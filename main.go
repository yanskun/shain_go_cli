package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/c-bata/go-prompt"
)

func menu(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "1", Description: "社員を登録"},
		{Text: "2", Description: "社員を参照"},
		{Text: "3", Description: "社員を更新"},
		{Text: "4", Description: "社員を削除"},
		{Text: "5", Description: "終わる"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func inputName() {
	fmt.Print("continue? (Y/n) >")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		p := scanner.Text()
		fmt.Println("in: ", p)
		switch p {
		case "", "Y":
			fmt.Println("デフォルトの処理をします")
		case "n":
			fmt.Println("何もしない")
		default:
			fmt.Println("コマンドが不正なのでもう一度入力を促す")
			continue
		}
	}
}

type Employee struct {
	name     string
	birthday int
	isMan    bool
	salary   int
}

// type Employees *[]Employee

func CreateEmployee() {
	employee := &Employee{"taro", 20001231, true, 300}

	employees := []Employee{*employee}

	fmt.Println(employees)
	inputName()
}

func main() {
	fmt.Println("メニュー選択してね")
	selectedMenu := prompt.Input("> ", menu)
	switch selectedMenu {
	case "1":
		CreateEmployee()
	}
}
