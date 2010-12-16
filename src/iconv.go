//
// iconv.go
//
package iconv

// #include "iconv_wrapper.h"
import "C"

import (
    "unsafe";
)

type Iconv struct {
    pointer unsafe.Pointer
}

func Open(tocode string, fromcode string) *Iconv {
    return &Iconv{C.IconvOpen(C.CString(tocode), C.CString(fromcode))};
}

func (cd *Iconv) Close() int {
    return int(C.IconvClose(cd.pointer));
}

func (cd *Iconv) Conv(inbuf string) string {
    outbuf := C.CString(*new(string));
    C.IconvIconv(cd.pointer, C.CString(inbuf), &outbuf);

    return C.GoString(outbuf);
}

