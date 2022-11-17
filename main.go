package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"strings"
)

type config struct {
	EdWidget     *widget.Entry
	PreWidget    *widget.RichText
	CurrentFile  fyne.URI
	SaveMenuItem *fyne.MenuItem
}

func (c *config) makeUI() (*widget.Entry, *widget.RichText) {
	edit := widget.NewMultiLineEntry()
	pre := widget.NewRichTextFromMarkdown("")
	c.EdWidget = edit
	c.PreWidget = pre
	edit.OnChanged = pre.ParseMarkdown

	return edit, pre
}

func (c *config) createMenuItems(win fyne.Window) {
	openMenu := fyne.NewMenuItem("Open...", c.Open(win))
	saveMenu := fyne.NewMenuItem("Save", c.Save(win))
	c.SaveMenuItem = saveMenu
	c.SaveMenuItem.Disabled = false
	saveAsMenu := fyne.NewMenuItem("Save as ...", c.SaveAs(win))

	fileMenu := fyne.NewMenu("File", openMenu, saveMenu, saveAsMenu)

	menu := fyne.NewMainMenu(fileMenu)
	win.SetMainMenu(menu)
}

func (c *config) SaveAs(win fyne.Window) func() {
	return func() {
		dia := dialog.NewFileSave(func(closer fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			if closer == nil {
				return
			}

			if c.EdWidget == nil {
				dialog.ShowError(nil, win)
				return
			}

			_, err = closer.Write([]byte(c.EdWidget.Text))
			if err != nil {
				return
			}

			if !strings.HasSuffix(strings.ToLower(closer.URI().String()), ".md") {
				dialog.ShowInformation("Error", "please rename your file with a .md extension!", win)
			}
			c.CurrentFile = closer.URI()
			win.SetTitle(win.Title() + "-" + c.CurrentFile.Name())
			defer func(closer fyne.URIWriteCloser) {
				err := closer.Close()
				if err != nil {

				}
			}(closer)
			c.SaveMenuItem.Disabled = false
		}, win)
		dia.SetFileName("untitled.md")
		dia.SetFilter(FileNameFliter)
		dia.Show()
	}
}

var FileNameFliter = storage.NewExtensionFileFilter([]string{".md", ".MD"})

func (c *config) Open(win fyne.Window) func() {
	return func() {
		dia := dialog.NewFileOpen(func(closer fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			if closer == nil {
				return
			}

			defer func(closer fyne.URIReadCloser) {
				err := closer.Close()
				if err != nil {

				}
			}(closer)

			data, err := ioutil.ReadAll(closer)
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			c.EdWidget.SetText(string(data))
			c.CurrentFile = closer.URI()
			win.SetTitle(win.Title() + "-" + c.CurrentFile.Name())
			c.SaveMenuItem.Disabled = false

		}, win)

		dia.SetFilter(FileNameFliter)
		dia.Show()
	}
}

func (c *config) Save(win fyne.Window) func() {
	return func() {
		if c.CurrentFile != nil {
			writer, err := storage.Writer(c.CurrentFile)
			if err != nil {
				return
			}
			defer func(writer fyne.URIWriteCloser) {
				err := writer.Close()
				if err != nil {

				}
			}(writer)
			writer.Write([]byte(c.EdWidget.Text))
		}
	}
}

var cfg config

func main() {
	//create a app
	a := app.New()

	//create window
	win := a.NewWindow("MarkDown")

	//function
	edit, preview := cfg.makeUI()
	cfg.createMenuItems(win)

	win.SetContent(container.NewHSplit(edit, preview))
	win.Resize(fyne.Size{800, 500})
	win.CenterOnScreen()

	//show
	win.ShowAndRun()

}
