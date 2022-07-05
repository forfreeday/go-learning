package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {

	var buffer bytes.Buffer

	for i := 0; i < 100; i++ {
		buffer.WriteString("this is line " + fmt.Sprintf("%d", i) + "\n")
	}

	if err := ioutil.WriteFile("testFile", buffer.Bytes(), 0644); err != nil {
		fmt.Println(err)
	}

}
