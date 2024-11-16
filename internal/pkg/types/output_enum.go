// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package types

import (
	"fmt"
	"strings"
)

const (
	// OutputTypePdf is a OutputType of type pdf.
	OutputTypePdf OutputType = "pdf"
	// OutputTypeHtml is a OutputType of type html.
	OutputTypeHtml OutputType = "html"
)

var ErrInvalidOutputType = fmt.Errorf("not a valid OutputType, try [%s]", strings.Join(_OutputTypeNames, ", "))

var _OutputTypeNames = []string{
	string(OutputTypePdf),
	string(OutputTypeHtml),
}

// OutputTypeNames returns a list of possible string values of OutputType.
func OutputTypeNames() []string {
	tmp := make([]string, len(_OutputTypeNames))
	copy(tmp, _OutputTypeNames)
	return tmp
}

// String implements the Stringer interface.
func (x OutputType) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x OutputType) IsValid() bool {
	_, err := ParseOutputType(string(x))
	return err == nil
}

var _OutputTypeValue = map[string]OutputType{
	"pdf":  OutputTypePdf,
	"html": OutputTypeHtml,
}

// ParseOutputType attempts to convert a string to a OutputType.
func ParseOutputType(name string) (OutputType, error) {
	if x, ok := _OutputTypeValue[name]; ok {
		return x, nil
	}
	return OutputType(""), fmt.Errorf("%s is %w", name, ErrInvalidOutputType)
}
