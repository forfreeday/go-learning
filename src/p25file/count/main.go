package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	chCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

func main() {
	fileName := "/Users/liukai/test.log"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Open file err=%v\n", err)
		return
	}
	defer file.Close()
	// var count CharCount

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		reader.ReadSlice('\n')
		if err == io.EOF {
			break
		}
		for _, v := range str {
			fmt.Println(v)
			switch v {
			case v >= 'a' && v <= 'z':
				counter.chCount
			}
		}
	}
}
