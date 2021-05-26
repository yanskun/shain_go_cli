package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("input? ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println("input is", scanner.Text())
}
