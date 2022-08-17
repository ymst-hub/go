package main

//sqliteの使い方、
//go mod into 名前
//go get github.com/mattn/go-sqlite3
//sqlite3 data.sqlite3 コマンド
//sqlを使う
import (
	"database/sql"
	"strconv"
	"fmt"
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
	//dbのオープン
	//変数１,変数２ := sql.Open(ドライバ名,データベース名)
	//dbのクローズ
	//defer DB名.close
	con,er := sql.Open("sqlite3","data.sqlite3")
	if er != nil {
		panic(er)
	}
	defer con.Close()

	q := "select * from mydata"
	rs, er := con.Query(q)
	if er != nil{
		panic(er)
	}
	
	for rs.Next(){
		var md Mydata
		er := rs.Scan(&md.ID,&md.Name,&md.Mail,&md.age)
		if er != nil {
			panic(er)
		}
		fmt.Println(md.Str())
	}

}