//
// iconv.go
//
package iconv

// #include "iconv_wrapper.h"
import "C"

import (
    "unsafe";
)

func Open(tocode string, fromcode string) unsafe.Pointer {
    return C.IconvOpen(C.CString(tocode), C.CString(fromcode));
}

func Close(cd unsafe.Pointer) int {
    return int(C.IconvClose(cd));
}

func Iconv(cd unsafe.Pointer, inbuf  string) string {

    outbuf := C.CString(*new(string));
    C.IconvIconv(cd, C.CString(inbuf), &outbuf);

    return C.GoString(outbuf);
}

