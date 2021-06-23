package shain

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Shain struct {
	ID       int
	Name     string
	Birthday time.Time
	IsMan    bool
	Salary   int
}

func inputText() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	p := scanner.Text()
	return p
}

func inputName() string {
	fmt.Println("名前を入力")

	return inputText()
}

func Create(id int) Shain {

	var name = inputName()
	var user = &Shain{}

	user.ID = id
	user.Name = name
	user.Birthday = time.Now()
	user.IsMan = true
	user.Salary = 300

	return *user
}

func Update(shain Shain) Shain {
	shain.Name = "jiro"

	return shain
}
