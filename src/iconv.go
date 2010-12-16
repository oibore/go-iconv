//
// iconv.go
//
package iconv

// #include <iconv.h>
// #include <errno.h>
import "C"

import (
	"os"
	"unsafe"
)

var EILSEQ = os.Errno(int(C.EILSEQ))

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

func (cd *Iconv) Conv(input string) (string, os.Error) {
	inbuf := []byte(input)
	outbuf := make([]byte, len(inbuf)*8)
	inbytes := C.size_t(len(inbuf))
	outbytes := C.size_t(len(outbuf))
	inptr := &inbuf[0]
	outptr := &outbuf[0]

	_, err := C.iconv(cd.pointer, (**C.char)(unsafe.Pointer(&inptr)), &inbytes,
	                              (**C.char)(unsafe.Pointer(&outptr)), &outbytes)
	if err != nil {
		return "", err
	}

	return string(outbuf[:len(outbuf)-int(outbytes)]), nil
}
