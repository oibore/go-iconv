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
var E2BIG = os.Errno(int(C.E2BIG))

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

func (cd *Iconv) Conv(input string) (result string, err os.Error) {
	if len(input) == 0 {
		return "", nil
	}

	inbuf := []byte(input)
	outbuf := make([]byte, len(inbuf))
	inbytes := C.size_t(len(inbuf))
	inptr := &inbuf[0]

	result = ""
	for inbytes > 0 {
		outbytes := C.size_t(len(outbuf))
		outptr := &outbuf[0]
		_, err = C.iconv(cd.pointer,
			(**C.char)(unsafe.Pointer(&inptr)), &inbytes,
			(**C.char)(unsafe.Pointer(&outptr)), &outbytes)
		result += string(outbuf[:len(outbuf)-int(outbytes)])
		if err != nil && err != E2BIG {
			return result, err
		}
	}

	return result, nil
}
