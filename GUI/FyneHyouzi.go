package main
//https://developer.fyne.io/explore/widgets
import "fyne.io/fyne/v2/app"
import "fyne.io/fyne/v2/container"
import "fyne.io/fyne/v2/widget"

func main(){
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	n := widget.NewEntry()
	p := widget.NewPasswordEntry()
	w.SetContent(
		container.NewVBox(
			container.NewAppTabs(
				container.NewTabItem(
					"1",
					l,
				),
				container.NewTabItem(
					"2",
					widget.NewButton("OK",func(){
					l.SetText(n.Text + p.Text)
					}),
				),
				
				
			),
			widget.NewForm(
				widget.NewFormItem("Name",n),
				widget.NewFormItem("Pass",p),
			),
			
			
		),
	)

	w.ShowAndRun()//この一文がないと表示されない（エラーも出ない）
}
