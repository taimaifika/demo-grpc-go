package main

import (
	"encoding/json"
	"fmt"
)

type QueryRepo struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Extra int    `json:"extra"`
}

type DataRes struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Extra int    `json:"extra1,omitempty"`
}

func main() {
	foo := QueryRepo{
		Id:    "1",
		Name:  "foo",
		// Extra: 1,
	}
	fmt.Println(foo)
	fooStr, _ := json.Marshal(foo)
	fmt.Println(string(fooStr))

	bar := DataRes(foo)

	fmt.Println(bar)
	barStr, _ := json.Marshal(bar)
	fmt.Println(string(barStr))
}
