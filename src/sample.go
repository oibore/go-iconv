//
// sample.go
//
package main

import (
    "fmt";
    //"strings";
    "iconv";
)

func main() {
    tocode   := "SJIS";
    fromcode := "UTF-8";

    str := "これは漢字です。";
    //fmt.Printf("%s\n", str);

    cd := iconv.Open(tocode, fromcode);

    str = iconv.Iconv(cd, str );
    fmt.Printf("str='%s'\n", str );

    iconv.Close(cd);
}

