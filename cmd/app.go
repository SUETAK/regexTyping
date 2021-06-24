package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("input....")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println("out put", scanner.Text())
}
