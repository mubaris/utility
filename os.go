// Package utility package contains bunch of functions to help you write better golang code.
package utility

import (
	"runtime"
)

// OsType returns the type of Operating System you're running
// Possible values are,
// android darwin dragonfly freebsd linux nacl netbsd openbsd plan9 solaris windows zos
func OsType() string {
	return runtime.GOOS
}

// OsArch returns the architecture of Operating System you're running
// Possible values are,
// 386 amd64 amd64p32 arm armbe arm64 arm64be ppc64 ppc64le mips mipsle mips64 mips64le mips64p32 mips64p32le ppc s390 s390x sparc sparc64
func OsArch() string {
	return runtime.GOARCH
}
