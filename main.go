package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/c-bata/go-prompt"
	"github.com/yasudanaoya/shain_go_cli/shain"
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

func CreateShain(list []shain.Shain) {
	var user = &shain.Shain{}
	// TODO: 入力できるように
	user.ID = 1
	user.Name = "taro"
	user.Birthday = 20001231
	user.IsMan = true
	user.Salary = 300

	newList := append(list, *user)

	showMenu(newList)
}

func ShowShainList(list []shain.Shain) {
	str := fmt.Sprintf("%v \n", list)
	fmt.Print(str)

	showMenu(list)
}

func UpdateShain(list []shain.Shain) {
	// TODO: 更新

	showMenu(list)
}

func DeleteShain(list []shain.Shain) {
	newList := []shain.Shain{}
	// TODO: 🆔 を入力できるように
	for _, shain := range list {
		if shain.ID != 1 {
			newList = append(newList, shain)
		}
	}

	showMenu(newList)
}

func showMenu(list []shain.Shain) {
	selectedMenu := prompt.Input("> ", menu)
	switch selectedMenu {
	case "1":
		CreateShain(list)
	case "2":
		ShowShainList(list)
	case "3":
		UpdateShain(list)
	case "4":
		DeleteShain(list)
	}
}

func main() {
	list := []shain.Shain{}
	fmt.Println("メニュー選択してね")
	showMenu(list)
}
