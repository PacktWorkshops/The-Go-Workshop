package main

// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	var cString *C.char
	cString = C.CString("Hello World!\n")
	defer C.free(unsafe.Pointer(cString))
	var b []byte
	b = C.GoBytes(unsafe.Pointer(cString), C.int(14))
	fmt.Print(string(b))
}
