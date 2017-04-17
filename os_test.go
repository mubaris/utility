package utility

import "testing"

func TestOsType(t *testing.T) {
	if OsType() == "" {
		t.Errorf("Test Failed")
	}
}

func TestOsArch(t *testing.T) {
	if OsArch() == "" {
		t.Errorf("Test Failed")
	}
}
