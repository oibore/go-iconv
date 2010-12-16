//
// iconv.go
//
package iconv

// #include "iconv_wrapper.h"
import "C"

import (
	"unsafe"
	"os"
)

type Iconv struct {
	pointer unsafe.Pointer
}

func Open(tocode string, fromcode string) (*Iconv, os.Error) {
	ret, err := C.IconvOpen(C.CString(tocode), C.CString(fromcode))
	if err != nil {
		return nil, err
	}
	return &Iconv{ret}, nil
}

func (cd *Iconv) Close() os.Error {
	_, err := C.IconvClose(cd.pointer)
	return err
}

func (cd *Iconv) Conv(inbuf string) (string, os.Error) {
	outbuf := C.CString(*new(string))
	_, err := C.IconvIconv(cd.pointer, C.CString(inbuf), &outbuf)
	if err != nil {
		return "", err
	}

	return C.GoString(outbuf), nil
}
