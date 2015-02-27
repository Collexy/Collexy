package helpersOld
 
import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"regexp"
	"strings"
)
 
func main() {
	db := dbConnect()
 
	makeTestTables(db)
 
	defer db.Close()
	defer cleanup(db)
 
	// Insert Some Data
	db.Exec(`INSERT INTO array_test VALUES ('{"String1", "String2"}')`)
 
	// arrays can be selected as strings...
	dataString := selectAsString(db)
	fmt.Println("SELECT as String:", dataString)
 
	// Or by using array functions...
	dataUnnest := selectUsingUnnest(db)
	fmt.Println("SELECT using Unnest:", dataUnnest)
 
	// Or by defining a scan type and parsing the return value
	dataSlice := selectAsSlice(db)
	fmt.Println("SELECT by parsing:", dataSlice)
 
	// Arrays can be updated by replacing the entire array:
	newArray := []interface{}{"String1", "String3", "String4", "String5"}
	updateArray(db, newArray)
	dataSlice = selectAsSlice(db)
	fmt.Println("UPDATE entire array", dataSlice)
 
	// or by appending / prepending value(s):
	AppendToArray(db, "String6")
	dataSlice = selectAsSlice(db)
	fmt.Println("UPDATE with append:", dataSlice)
 
	// or by replacing individual values:
	ReplaceInArray(db, 2, "NULL")
	dataSlice = selectAsSlice(db)
	fmt.Println("UPDATE with replace:", dataSlice)
 
	// Deleting by index requires slicing and is inefficient:
	DeleteFromArray(db, 3)
	dataSlice = selectAsSlice(db)
	fmt.Println("UPDATE deleting index:", dataSlice)
 
}
 
// Arrays are serialized to strings {value, value...} by the database.
// these strings selected, updated and inserted like any string
func selectAsString(db *sql.DB) string {
	row := db.QueryRow("SELECT data FROM array_test")
	var asString string
	err := row.Scan(&asString)
	if err != nil {
		panic(err)
	}
	return asString
}
 
// The UNNEST function expands an array into multiple rows.  Each row
// can then be scanned individually.
func selectUsingUnnest(db *sql.DB) []string {
	results := make([]string, 0)
	rows, err := db.Query("SELECT UNNEST(data) FROM array_test")
	if err != nil {
		panic(err)
	}
	var scanString string
	for rows.Next() {
		rows.Scan(&scanString)
		results = append(results, scanString)
	}
	return results
}
 
// By defining a wrapper type around a slice which implements
// sql.Scanner, we can scan the array directly into the type.
func selectAsSlice(db *sql.DB) StringSlice {
	row := db.QueryRow("SELECT data FROM array_test")
	var asSlice StringSlice
	err := row.Scan(&asSlice)
	if err != nil {
		panic(err)
	}
	return asSlice
}
 
// Update an array by replacing the whole array with new values.
// This _could_ be done by serializing the StringSlice type using
// sql.driver.Valuer, but then we would have to validate the type
// of each value manually and format it for insert by hand.  Instead,
// the ARRAY[...] format allows us to use query parameters to construct
// the array, ie ARRAY[$1, $2, $3], which then allows the database
// driver to coerce the variables into the right format for us.
func updateArray(db *sql.DB, array []interface{}) {
	params := make([]string, 0, len(array))
	for i := range array {
		params = append(params, fmt.Sprintf("$%v", i+1))
	}
	query := fmt.Sprintf("UPDATE array_test SET data = ARRAY[%s]", strings.Join(params, ", "))
	db.Exec(query, array...)
}
 
// The ARRAY_APPEND and ARRAY_PREPEND functions can be used to add single
// values to arrays.  ARRAY_CAT combines two arrays.  The || operator can
// do the same thing:
//   SET data = data || <value>
//   SET data = data || ARRAY[<value1>, <value2>]
 
func AppendToArray(db *sql.DB, value string) {
	_, err := db.Exec("UPDATE array_test SET data = ARRAY_APPEND(data, $1)", value)
	if err != nil {
		panic(err)
	}
}
 
// Arrays are 1-indexed. Individual elements can be used in expressions,
// updated, or selected by indexing the array.
func ReplaceInArray(db *sql.DB, index int, newValue string) {
	_, err := db.Exec("UPDATE array_test SET data[$1] = $2", index, newValue)
	if err != nil {
		panic(err)
	}
}
 
