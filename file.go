// Package utility package contains bunch of functions to help you write better golang code.
package utility

import (
	"path/filepath"
)

// FileExt returns the extension of the file passed.
// path is a the filepath of the file.
func FileExt(path string) string {
	ext := filepath.Ext(path)
	if len(ext) == 0 {
		return ext
	}
	return ext[1:]
}
