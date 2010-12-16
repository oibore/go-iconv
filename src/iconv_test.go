//
// iconv_test.go
//
package iconv

import (
	"testing"
	"fmt"
	"os"
)

func TestIconv(t *testing.T) {
	tocode := "SJIS"
	fromcode := "UTF-8"

	str := "これは漢字です。"
	//fmt.Printf("%s\n", str);

	cd, err := Open(tocode, fromcode)
	if err != nil {
		t.Errorf("Error on opening: %s\n", err)
		return
	}

	str, err = cd.Conv(str)
	if err != nil {
		t.Errorf("Error on conversion: %s\n", err)
		return
	}
	fmt.Printf("str='%s'\n", str)

	err = cd.Close()
	if err != nil {
		t.Errorf("Error on close: %s\n", err)
	}
}

func TestError(t *testing.T) {
	_, err := Open("INVALID_ENCODING", "INVALID_ENCODING")
	if err != os.EINVAL {
		t.Errorf("Unexpected error: %#s (expected %#s)", err, os.EINVAL)
	}

	cd, _ := Open("ISO-8859-15", "UTF-8")
	_, err = cd.Conv("\xc3a")
	if err != EILSEQ {
		t.Errorf("Unexpected error: %#s (expected %#s)", err, EILSEQ)
	}
}
