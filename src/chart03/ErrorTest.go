package chart03

import (
	"errors"
	"fmt"
)

func Do() (int, error) {
	return 0, errors.New("异常")
}

func main() {
	_, err := Do()
	if err != nil {
		fmt.Println(err)
	}
}
