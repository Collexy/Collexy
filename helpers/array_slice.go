package helpers

import (
	"regexp"
	"strings"
	"fmt"
)

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

func StringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func IntInSlice(a int, list []int) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}