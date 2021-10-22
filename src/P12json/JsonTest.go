package P12json

import (
	"encoding/json"
	"fmt"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

type Post2 struct {
	Id      int      `json:"ID"`
	Content string   `json:"content"`
	Author  string   `json:"author"`
	Label   []string `json:"label"`
}

func JsonParseTest() {
	post := &Post{1, "Hello World", "userA"}
	//转换结构体为 json
	b, err := json.Marshal(post)
	if err != nil {
		fmt.Println(nil)
	}
	fmt.Println(string(b))
}

func JsonParseTest2() {
	post := &Post{1, "Hello World", "userA"}
	c, err := json.MarshalIndent(post, "", "\t")
	if err != nil {
		fmt.Println(nil)
	}
	fmt.Println(string(c))
}

func JsonParseTest3() {
	// slice -> json
	s := []string{"a", "b", "c"}
	d, _ := json.MarshalIndent(s, "", "\t")
	fmt.Println(string(d))

	// map -> json
	m := map[string]string{
		"a": "aa",
		"b": "bb",
		"c": "cc",
	}
	e, _ := json.MarshalIndent(m, "", "\t")
	fmt.Println(string(e))
}

func JsonParseTest4() {
	postp := &Post2{
		2,
		"Hello World",
		"userB",
		[]string{"linux", "shell"},
	}

	p, _ := json.MarshalIndent(postp, "", "\t")
	fmt.Println(string(p))
}
