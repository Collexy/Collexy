package globals

import (
	"errors"
	//corehelpers "collexy/core/helpers"
	"database/sql/driver"
	"fmt"
	"regexp"
	"strings"
)

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
	parsed := ParseArray(asString)
	(*s) = StringSlice(parsed)

	return nil
}

func (b StringSlice) Value() (driver.Value, error) {
	var str string = "{"
	var myarr []string = b
	fmt.Println("driver.Value 1: ")
	fmt.Println(b)
	for i := 0; i < len(myarr); i++ {
		str = str + myarr[i]
		if i < len(myarr)-1 {
			str = str + ","
		}
	}
	str = str + "}"
	fmt.Println("driver.Value 2: ")
	fmt.Println(str)
	return str, nil
}

func RemoveDuplicatesStringSlice(xs *[]string) {
	found := make(map[string]bool)
	j := 0
	for i, x := range *xs {
		if !found[x] {
			found[x] = true
			(*xs)[j] = (*xs)[i]
			j++
		}
	}
	*xs = (*xs)[:j]
}

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
	//arrayExp = regexp.MustCompile(fmt.Sprintf("((%s)(,)?)", arrayValue))
	arrayExp = regexp.MustCompile(fmt.Sprintf("((%s)()?)", arrayValue))

	valueIndex int
)

// Parse the output string from the array type.
// Regex used: (((?P<value>(([^",\\{}\s(NULL)])+|"([^"\\]|\\"|\\\\)*")))(,)?)
func ParseArray(array string) []string {
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
