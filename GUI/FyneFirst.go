package main

//パッケージの利用方法
//1,import文を記載する
//2,go mod init モジュール名でgo.modファイルを作成する
//3,go mod tidy
//4,go install （？）
//5,go run ファイル名
import "fyne.io/fyne/v2/app"
import "fyne.io/fyne/v2/container"
import "fyne.io/fyne/v2/widget"
import "strconv"

func main(){
	c := 0
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	w.SetContent(
		container.NewVBox(
			l,
			widget.NewButton("Push",func(){
				c++
				l.SetText("Count = " + strconv.Itoa(c))
			}),
		),
	)

	w.ShowAndRun()//この一文がないと表示されない（エラーも出ない）
}

/*
//パッケージのサンプルコード
func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()
}
*/
