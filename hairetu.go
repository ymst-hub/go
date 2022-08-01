//配列の書き方
/*
var 変数[整数] 型

文字を分ける方法（stringパッケージのインポートをすること）
var 変数 = string.Split(対象文字列,区切り文字)

配列全てを出力する方法
変数,変数２ := range 配列
変数にはインデックス（何個目か）変数２には値がそれぞれ入る
*/

//スライスの書き方(配列の参照)
/*
var 変数[] 型{値~~,~~}

配列をスライスとして取得する方法
変数 := 配列[開始位置,終了位置]

値を追加する方法(元の配列にも影響がある)
変数 = append(スライス,値~,~~)

//途中への挿入はできない
※その部分までのスライスを取得して、そこにappendをすると間に差し込める
*/
