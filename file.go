package main

import(
	"os"
	"fmt"
	"bufio"
)
func main(){
	//値の書き込みをまとめている
	wt := func(f *os.File,s string){
		_,er := f.WriteString(s + "\n")
		if er != nil{
			panic(er)
		}
	}
	fn := "data.txt"
	f,er := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY,os.ModePerm)
	//書き出し専用モードで、なければ作り、あれば最後に追記する
	if er != nil{
		//fmt.Println(er)
		//return
		panic(er)//エラーにする
	}

	defer f.Close()//最後に必ず行うことを記載する(この場合はファイルを閉じる)

	fmt.Println("***start***")
	wt(f,"***start***")//ファイルに書き込む
	for{
		s := input("type message")
		if s == ""{
			break
		}
		wt(f,s)
	}
	wt(f,"***end***\n\n")
	fmt.Println("***end***")
	er = f.Close()
	if er != nil{
		fmt.Println(er)
	}

}

func input(msg string)string{
	//scannerという標準入力の型を用意する
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("msg"+": ")
	//標準入力を取得する
	scanner.Scan()
	//Scanした入力を引き渡す
	return scanner.Text()
}