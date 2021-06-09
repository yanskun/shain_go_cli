package employee

import (
	"fmt"
)

type employee struct {
	name     string
	birthday int
	isMan    bool
	salary   int
}

type employees []*employee

func CreateEmployee() {
	e := employee{"taro", 20001231, true, 300}
	fmt.Println(e)
}
