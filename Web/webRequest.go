package main

//パッケージで迷ったら
//https://zenn.dev/spiegel/articles/20210223-go-module-aware-mode
import (
	/*
	"fmt"
	"io/ioutil"
	"net/http"
	*/

	"github.com/PuerkitoBio/goquery"
)

func main(){
	p := "https://golang.org"
	/*
	re,er := http.Get(p)//request送信
	if er != nil{
		panic(er)
	}
	defer re.Body.Close()
	
	s,er := ioutil.ReadAll(re.Body)//htmlを読み込む
	if er != nil{
		panic(er)
	}

	fmt.Println(string(s))
	*/
	//goqueryを使う
	doc, er := goquery.NewDocument(p)//documentの作成
	if er != nil{
		panic(er)
	}

	doc.Find("a").Each(func(n int,sel *goquery.Selection){//aタグを探す、１つ目はインデックスが、２つ目はSelectionが渡される
		lk,_ := sel.Attr("href")
		println(n,sel.Text(),"(",lk,")")
	})
}
