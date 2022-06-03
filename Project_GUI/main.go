package main

// import fyne
import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// label empty
	labelInfo := widget.NewLabel("")
	// new app
	a := app.New()
	wind := a.NewWindow("Access Data")
	wind.Resize(fyne.NewSize(400, 400))

	/////Persoin operation
	title := canvas.NewText("Person Operation", color.White)
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24
	///// entry person
	Card_Id := widget.NewEntry()
	Person_Name := widget.NewEntry()
	Person_Id := widget.NewEntry()
	Person_Age := widget.NewEntry()
	Card_Id.SetPlaceHolder("Enter ID Card")
	Person_Name.SetPlaceHolder("Enter Person Name")
	Person_Id.SetPlaceHolder("Enter Person Id")
	Person_Age.SetPlaceHolder("Enter Person Age")
	///// Insert on database
	buttonIn := widget.NewButton("Insert", func() {
		labelInfo.Text = "Insert"
		labelInfo.Refresh()
	})
	///// update on database
	buttonUp := widget.NewButton(" Update ", func() {
		labelInfo.Text = "Update"
		labelInfo.Refresh()
	})
	///// Select on database
	buttonSel := widget.NewButton(" Select ", func() {
		labelInfo.Text = "Slect"
		labelInfo.Refresh()
	})

	///// Delete on database
	buttonDel := widget.NewButton(" Delete ", func() {
		labelInfo.Text = "Del"
		labelInfo.Refresh()
	})
	//////////////////////////////////////////////Card
	title1 := canvas.NewText("Card Operation", color.White)
	title1.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	title1.Alignment = fyne.TextAlignCenter
	title1.TextSize = 24
	/// entry card
	CardId_id := widget.NewEntry()
	Father_Name := widget.NewEntry()
	number_Phone := widget.NewEntry()
	CardId_id.SetPlaceHolder("Enter ID Card")
	Father_Name.SetPlaceHolder("Enter Father ID")
	number_Phone.SetPlaceHolder("Enter Phone Number")

	///// Insert on database
	buttonInC := widget.NewButton("Insert", func() {
		labelInfo.Text = "InsertCar"
		labelInfo.Refresh()
		
	})
	///// update on database
	buttonUpC := widget.NewButton(" Update ", func() {
		labelInfo.Text = "UpdateCar"
		labelInfo.Refresh()
	})
	///// Select on database
	buttonSelC := widget.NewButton(" Select ", func() {
		labelInfo.Text = "SlectCar"
		labelInfo.Refresh()
	})

	///// Delete on database
	buttonDelC := widget.NewButton(" Delete ", func() {
		labelInfo.Text = "DelCa"
		labelInfo.Refresh()
	})
	// we are almost done
	H1Box := container.New(layout.NewHBoxLayout(), buttonIn, buttonUp, buttonSel, buttonDel)
	H2Box := container.New(layout.NewHBoxLayout(), buttonInC, buttonUpC, buttonSelC, buttonDelC)
	H3Box := container.New(layout.NewVBoxLayout(), title1, CardId_id, Father_Name, number_Phone, H2Box)
	vBox := container.New(layout.NewVBoxLayout(), title, Card_Id, Person_Name, Person_Id, Person_Age, H1Box,
		widget.NewSeparator(), H3Box, labelInfo)

	wind.SetContent(vBox)
	wind.ShowAndRun()
}
