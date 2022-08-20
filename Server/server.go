package main

import (
	"net/http"
	"log"
	"text/template"
)
func main(){
	pk2()
}

//1ページ
func pk1(){
	//サーバー起動(ホットリロードはない)
	//http.ListenAndServe(アドレス,Handler)
	//http.ListenAndServe("",http.NotFoundHandler())//404
	//http.ListenAndServe("",http.FileServer(http.Dir(".")))//ディレクトリ直下のソースが表示
	//htmlを作成するとそこにいく
	
	/*
	goではサーバー側で処理を実行し、webページの内容を生成して表示する。
	などはできない。
	そのため、HandleFuncを使用する
	http.HandleFunc(アドレス,関数)
	関数を以下のように指定する
	func(w http.ResponseWriter,rq *http.Request){
		実行する処理
	}
	*/
	/*
	msg := `<html><body>
			<h1>Hello</h1>
			<p>Hello,This is Go Server</p>
			</body></html>`
	hh := func(w http.ResponseWriter,rq *http.Request){
		w.Write([]byte(msg))
	}
	http.HandleFunc("/hello",hh)
	http.ListenAndServe("",nil)
	*/

	/*
	Templateとは
	HTMLを元に、データをカスタマイズする
	Templateの使い方
	作成
	変数 := templete.New(名前)
	ソースを突っ込む
	変数１,変数２ := Template.Parse(HTMLのソースコード)
	Webページとして出力する
	変数 := Template.Excute(Writer,interface{})
	*/
	/*
	html := `<html>
	<body>
	<h1>Hello</h1>
	<p>This is sample message</p>
	</body></html>`
	*/
	//tf,er := template.New("index").Parse(html)
	tf,er := template.ParseFiles(`templates/hello.html`)
	if er != nil{
		tf, _ = template.New("index").Parse(`<html>
			<body>
			<h1>Hello</h1>
			<p>No TEMPLATE</p>
			</body></html>`)
	}

	hh := func(w http.ResponseWriter,rq *http.Request){
		er = tf.Execute(w,nil)
		if er != nil{
			log.Fatal(er)
		}
	}
	http.HandleFunc("/hello",hh)
	http.ListenAndServe("",nil)
}

//複数ページの対応
//Temps is template structure
type Temps struct{
	notemp *template.Template
	indx *template.Template
	helo *template.Template
}

//Template for no-template
func notemp() *template.Template{
	src := "<html><body><h1>No Template</h1></body></html>"
	tmp,_ := template.New("index").Parse(src)
	return tmp
}

//setup template function
func setupTemp() *Temps{
	temps := new(Temps)
	temps.notemp = notemp()

	//set index template
	indx,er := template.ParseFiles("templates/index.html")
	if er != nil{
		indx = temps.notemp
	}
	temps.indx = indx

	//set hello template
	helo,er := template.ParseFiles("templates/hello.html")
	if er != nil{
		helo = temps.notemp
	}
	temps.helo = helo
	return temps
}

//index handler
func index(w http.ResponseWriter,rq *http.Request,tmp *template.Template){
	er := tmp.Execute(w,nil)
	if er != nil{
		log.Fatal(er)
	}
}


//hello handler
//var flg bool = true
//var message2 = "メッセージ２だよ"
func hello(w http.ResponseWriter,rq *http.Request,tmp *template.Template){
	/*
	item := struct{
		Title string
		Message string
		Flg bool
		JMessage string
		Items []string
		Message2 string

	}{
		Title:"Send Values",
		Message:"Trueだよ.<br>これはサンプルです。",
		Flg: flg,
		JMessage:"Falseだよ.<br>これはサンプルです。",
		Items: []string{"One","Two","Three"},
		Message2: message2,
	}

	er := tmp.Execute(w,item)//送るもの
	if er != nil{
		log.Fatal(er)
	}
	flg = !flg
	if flg{
		message2 = ""
	}else{
		message2 = "メッセージ２だよ"
	}
	*/

	//パラメータを取得する方法
	//変数 := Request.FormValie(キー)
	msg := "type name and password:"
	if rq.Method == "POST"{
		nm := rq.PostFormValue("name")
		pw := rq.PostFormValue("pass")
		msg = "name: "+nm+",password:"+pw
	}
	item := struct{
		Title string
		Message string
	}{
		Title: "Send Values",
		Message:msg,
	}
	er := tmp.Execute(w,item)
	if er != nil{
		log.Fatal(er)
	}
}

//main部分
func pk2(){
	temps := setupTemp()
	//index handling
	http.HandleFunc("/",func(w http.ResponseWriter, rq *http.Request){
		index(w,rq,temps.indx)
	})
	//hello handling
	http.HandleFunc("/hello",func(w http.ResponseWriter, rq *http.Request){
		hello(w,rq,temps.helo)
	})

	http.ListenAndServe("",nil)
}
