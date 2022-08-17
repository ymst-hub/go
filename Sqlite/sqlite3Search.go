package main

//sqliteの使い方、
//go mod into 名前
//go get github.com/mattn/go-sqlite3
//sqlite3 data.sqlite3 コマンド
//sqlを使う
import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

//Mydataを設定
type Mydata struct{
	ID int
	Name string
	Mail string
	age int
}
func (m *Mydata) Str() string{
	return "<\""+ strconv.Itoa(m.ID) + ":" + m.Name + "\" "+m.Mail + "," + strconv.Itoa(m.age) +">"
}

var qry string = "select * from mydata where id = ?"
//like検索 where name like ?
//qry,"%"+s+"%"

func main(){
	con,er := sql.Open("sqlite3","data.sqlite3")
	if er != nil{
		panic(er)
	}
	defer con.Close()

	for true {
		//begin
		s := input("id")
		if s == "" {
			break
		}
		n,er := strconv.Atoi(s)
		if er != nil{
			panic(er)
		}
		//rs,er := con.Query(qry,n)//nを追加できる
		rs := con.QueryRow(qry,n)
		/*
		queryRowは１つしか返さない
		if er != nil{
			panic(er)
		}
		
		for rs.Next(){
			var md Mydata
			er := rs.Scan(&md.ID , &md.Name , &md.Mail , &md.age)
			if er != nil{
				panic(er)
			}
			fmt.Println(md.Str())
		}
		*/
		var md Mydata
		er2 := rs.Scan(&md.ID , &md.Name , &md.Mail , &md.age)
		if er2 != nil{
			panic(er2)
		}
		fmt.Println(md.Str())

	}
	fmt.Println("***end***")
}

func input(msg string)string{
	//scannerという標準入力の型を用意する
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg+": ")
	//標準入力を取得する
	scanner.Scan()
	fmt.Println("")
	//Scanした入力を引き渡す
	return scanner.Text()
}