package P05error

import (
	"errors"
)

func Do() (int, error) {
	return 222, errors.New("异常")
}
