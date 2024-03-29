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
	// PageSizeA4 is a PageSize of type A4.
	PageSizeA4 PageSize = "A4"
	// PageSizeB4 is a PageSize of type B4.
	PageSizeB4 PageSize = "B4"
	// PageSizeA is a PageSize of type A.
	PageSizeA PageSize = "A"
	// PageSizeArchA is a PageSize of type Arch-A.
	PageSizeArchA PageSize = "Arch-A"
	// PageSizeLetter is a PageSize of type Letter.
	PageSizeLetter PageSize = "Letter"
)

var ErrInvalidPageSize = fmt.Errorf("not a valid PageSize, try [%s]", strings.Join(_PageSizeNames, ", "))

var _PageSizeNames = []string{
	string(PageSizeA4),
	string(PageSizeB4),
	string(PageSizeA),
	string(PageSizeArchA),
	string(PageSizeLetter),
}

// PageSizeNames returns a list of possible string values of PageSize.
func PageSizeNames() []string {
	tmp := make([]string, len(_PageSizeNames))
	copy(tmp, _PageSizeNames)
	return tmp
}

// String implements the Stringer interface.
func (x PageSize) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x PageSize) IsValid() bool {
	_, err := ParsePageSize(string(x))
	return err == nil
}

var _PageSizeValue = map[string]PageSize{
	"A4":     PageSizeA4,
	"a4":     PageSizeA4,
	"B4":     PageSizeB4,
	"b4":     PageSizeB4,
	"A":      PageSizeA,
	"a":      PageSizeA,
	"Arch-A": PageSizeArchA,
	"arch-a": PageSizeArchA,
	"Letter": PageSizeLetter,
	"letter": PageSizeLetter,
}

// ParsePageSize attempts to convert a string to a PageSize.
func ParsePageSize(name string) (PageSize, error) {
	if x, ok := _PageSizeValue[name]; ok {
		return x, nil
	}
	// Case insensitive parse, do a separate lookup to prevent unnecessary cost of lowercasing a string if we don't need to.
	if x, ok := _PageSizeValue[strings.ToLower(name)]; ok {
		return x, nil
	}
	return PageSize(""), fmt.Errorf("%s is %w", name, ErrInvalidPageSize)
}
