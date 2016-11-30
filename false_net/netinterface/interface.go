package netinterface

import (
	"io"
	"strings"
	"syscall"
	"unsafe"
)

type Interface struct {
	io.ReadWriteCloser
	IsTap bool
	IsTun bool
	Name  string
}

func New(rwc io.ReadWriteCloser, tap bool, name string) *Interface {
	return &Interface{
		ReadWriteCloser: rwc,
		IsTap:           tap,
		IsTun:           !tap,
		Name:            name,
	}
}

type IFReq struct {
	Name  [0x10]byte
	Flags uint16
	pad   [0x28 - 0x10 - 2]byte
}

func NewReq(fd uintptr, ifName string, flags uint16) (name string, err error) {
	var req IFReq
	req.Flags = flags
	copy(req.Name[:], ifName)
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, fd, (uintptr)(syscall.TUNSETIFF), uintptr(unsafe.Pointer(&req)))
	if errno != 0 {
		err = errno
		return
	}
	name = strings.Trim((string)(req.Name[:]), "\x00")
	return
}
