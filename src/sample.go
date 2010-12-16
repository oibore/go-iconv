//
// sample.go
//
package main

import (
	"fmt"
	//"strings";
	"iconv"
)

func main() {
	tocode := "SJIS"
	fromcode := "UTF-8"

	str := "これは漢字です。"
	//fmt.Printf("%s\n", str);

	cd, err := iconv.Open(tocode, fromcode)
	if err != nil {
		fmt.Printf("Error on opening: %s\n", err)
		return
	}

	str, err = cd.Conv(str)
	if err != nil {
		fmt.Printf("Error on conversion: %s\n", err)
		return
	}
	fmt.Printf("str='%s'\n", str)

	err = cd.Close()
	if err != nil {
		fmt.Printf("Error on close: %s\n", err)
	}
}
