//
// iconv.go
//
package iconv

// #include <iconv.h>
// #include "iconv_wrapper.h"
import "C"

import (
	"os"
)

type Iconv struct {
	pointer C.iconv_t
}

func Open(tocode string, fromcode string) (*Iconv, os.Error) {
	ret, err := C.iconv_open(C.CString(tocode), C.CString(fromcode))
	if err != nil {
		return nil, err
	}
	return &Iconv{ret}, nil
}

func (cd *Iconv) Close() os.Error {
	_, err := C.iconv_close(cd.pointer)
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
