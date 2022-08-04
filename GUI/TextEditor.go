package main


//go build -ldflags="-H windowsgui" TextEditor.go
import "fyne.io/fyne/v2/app"
import "fyne.io/fyne/v2/widget"
import "fyne.io/fyne/v2/container"
import "fyne.io/fyne/v2/dialog"
import "fyne.io/fyne/v2"
import "fyne.io/fyne/v2/layout"
import "fyne.io/fyne/v2/theme"
import "io/ioutil"
import "os"

func main(){
	a := app.New()
	w := a.NewWindow("Text")
	edit := widget.NewEntry()
	edit.MultiLine = true
	sc := container.NewScroll(edit)
	inf := widget.NewLabel("infomation bar")

	//新規ファイル
	nf := func(){
		dialog.ShowConfirm("Alert","Create New document?",func(f bool){
			if f{
				edit.SetText("")
				inf.SetText("Create new document")
			}
		},w)
	}

	//ファイルを開く
	of := func(){
		f := widget.NewEntry()
		dialog.ShowCustomConfirm("Open file name","OK","Cancel",f,
			func(b bool){
				if b{
					fn := f.Text + ".txt"
					ba,er := ioutil.ReadFile(fn)
					if er != nil{
						dialog.ShowError(er,w)
					}else{
						edit.SetText(string(ba))
						inf.SetText("Open from file" + fn)
					}
				}
			},w)
	}

	//ファイル保存
	sf := func (){
		f := widget.NewEntry()
		dialog.ShowCustomConfirm("Save file name","OK","Cancel",f,func(b bool){
			if b{
				fn := f.Text + ".txt"
				er := ioutil.WriteFile(fn,[]byte(edit.Text),os.ModePerm)
				if er != nil {
					dialog.ShowError(er,w)
					return
				}
				inf.SetText("Save to file" + fn)
			}
		},w)
	}

	//ファイルを閉じる
	qf := func(){
		dialog.ShowConfirm("Alert","Quit application?",func(b bool){
			if b{
				a.Quit()
			}
		},w)
	}
	

	//テーマを変える
	tf := true
	cf := func(){
		if tf{
			a.Settings().SetTheme(theme.LightTheme())
			inf.SetText("change to Light Theme")
		}else{
			a.Settings().SetTheme(theme.DarkTheme())
			inf.SetText("change to Dark Theme")
		}
		tf = !tf
	}

	//メニューバー
	createMenubar := func() *fyne.MainMenu{
		return fyne.NewMainMenu(
			fyne.NewMenu("File",
				fyne.NewMenuItem("New",func(){
					nf()
				}),
				fyne.NewMenuItem("Open",func(){
					of()
				}),
				fyne.NewMenuItem("Save",func(){
					sf()
				}),
				fyne.NewMenuItem("Change Theme",func(){
					cf()
				}),
				fyne.NewMenuItem("Quit",func(){
					qf()
				}),
			),
			fyne.NewMenu("Edit",
				fyne.NewMenuItem("Cut",func(){
					edit.TypedShortcut(
						&fyne.ShortcutCut{
							Clipboard: w.Clipboard()})
					inf.SetText("Cut text")
				}),
				fyne.NewMenuItem("Copy",func(){
					edit.TypedShortcut(
						&fyne.ShortcutCopy{
							Clipboard: w.Clipboard()})
					inf.SetText("Copy text")
				}),
				fyne.NewMenuItem("Paste",func(){
					edit.TypedShortcut(
						&fyne.ShortcutPaste{
							Clipboard: w.Clipboard()})
					inf.SetText("Paste text")
				}),
			),
		)
	}

	//ツールバー
	createToolbar := func() *widget.Toolbar{
		return widget.NewToolbar(
			widget.NewToolbarAction(
				theme.DocumentCreateIcon(),func(){
					nf()
				}),
			widget.NewToolbarAction(
				theme.FolderOpenIcon(),func(){
					of()
				}),
			widget.NewToolbarAction(
				theme.DocumentSaveIcon(),func(){
					sf()
				}),
		)
	}

	mb := createMenubar()
	tb := createToolbar()

	w.SetMainMenu(mb)
	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewBorderLayout(
				tb,inf,nil,nil,
			),
			tb,inf,sc,
		),
	)
	w.Resize(fyne.NewSize(500,500))
	w.ShowAndRun()




}

