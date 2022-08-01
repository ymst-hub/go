package main

import (
	"bufio"//標準入力取得部分
	"fmt"//出力
	"os"//標準入力部分
)

func main(){
	name := input("type your name")
	fmt.Println("Hello,"+name+"!!")
}

func input(msg string)string{
	//scannerという標準入力の型を用意する
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("msg"+": ")
	//標準入力を取得する
	scanner.Scan()
	//Scanした入力を引き渡す
	return scanner.Text()
}