package utility

import "testing"

type Paths struct {
	path, ext string
}

var exttests = []Paths{
	{"path.go", "go"},
	{"abc/path.go", "go"},
	{"abc/def", ""},
	{"videos/moana.mp4", "mp4"},
}

func TestFileExt(t *testing.T) {
	for _, test := range exttests {
		if x := FileExt(test.path); x != test.ext {
			t.Errorf("Test Failed. FileExt(%q) = %q want %q", test.path, x, test.ext)
		}
	}
}
