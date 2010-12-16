//
// iconv_test.go
//
package iconv

import (
	"testing"
	"fmt"
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
