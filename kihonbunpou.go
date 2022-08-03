package main

import "strconv"

func main(){
	//整数
	var a int//8~64まである
	var b uint//上記と同じく、符号なし
	var c uintptr//ポインタ
	var byte byte//uint8の別名
	var rune rune//int32の別名
	var d float32//64もある。実数
	var hukusosuu complex64//実数部と虚数部をそれぞれ32ずつ記載可能128もある
	var bool bool//真偽
	var text string//文字
	//runeはchar型、stringはchar配列''と””で異なる
	a = -10
	b = 1 + 4
	c = uintptr(a)//型変換はこんな
	//暗黙の型変換はやってくれないため、自身で記載する
	//型名は省略も可
	x,y,z := 100 ,200 ,300
	//:=はvarを省略して、代入するよという明示
	n,err := strconv.Atoi(text)
	//AtoiとItoaはCと同様に使える
	//nは型変換した結果、エラーになった場合はerrの場所に格納される（nullか、値があるか）
	const teisu
	//constで定数を宣言する

	/*
	if文
	if 条件 {
		処理
	}else{

	}
	*/

	//比較演算子はCと同様

	/*
	switch 条件{
	case 値:
		処理
		//fallthrough
		//を記載すると、breakを記載しない場合のswitchの動きができる
	case 値２:
		処理
	}

	短く下記にもできる
	switch 文;条件{

	}
	例)
	swicth n,err = strconv.Atoi(string);n{
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	}

	条件分がない場合
	switch{
	case n=="/":
		fmt.Println("ルート")
	case n=="/src"
		fmt.Println("メイン画面")
	}
	これはtrueの条件ケースを探して、当てはまるときに実行させられる
	*/
	
	/*
	for 条件{
		処理
	}

	また、下記の書き方も可能
	for 初期化;条件;後処理{
		処理
		//continue 繰り返しをぬけ、最初に戻る
		//break 繰り返しをぬけ、forの次の処理へ行く
	}
	*/

	/*
	goto構文
	goto ラベル名
	~~~~~~~
	ラベル名:
		処理
	上記のようにすると、そこに飛ぶ
	goto err

	err:
		fmt.Println("Error")
	みたいにエラー処理に使う
	ただ、使い方を考えないとスパゲティコードになる
	*/

	/*
	ポインタ
	var 変数名　*型
	var p *int
	n = 10
	p := &n//＆をつけてメモリを代入する

	*p = 1
	とするとnも1になる
	*/
	
	/*
	構造体
	struct{
		変数 型
	}
	var person struct{
		Name string
		year int
	}
	上記のように変数にも突っ込める

	func (p person)PrintData(){
		fmt.Println(p.name)
	}
	のようにして、構造体のメソッドも記載できる

	hoge := new(person)で新しい構造体の定義ができる
	配列、スライス、マップ、チャンネルの場合は
	makeを使用する

	インターフェイス
	//下記のメソッドをプログラムに必ず実装してねという規範
	type 名前 interface{
		メソッド１
		メソッド２
	}
	インターフェイスは型として使用でき、
	neko
	inu
	を動物で括るみたいなことができる
	*/

	/*
	並行処理
	go 関数()
	でmain処理と並列して行うようになる

	チャンネル
	値の受け渡しに使う（スレッド間）
	変数 := make(chan 型)
	値の取り出し
	変数 := <-チャンネル
	値の追加
	チャンネル <- 値

	チャンネルごとに処理を行う
	select{
	case re := <- channel1:
		チャンネル１に値が送られた場合の処理
	}
	というように使う

	Mutexパッケージを使うと、
	.Lock()
	.Unlockで排他制御も可能
	*/

}