// Arrays support slice indexing:
//    ARRAY['a', 'b', 'c'][1:2] == ARRAY['a', 'b']
// The ARRAY_UPPER function gets the length of an array for a specified dimension
// Deleting a value from an array amounts to slicing the array into two parts
// and combining them back together.
func DeleteFromArray(db *sql.DB, i int) {
	_, err := db.Exec("UPDATE array_test SET data = array_cat(data[0:$1], data[$2:ARRAY_UPPER(data, 1) + 1])", i-1, i+1)
	if err != nil {
		panic(err)
	}
}
 
type StringSlice []string
 
// Implements sql.Scanner for the String slice type
// Scanners take the database value (in this case as a byte slice)
// and sets the value of the type.  Here we cast to a string and
// do a regexp based parse
func (s *StringSlice) Scan(src interface{}) error {
	asBytes, ok := src.([]byte)
	if !ok {
		return error(errors.New("Scan source was not []bytes"))
	}
 
	asString := string(asBytes)
	parsed := parseArray(asString)
	(*s) = StringSlice(parsed)
 
	return nil
}
 
// PARSING ARRAYS
// SEE http://www.postgresql.org/docs/9.1/static/arrays.html#ARRAYS-IO
// Arrays are output within {} and a delimiter, which is a comma for most
// postgres types (; for box)
//
// Individual values are surrounded by quotes:
// The array output routine will put double quotes around element values if
// they are empty strings, contain curly braces, delimiter characters,
// double quotes, backslashes, or white space, or match the word NULL.
// Double quotes and backslashes embedded in element values will be
// backslash-escaped. For numeric data types it is safe to assume that double
// quotes will never appear, but for textual data types one should be prepared
// to cope with either the presence or absence of quotes.
 
// construct a regexp to extract values:
var (
	// unquoted array values must not contain: (" , \ { } whitespace NULL)
	// and must be at least one char
	unquotedChar  = `[^",\\{}\s(NULL)]`
	unquotedValue = fmt.Sprintf("(%s)+", unquotedChar)
 
	// quoted array values are surrounded by double quotes, can be any
	// character except " or \, which must be backslash escaped:
	quotedChar  = `[^"\\]|\\"|\\\\`
	quotedValue = fmt.Sprintf("\"(%s)*\"", quotedChar)
 
	// an array value may be either quoted or unquoted:
	arrayValue = fmt.Sprintf("(?P<value>(%s|%s))", unquotedValue, quotedValue)
 
	// Array values are separated with a comma IF there is more than one value:
	arrayExp = regexp.MustCompile(fmt.Sprintf("((%s)(,)?)", arrayValue))
 
	valueIndex int
)
 
// Find the index of the 'value' named expression
func init() {
	for i, subexp := range arrayExp.SubexpNames() {
		if subexp == "value" {
			valueIndex = i
			break
		}
	}
}
 
// Parse the output string from the array type.
// Regex used: (((?P<value>(([^",\\{}\s(NULL)])+|"([^"\\]|\\"|\\\\)*")))(,)?)
func parseArray(array string) []string {
	results := make([]string, 0)
	matches := arrayExp.FindAllStringSubmatch(array, -1)
	for _, match := range matches {
		s := match[valueIndex]
		// the string _might_ be wrapped in quotes, so trim them:
		s = strings.Trim(s, "\"")
		results = append(results, s)
	}
	return results
}
 
// DB HELPERs
 
func dbConnect() *sql.DB {
	datname := os.Getenv("PGDATABASE")
	sslmode := os.Getenv("PGSSLMODE")
 
	if datname == "" {
		os.Setenv("PGDATABASE", "pqgotest")
	}
 
	if sslmode == "" {
		os.Setenv("PGSSLMODE", "disable")
	}
 
	conn, err := sql.Open("postgres", "")
	if err != nil {
		panic(err)
	}
 
	return conn
}
 
// Create a table with an array type
// Can also use the syntax CREATE TABLE array_test (data varchar ARRAY)
func makeTestTables(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE array_test (data varchar[])")
	if err != nil {
		panic(err)
	}
}
 
func cleanup(db *sql.DB) {
	db.Exec("DROP TABLE array_test")
}