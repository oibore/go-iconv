//
// iconv.go
//
package iconv

// #include <iconv.h>
// #include <errno.h>
import "C"

import (
	"bytes"
	"syscall"
	"unsafe"
)

type Iconv struct {
	pointer C.iconv_t
}

func Open(tocode string, fromcode string) (*Iconv, error) {
	ret, err := C.iconv_open(C.CString(tocode), C.CString(fromcode))
	if err != nil {
		return nil, err
	}
	return &Iconv{ret}, nil
}

func (cd *Iconv) Close() error {
	_, err := C.iconv_close(cd.pointer)
	return err
}

func (cd *Iconv) Conv(input string) (result string, err error) {
	var buf bytes.Buffer

	if len(input) == 0 {
		return "", nil
	}

	inbuf := []byte(input)
	outbuf := make([]byte, len(inbuf))
	inbytes := C.size_t(len(inbuf))
	inptr := &inbuf[0]

	for inbytes > 0 {
		outbytes := C.size_t(len(outbuf))
		outptr := &outbuf[0]
		_, err = C.iconv(cd.pointer,
			(**C.char)(unsafe.Pointer(&inptr)), &inbytes,
			(**C.char)(unsafe.Pointer(&outptr)), &outbytes)
		buf.Write(outbuf[:len(outbuf)-int(outbytes)])
		if err != nil && err != syscall.E2BIG {
			return buf.String(), err
		}
	}

	return buf.String(), nil
}
