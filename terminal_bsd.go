// +build darwin freebsd openbsd netbsd dragonfly

package banner

import "syscall"

const ioctlReadTermios = syscall.TIOCGETA
