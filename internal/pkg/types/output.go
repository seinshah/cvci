package types

import (
	"context"
	"path/filepath"
	"strings"
)

//go:generate go-enum --names

// OutputType is the type of CV output being generated by the binary.
// ENUM(pdf, html).
type OutputType string

// OutputGenerator is an interface that each output generator need to implement.
// These generators receive a parsed HTML template and need to generate a proper output.
type OutputGenerator interface {
	Generate(ctx context.Context, content []byte) ([]byte, error)
}

func DetectOutputType(outputPath string) OutputType {
	ext := strings.Replace(filepath.Ext(outputPath), ".", "", 1)

	return OutputType(strings.ToLower(ext))
}
