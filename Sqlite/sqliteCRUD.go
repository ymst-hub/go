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

func main(){
	con,er := sql.Open("sqlite3","data.sqlite3")
	if er != nil {
		panic(er)
	}
	defer con.Close()

	nm := input("name")
	ml := input("mail")
	age := input("age")
	ag,_ := strconv.Atoi(age)
	qry := "insert into mydata (name,mail,age) values(?,?,?)"
	con.Exec(qry,nm,ml,ag)
	showRecord(con)
	
}

func showRecord(con *sql.DB){
	qry := "select * from mydata"
	rs,_ := con.Query(qry)
	for rs.Next(){
		fmt.Println(MydatafmRws(rs).Str())
	}
}

func MydatafmRws(rs *sql.Rows) *Mydata{
	var md Mydata
	er := rs.Scan(&md.ID,&md.Name,&md.Mail,&md.age)
	if er != nil{
		panic(er)
	}
	return &md
}

func MydatafmRw(rs *sql.Row) *Mydata{
	var md Mydata
	er := rs.Scan(&md.ID,&md.Name,&md.Mail,&md.age)
	if er != nil {
		panic(er)
	}
	return &md
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
