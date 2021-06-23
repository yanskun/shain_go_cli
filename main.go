package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/yasudanaoya/shain_go_cli/shain"
)

func showMenu(list []shain.Shain) {
	fmt.Println("メニュー選択してね")
	fmt.Print(`1. 社員を登録 
		2. 社員を参照
		3. 社員を更新
		4. 社員を削除
		5. 終わる
	`)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		p := scanner.Text()
		switch p {
		case "1":
			createShain(list)
		case "2":
			showShainList(list)
		case "3":
			updateShain(list)
		case "4":
			deleteShain(list)
		default:
			fmt.Println("コマンドが不正なのでもう一度入力を促す")
			continue
		}
	}
}

func createShain(list []shain.Shain) {
	var id = len(list)

	newList := append(list, shain.Create(id))

	showMenu(newList)
}

func showShainList(list []shain.Shain) {
	str := fmt.Sprintf("%v \n", list)
	fmt.Print(str)

	showMenu(list)
}

func updateShain(list []shain.Shain) {
	// TODO: 更新

	showMenu(list)
}

func deleteShain(list []shain.Shain) {
	newList := []shain.Shain{}
	// TODO: 🆔 を入力できるように
	for _, shain := range list {
		if shain.ID != 1 {
			newList = append(newList, shain)
		}
	}

	showMenu(newList)
}

func main() {
	list := []shain.Shain{}
	showMenu(list)
}
