package globals

import (
	"errors"
	"collexy/helpers"
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
	parsed := helpers.ParseArray(asString)
	(*s) = StringSlice(parsed)
 
	return nil
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