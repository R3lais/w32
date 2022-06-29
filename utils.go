package w32

import (
	"syscall"
	"unsafe"
)

func StrToUTF16Ptr(str string) *uint16 {
	v, _ := syscall.UTF16PtrFromString(str)
	return v
}

func StrToUTF16(str string) []uint16 {
	v, _ := syscall.UTF16FromString(str)
	return v
}

func UTF16ToStr(s []uint16) string {
	return syscall.UTF16ToString(s)
}

func GetLastError() error {
	return syscall.GetLastError()
}

func SizeOf(x interface{}) uintptr {
	return unsafe.Sizeof(x)
}

func NewCallback(x interface{}) uintptr {
	return syscall.NewCallback(x)
}

func Pointer(x any) unsafe.Pointer {
	return x.(unsafe.Pointer)
}
