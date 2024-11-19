package loader_test

import (
	"os"
	"testing"

	"github.com/seinshah/cvci/internal/pkg/loader"
	"github.com/stretchr/testify/require"
)

type expectError struct {
	hasError bool
	err      error
}

func localFileSetUp(content string) func(t *testing.T) (string, func(t *testing.T)) {
	return func(t *testing.T) (string, func(t *testing.T)) {
		t.Helper()

		f, err := os.CreateTemp("/tmp", "cvci-test-*.txt")

		require.NoError(t, err)

		_, err = f.WriteString(content)

		require.NoError(t, err)

		return f.Name(), func(t *testing.T) {
			errD := f.Close()

			require.NoError(t, errD)
		}
	}
}

func TestNewGeneralLoader(t *testing.T) {
	t.Parallel()

	tests := []struct {
		expectError
		name         string
		setUp        func(t *testing.T) (string, func(t *testing.T))
		path         string
		validateType func(t *testing.T, l loader.Loader)
	}{
		{
			name:  "success-local",
			setUp: localFileSetUp("success-local-test"),
			validateType: func(t *testing.T, l loader.Loader) {
				_, ok := l.(*loader.LocalLoader)

				require.True(t, ok)
			},
		},
		{
			name: "success-remote",
			path: "https://raw.githubusercontent.com/seinshah/cvci/refs/heads/main/assets/sample_config.yaml",
			validateType: func(t *testing.T, l loader.Loader) {
				_, ok := l.(*loader.RemoteLoader)

				require.True(t, ok)
			},
		},
		{
			name: "undetected-loader",
			path: "invaid-path/",
			expectError: expectError{
				hasError: true,
				err:      loader.ErrInvalidPath,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			path := tc.path

			if tc.setUp != nil {
				p, cancel := tc.setUp(t)
				defer cancel(t)

				if p != "" {
					path = p
				}
			}

			l, err := loader.NewGeneralLoader(path)

			if tc.hasError {
				require.Error(t, err)

				if tc.err != nil {
					require.ErrorIs(t, err, tc.err)
				}

				return
			}

			require.NoError(t, err)
			require.NotNil(t, l)
			t.Logf("loader type %T", l.Loader())
			tc.validateType(t, l.Loader())
		})
	}
}