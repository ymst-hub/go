package main

import (

)

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
	

}
