package main

import "fmt"
import "golearning/src/P05error"

func main() {
	result, err := P05error.Do()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
