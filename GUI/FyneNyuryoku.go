package main

import "fyne.io/fyne/v2/app"
import "fyne.io/fyne/v2/container"
import "fyne.io/fyne/v2/widget"
import "fyne.io/fyne/v2/theme"
import "strconv"

func main(){
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	l2 := widget.NewLabel("Radio")
	s := widget.NewSlider(0,100)
	l3 := widget.NewLabel("Slider")
	sb := widget.NewButton("Slider",func(){
		l3.SetText(strconv.Itoa(int(s.Value)))
	})//sliderの表示処理
	e := widget.NewEntry()//テキストフィールド
	c := widget.NewCheck("Check",func(f bool){
		if f{
			l.SetText("OK")
		}else{
			l.SetText("NO")
		}
	})//checkBOX(引数は２つで、ラベルと関数が必要)func(f bool)はボタンの状態を表示する
	r := widget.NewRadioGroup(
		[]string{"1","2","3"},
		func(s string){
			if s == ""{
				l2.SetText("No Select")
			}else{
				l2.SetText("selected" + s)
			}
		})//Radioボタン
	r.SetSelected("1")//初期値
	e.SetText("0")
	sl := widget.NewSelect([]string{
		"hi","my","name","is","Selector",
	},func(s string){
		l3.SetText(s)
	})


	w.SetContent(//画面部分
		container.NewVBox(
			l3,
			s,
			sl,
			sb,
			l2,
			r,
			l,
			e,
			c,
			widget.NewButton("Push",func(){
				n, _ := strconv.Atoi(e.Text)
				l.SetText("Total = " + strconv.Itoa(total(n)))
			}),
		),
	)
	//ライトテーマとダークテーマを変更できる（theme）のインポートが必要
	a.Settings().SetTheme(theme.DarkTheme())//LightThemeもある
	w.ShowAndRun()//この一文がないと表示されない（エラーも出ない）
}

func total(n int)int{
	t := 0
	for i := 0; i <= n;i++{
		t += i
	}
	return t
}