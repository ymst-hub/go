package main

import(
	"os"
	"fmt"
	//"io/ioutil"
	"bufio"
)
func main(){
	rt := func(f *os.File){
		/*
		s,er := ioutil.ReadAll(f)
		if er != nil{
			panic(er)
		}
		fmt.Println(string(s))//sそのままだと数値配列(おそらく文字コード)が表示される
		*/

		//一行ずつ読み込む
		r := bufio.NewReaderSize(f,4096)
		for i := 1;true;i++{
			s,_,er := r.ReadLine()
			if er != nil{
				break
			}
			fmt.Println(i,":",string(s))
		}
	}
	fn := "data.txt"

	f,er := os.OpenFile(fn,os.O_RDONLY,os.ModePerm)
	if er != nil{
		panic(er)
	}
	defer f.Close()

	fmt.Println("<<<start>>>")
	rt(f)
	fmt.Println("<<<end>>>")
}

//ファイル情報を調べる
//変数１,変数２ := os.Stat(パス)//ファイルインフォが１に、エラーなら２に入る
//特定のディレクトリを調べる
//変数 := ioutil.ReadDir(パス)//ファイルインフォをまとめて調べられる
/*
FileInfoの情報
Name()：ファイル名 => string
Size():ファイルサイズ => int64
Mode():ファイルモード => FileMode
ModTime():更新日時 => Time
IsDir():ディレクトリかどうか => bool(ディレクトリならtrue)
Sys():プロセスに関する情報を返す
*/
