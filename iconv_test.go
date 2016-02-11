//
// iconv_test.go
//
package iconv

import (
	"syscall"
	"testing"
)

var testData = []struct{ utf8, other, otherEncoding string }{
	{"これは漢字です。", "\x82\xb1\x82\xea\x82\xcd\x8a\xbf\x8e\x9a\x82\xc5\x82\xb7\x81B", "SJIS"},
	{"これは漢字です。", "S0\x8c0o0\"oW[g0Y0\x020", "UTF-16LE"},
	{"これは漢字です。", "0S0\x8c0oo\"[W0g0Y0\x02", "UTF-16BE"},
	{"€1 is cheap", "\xa41 is cheap", "ISO-8859-15"},
	{"", "", "SJIS"},
}

func TestIconv(t *testing.T) {
	for _, data := range testData {
		cd, err := Open("UTF-8", data.otherEncoding)
		if err != nil {
			t.Errorf("Error on opening: %s\n", err)
			continue
		}

		str, err := cd.Conv(data.other)
		if err != nil {
			t.Errorf("Error on conversion: %s\n", err)
			continue
		}

		if str != data.utf8 {
			t.Errorf("Unexpected value: %#v (expected %#v)", str, data.utf8)
		}

		err = cd.Close()
		if err != nil {
			t.Errorf("Error on close: %s\n", err)
		}
	}
}

func TestIconvReverse(t *testing.T) {
	for _, data := range testData {
		cd, err := Open(data.otherEncoding, "UTF-8")
		if err != nil {
			t.Errorf("Error on opening: %s\n", err)
			continue
		}

		str, err := cd.Conv(data.utf8)
		if err != nil {
			t.Errorf("Error on conversion: %s\n", err)
			continue
		}

		if str != data.other {
			t.Errorf("Unexpected value: %#v (expected %#v)", str, data.other)
		}

		err = cd.Close()
		if err != nil {
			t.Errorf("Error on close: %s\n", err)
		}
	}
}

func TestError(t *testing.T) {
	_, err := Open("INVALID_ENCODING", "INVALID_ENCODING")
	if err != syscall.EINVAL {
		t.Errorf("Unexpected error: %#s (expected %#s)", err, syscall.EINVAL)
	}

	cd, _ := Open("ISO-8859-15", "UTF-8")
	_, err = cd.Conv("\xc3a")
	if err != syscall.EILSEQ {
		t.Errorf("Unexpected error: %#s (expected %#s)", err, syscall.EILSEQ)
	}
}

func TestTranslit(t *testing.T) {
	result, err := Conv("1€ is nothing nowadays", "ASCII//TRANSLIT", "UTF-8")
	if err != nil {
		t.Errorf("Error on conversion: %s", err)
		return
	}
	if result != "1EUR is nothing nowadays" {
		t.Errorf("Unexpected transliteration: %s", result)
		return
	}

	result, err = Conv("1€ is nothing nowadays", "ASCII", "UTF-8")
	if err != syscall.EILSEQ {
		t.Errorf("Unexpected error: %s", err)
		return
	}
}
