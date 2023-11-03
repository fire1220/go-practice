package main

import (
	"runtime"
	"syscall"
	"unsafe"
)

func SyscallWriteDarwin(fd int, msg string) int
func SyscallDarwin(trap, a1, a2, a3 uintptr) (r1, r2, r3 uintptr)

func main() {
	// Mac系统
	if runtime.GOOS == "darwin" {
		x := SyscallWriteDarwin(1, "hello world syscall!\n")
		println(x)
		by := []byte("hello syscall")
		p := unsafe.Pointer(&(by[0]))
		a, b, c := SyscallDarwin(syscall.SYS_WRITE,
			uintptr(syscall.Stdout), uintptr(p), uintptr(len(by)))
		println(a, b, c)
	}
}
