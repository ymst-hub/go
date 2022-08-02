package main

//最初に処理する関数名
import (
	"fmt"
	//パッケージ
)

func main(){
	fmt.Println("Hello Go-lang!")
	//実行する処理
}

//go run ~~で実行、go build ~~でexeなどを生成できる

/*
関数の構成
func 関数名 (引数) 戻り値{
処理
}
//戻り値は(戻り値１,戻り値２...)と記載可能
//また、
//引数は(a []string,v ...string)とすれば
関数名(a,v1,v2,v3)
//として使用できる
利用はv...でOK
*/

/*
無名関数
変数 := func(引数)戻り値{処理}
利用方法は、変数()で呼び出し可能

func(){
	即時実行する処理
}()
*/
