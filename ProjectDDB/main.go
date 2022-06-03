package main

// import fyne
import (
	"database/sql"
	"fmt"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	cfg := mysql.Config{
		User:   "bot",
		Passwd: "bot",
		Net:    "tcp",
		Addr:   "192.168.1.13:3306",
		DBName: "go",
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	// new app
	GoProject := app.New()

	//////////////////////////////////////////////////////////////////////// Start create Page form
	//// Main menu page
	Menu := GoProject.NewWindow("Menu")
	Menu.Resize(fyne.NewSize(250, 250))
	MenuTitle := canvas.NewText("Main Menu", color.White)
	MenuTitle.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	MenuTitle.Alignment = fyne.TextAlignCenter
	MenuTitle.TextSize = 24

	///// Add table page
	AddTable := GoProject.NewWindow("Add Table")
	AddTable.Resize(fyne.NewSize(400, 400))
	CraeteTable_title := canvas.NewText("Add Table", color.White)
	CraeteTable_title.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	CraeteTable_title.Alignment = fyne.TextAlignCenter
	CraeteTable_title.TextSize = 24

	///// Insert data page
	InsertPage := GoProject.NewWindow("Insert Data")
	InsertPage.Resize(fyne.NewSize(400, 400))
	InsertPage_title := canvas.NewText("Insert Data", color.White)
	InsertPage_title.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	InsertPage_title.Alignment = fyne.TextAlignCenter
	InsertPage_title.TextSize = 24

	///// Update data page
	UpdatePage := GoProject.NewWindow("Update Data")
	UpdatePage.Resize(fyne.NewSize(400, 400))
	UpdatePage_title := canvas.NewText("Update Data", color.White)
	UpdatePage_title.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	UpdatePage_title.Alignment = fyne.TextAlignCenter
	UpdatePage_title.TextSize = 24

	///// Show data page
	ShowPage := GoProject.NewWindow("Show Data")
	ShowPage.Resize(fyne.NewSize(400, 400))
	ShowPage_title := canvas.NewText("Show Data", color.White)
	ShowPage_title.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	ShowPage_title.Alignment = fyne.TextAlignCenter
	ShowPage_title.TextSize = 24

	///// Delete data page
	DeletePage := GoProject.NewWindow("Delete Data")
	DeletePage.Resize(fyne.NewSize(400, 400))
	DeletePage_title := canvas.NewText("Delete Data", color.White)
	DeletePage_title.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	DeletePage_title.Alignment = fyne.TextAlignCenter
	DeletePage_title.TextSize = 24

	/////////////////////////////////////////////////// End Create Page Form  ///////////////////////////////

	//////////////////////////////////////////////////////////////////(Form 6)########## body of Dlete data page
	labelDelCom := widget.NewLabel("")

	Table_NameDel := getTablesName()
	icolDel := 1
	var coulmDelNames []string
	var coulmDelType []string

	labelCoulmDel := widget.NewLabel("Coulm")

	/// choose the table you want Delete the data
	SelectTable_Del := widget.NewSelect(Table_NameDel, func(s string) {
		coulmDelType, coulmDelNames, err = coulmName_DataType(s)
		if err != nil {
			log.Fatal(err)
		}
		labelCoulmDel.Text = coulmDelNames[1]
		labelCoulmDel.Refresh()
	})
	SelectTable_Del.SetSelected("Select table")
	/// choose coulm
	butt_nextCoulmD := widget.NewButton(" Next Coulm ", func() {
		icolDel++
		labelCoulmDel.Text = coulmDelNames[icolDel%len(coulmDelNames)]
		labelCoulmDel.Refresh()
	})
	butt_PrevCoulmD := widget.NewButton(" Previc Coulm ", func() {
		icolDel--
		labelCoulmDel.Text = coulmDelNames[icolDel%len(coulmDelNames)]
		labelCoulmDel.Refresh()
	})
	///// Enter base Delete
	Enter_Name_Del := widget.NewEntry()
	Enter_Name_Del.SetPlaceHolder("Enter Data Delete")
	/////// Finsh delete
	butt_Finsh_Del := widget.NewButton(" Finsh Delete ", func() {
		deleterecord(SelectTable_Del.Selected, coulmDelNames, coulmDelType, Enter_Name_Del.Text, labelCoulmDel.Text)
		coulmDelNames = nil
		coulmDelType = nil
		Enter_Name_Del.SetText("")
		labelCoulmDel.SetText("")
		labelDelCom.Text = "Complet Delete"
		labelDelCom.Refresh()
	})

	/// Button back to main menu
	butt_GoMenu_Del := widget.NewButton(" Back ", func() {
		labelCoulmDel.Text = "Coulm"
		labelCoulmDel.Refresh()
		labelDelCom.Text = ""
		labelDelCom.Refresh()
		DeletePage.Hide()
		Menu.Show()
	})
	HBox_Del0 := container.New(layout.NewHBoxLayout(), butt_PrevCoulmD, labelCoulmDel, butt_nextCoulmD)
	vBox_Del := container.New(layout.NewVBoxLayout(), DeletePage_title, SelectTable_Del, HBox_Del0, Enter_Name_Del,
		butt_Finsh_Del, butt_GoMenu_Del, labelDelCom)
	DeletePage.SetContent(vBox_Del)

	//////////////////////////////////////////////////////////////////(Form 5)########## body of Show data page

	/// Button back to main menu
	butt_GoMenu_Sh := widget.NewButton(" Back ", func() {

		ShowPage.Hide()
		Menu.Show()
	})
	vBox_Sh := container.New(layout.NewVBoxLayout(), ShowPage_title, butt_GoMenu_Sh)
	ShowPage.SetContent(vBox_Sh)

	//////////////////////////////////////////////////////////////////(Form 4)########## body of Update data page
	labelUpCom := widget.NewLabel("")

	var butt_Next_Up *widget.Button
	// lable change
	label_Name_Up := widget.NewLabel("")
	label_type_Up := widget.NewLabel("")
	/// choose the table you want Update the data
	Table_Name := getTablesName()
	var coulmUpNames []string
	var coulmUpTypes []string
	iUp := 1
	icol := 1
	var valuesUp []string
	labelCoulm := widget.NewLabel("Coulm")

	label_UPType := widget.NewLabel(" Type change")
	label_UPType.Refresh()
	SelectTable_UP := widget.NewSelect(Table_Name, func(s string) {
		coulmUpTypes, coulmUpNames, err = coulmName_DataType(s)

		if err != nil {
			log.Fatal(err)
		}
		label_Name_Up.Text = coulmUpNames[1]
		label_type_Up.Text = coulmUpTypes[1]
		labelCoulm.Text = coulmUpNames[1]
		label_UPType.Text = coulmUpTypes[1]
		label_UPType.Refresh()
		labelCoulm.Refresh()
		label_Name_Up.Refresh()
		label_type_Up.Refresh()
		iUp = 2
		butt_Next_Up.Enable()
	})
	SelectTable_UP.SetSelected("Select table")

	/// choose the table you want Update the data

	butt_nextCoulm := widget.NewButton(" Next Coulm ", func() {
		icol++

		labelCoulm.Text = coulmUpNames[icol%len(coulmUpNames)]
		label_UPType.Text = coulmUpTypes[icol%len(coulmUpNames)]
		label_UPType.Refresh()
		labelCoulm.Refresh()
	})
	butt_PrevCoulm := widget.NewButton(" Previc Coulm ", func() {
		icol--
		labelCoulm.Text = coulmUpNames[icol%len(coulmUpNames)]
		label_UPType.Text = coulmUpTypes[icol%len(coulmUpNames)]
		label_UPType.Refresh()
		labelCoulm.Refresh()
	})

	///// Enter base change
	Enter_Name_Up := widget.NewEntry()
	Enter_Name_Up.SetPlaceHolder("Enter base Data change")
	/////////// lable type

	// lable fixed
	label_TextName_Up := widget.NewLabel("  Name  ")
	label_Space_Up := widget.NewLabel("             ")
	label_TextType_Up := widget.NewLabel("  Type  ")
	///// Enter update value
	Enter_UpNamevalue := widget.NewEntry()
	Enter_UpNamevalue.SetPlaceHolder("Enter update Value")
	///// Button Next to Coulm
	butt_Next_Up = widget.NewButton(" Next ", func() {

		valuesUp = append(valuesUp, Enter_UpNamevalue.Text)
		if iUp >= len(coulmUpNames) {
			butt_Next_Up.Disable()
		} else {
			label_Name_Up.Text = coulmUpNames[iUp]
			label_type_Up.Text = coulmUpTypes[iUp]
			label_Name_Up.Refresh()
			label_type_Up.Refresh()
			iUp++
			if iUp == len(coulmUpNames) {
				butt_Next_Up.Disable()
			}
		}
		Enter_UpNamevalue.SetText("")

	})

	///// Button Finsh coulm table
	butt_Finsh_Up := widget.NewButton(" Finsh ", func() {
		valuesUp = append(valuesUp, Enter_UpNamevalue.Text)
		update(SelectTable_UP.Selected, coulmUpNames, coulmUpTypes, valuesUp, Enter_Name_Up.Text, labelCoulm.Text)

		coulmUpNames = nil
		coulmUpTypes = nil
		valuesUp = nil
		labelCoulm.Text = "coulm"
		labelUpCom.Text = "Complet Update"
		Enter_UpNamevalue.Text = ""
		Enter_Name_Up.Text = ""
		Enter_Name_Up.Refresh()
		Enter_UpNamevalue.Refresh()
		labelCoulm.Refresh()
		labelUpCom.Refresh()
	})
	/// Button back to main menu
	butt_GoMenu_Up := widget.NewButton(" Back ", func() {
		labelUpCom.Text = ""
		labelUpCom.Refresh()
		UpdatePage.Hide()
		Menu.Show()
	})
	HBox_Up1 := container.New(layout.NewHBoxLayout(), butt_PrevCoulm, labelCoulm, butt_nextCoulm)
	HBox_Up2 := container.New(layout.NewHBoxLayout(), label_Space_Up, label_TextName_Up, label_Space_Up, label_TextType_Up)
	HBox_Up3 := container.New(layout.NewHBoxLayout(), label_Space_Up, label_Name_Up, label_Space_Up, label_type_Up)
	HBox_Up4 := container.New(layout.NewHBoxLayout(), butt_Next_Up, butt_Finsh_Up)
	vBox_UP := container.New(layout.NewVBoxLayout(), UpdatePage_title, SelectTable_UP, HBox_Up1, label_UPType, Enter_Name_Up,
		HBox_Up2, HBox_Up3, Enter_UpNamevalue, HBox_Up4, butt_GoMenu_Up, labelUpCom)
	UpdatePage.SetContent(vBox_UP)

	//////////////////////////////////////////////////////////////////(Form 3)########## body of Insert data page
	labelInCom := widget.NewLabel("")

	var butt_Next_In *widget.Button
	Tables := getTablesName()
	/// choose the table you want insert the data
	var coulmnNames []string
	var coulmnTypes []string
	i := 1
	var values []string
	// lable change
	label_Name_in := widget.NewLabel("")
	label_type_in := widget.NewLabel("")
	Coulm_Type_In := widget.NewSelect(Tables, func(s string) {
		coulmnTypes, coulmnNames, err = coulmName_DataType(s)
		if err != nil {
			log.Fatal(err)
		}
		label_Name_in.Text = coulmnNames[1]
		label_type_in.Text = coulmnTypes[1]
		label_Name_in.Refresh()
		label_type_in.Refresh()
		i = 2
		butt_Next_In.Enable()
	})
	Coulm_Type_In.SetSelected("Select table")
	// lable fixed
	label_TextName_in := widget.NewLabel("  Name  ")
	label_Space_in := widget.NewLabel("             ")
	label_TextType_in := widget.NewLabel("  Type  ")

	/// textbox inter data in table
	Enter_Name_In := widget.NewEntry()
	Enter_Name_In.SetPlaceHolder("Enter Data")
	///// Button Next to Coulm

	butt_Next_In = widget.NewButton(" Next ", func() {
		values = append(values, Enter_Name_In.Text)
		if i >= len(coulmnNames) {
			butt_Next_In.Disable()
		} else {
			label_Name_in.Text = coulmnNames[i]
			label_type_in.Text = coulmnTypes[i]
			label_Name_in.Refresh()
			label_type_in.Refresh()
			i++
			if i == len(coulmnNames) {
				butt_Next_In.Disable()
			}
		}
		Enter_Name_In.Text = ""
	})

	///// Button Finsh coulm table
	butt_Finsh_In := widget.NewButton(" Finsh ", func() {
		values = append(values, Enter_Name_In.Text)
		insert(Coulm_Type_In.Selected, coulmnNames, coulmnTypes, values)
		labelInCom.Text = "Complet Insert data"
		labelInCom.Refresh()
		Enter_Name_In.SetText("")
		coulmnNames = nil
		coulmnTypes = nil
		values = nil
	})
	/////// Button back to main menu
	butt_GoMenu_In := widget.NewButton(" Back ", func() {
		labelInCom.Text = ""
		labelInCom.Refresh()
		Menu.Show()
		InsertPage.Hide()
	})
	HBox_In1 := container.New(layout.NewHBoxLayout(), label_Space_in, label_TextName_in, label_Space_in, label_TextType_in)
	HBox_In2 := container.New(layout.NewHBoxLayout(), label_Space_in, label_Name_in, label_Space_in, label_type_in)
	HBox_In3 := container.New(layout.NewHBoxLayout(), butt_Next_In, butt_Finsh_In)
	vBox_In := container.New(layout.NewVBoxLayout(), InsertPage_title, Coulm_Type_In, HBox_In1, HBox_In2, Enter_Name_In,
		HBox_In3, butt_GoMenu_In, labelInCom)
	InsertPage.SetContent(vBox_In)

	//////////////////////////////////////////////////////////////////(Form 2)########## body of creat table page
	labelCreateCom := widget.NewLabel("")

	////// Textbox name the table
	Entry_Table_Name := widget.NewEntry()
	Entry_Table_Name.SetPlaceHolder("Enter Teble Name")
	var tname string
	///// Button create table
	butt_CreatTable := widget.NewButton(" Create Table ", func() {
		tname = Entry_Table_Name.Text
	})
	////// text box entry name coulm and choose the type
	coulmHeader := []string{}
	coulmType := []string{}

	Coulm_Name := widget.NewEntry()
	Coulm_Name.SetPlaceHolder("Enter Coulm Name")

	Coulm_Type := widget.NewSelect([]string{"Text", "int", "boolean"}, func(s string) {})
	Coulm_Type.SetSelected("Text")
	///// Button Next to Coulm
	butt_Next := widget.NewButton(" Next ", func() {
		coulmHeader = append(coulmHeader, Coulm_Name.Text)
		coulmType = append(coulmType, Coulm_Type.Selected)
		Coulm_Name.SetText("")
		Coulm_Type.SetSelected("Text")
	})
	///// Button Finsh coulm table
	butt_Finsh := widget.NewButton(" Finsh ", func() {
		labelCreateCom.Text = "Complet create table"
		createtable(tname, coulmHeader, coulmType)
		Coulm_Name.SetText("")
		Coulm_Type.SetSelected("Text")
		Entry_Table_Name.SetText("")
		labelCreateCom.Refresh()
		coulmHeader = nil
		coulmType = nil
	})
	/////// Button back to main menu
	butt_GoMenu_tab := widget.NewButton(" Back ", func() {
		Tables = getTablesName()
		labelCreateCom.Text = ""
		labelCreateCom.Refresh()
		AddTable.Hide()
		Menu.Show()
	})
	///// layout the create table page
	HBox := container.New(layout.NewHBoxLayout(), butt_Next, butt_Finsh)
	vBox := container.New(layout.NewVBoxLayout(), CraeteTable_title, Entry_Table_Name, butt_CreatTable,
		widget.NewSeparator(), Coulm_Name, Coulm_Type, HBox, butt_GoMenu_tab, labelCreateCom)
	AddTable.SetContent(vBox)

	//////////////////////////////////////////////////////////////////(Form 1)########## body of main menu page
	///// button crete table in main menu page
	CreateForm := widget.NewButton(" Create Table", func() {
		AddTable.Show()
		Menu.Hide()
	})
	///// button Insert data table in main menu page
	butt_Insert := widget.NewButton("  Insert Data  ", func() {
		//	Tables = getTablesName()
		InsertPage.Show()
		Menu.Hide()
	})
	///// button update data table in main menu page
	butt_update := widget.NewButton("  Update Data  ", func() {
		UpdatePage.Show()
		Menu.Hide()
	})
	///// button Show data table in main menu page
	butt_Show := widget.NewButton("  Show Data   ", func() {
		ShowPage.Show()
		Menu.Hide()
	})
	///// button delete data table in main menu page
	butt_delete := widget.NewButton("  Delete Data   ", func() {
		DeletePage.Show()
		Menu.Hide()
	})
	//// layout main menu page
	HBox1 := container.New(layout.NewHBoxLayout(), butt_Insert, butt_update)
	HBox2 := container.New(layout.NewHBoxLayout(), butt_Show, butt_delete)
	MenuBoxH := container.New(layout.NewVBoxLayout(), MenuTitle, CreateForm, HBox1, HBox2)
	Menu.SetContent(MenuBoxH)
	Menu.ShowAndRun()

}

///// function create table used on button finsh in create table page
func createtable(tname string, a []string, b []string) (bool, error) {
	x := "CREATE TABLE `go`.`" + tname + "`(`id` INT NOT NULL, "
	print(x)
	i := 0
	for i < len(a) {
		if a[i] == "" {
			a[i] = fmt.Sprint("x", i)
		}
		x += "`" + a[i] + "` " + b[i] + " null, "
		i++
	}
	x += "PRIMARY KEY (`id`));"
	print(x)
	_, err := db.Exec(x)
	if err != nil {
		return false, fmt.Errorf("%v", err)
	}
	return true, nil
}

func coulmName_DataType(tname string) ([]string, []string, error) {
	x := "SELECT DATA_TYPE from INFORMATION_SCHEMA.COLUMNS where   table_name = '" + tname + "';"

	rows, err := db.Query(x)

	if err != nil {
		log.Fatal(err)
	}
	var str string
	var DATA_TYPE []string
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&str)
		if err != nil {
			log.Fatal(err)
		}
		DATA_TYPE = append(DATA_TYPE, str)
	}

	y := "SELECT COLUMN_NAME FROM INFORMATION_SCHEMA.COLUMNS 	where TABLE_NAME = '" + tname + "';"
	NameCoulmn, err := db.Query(y)
	if err != nil {
		log.Fatal(err)
	}
	var CoulmName []string
	defer NameCoulmn.Close()
	for NameCoulmn.Next() {
		err := NameCoulmn.Scan(&str)
		if err != nil {
			log.Fatal(err)
		}
		CoulmName = append(CoulmName, str)
	}
	return DATA_TYPE, CoulmName, nil
}

func getTablesName() []string {
	x := "SELECT TABLE_NAME FROM INFORMATION_SCHEMA.COLUMNS 	where TABLE_SCHEMA = 'go' group by TABLE_NAME;"

	rows, err := db.Query(x)
	if err != nil {
		log.Fatal(err)
	}
	var str string
	var Table_Name []string
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&str)
		if err != nil {
			log.Fatal(err)
		}
		Table_Name = append(Table_Name, str)
	}
	return Table_Name
}

func insert(tname string, a []string, b []string, c []string) {
	x := "INSERT INTO `go`.`" + tname + "`("
	valestring := ""
	i := 1
	for i < len(a) {
		x += "`" + a[i] + "`,"
		if b[i] == "text" {
			valestring = valestring + ",'" + c[i-1] + "'"
		} else {
			valestring = valestring + "," + c[i-1]
		}
		i++
	}
	valestring = valestring[1:]
	x = x[:(len(x) - 1)]
	x += ") VALUES (" + valestring + ");"
	_, err := db.Exec(x)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(x)
}

func update(tname string, a []string, b []string, c []string, wherevale string, wherename string) {

	println(c)
	x := "UPDATE `go`.`" + tname + "`SET"
	i := 1
	z := ""
	for i < len(a) {
		if c[i-1] != "" {
			if b[i] == "text" {
				x += "`" + a[i] + "`= '" + c[i-1] + "',"
			} else {
				x += "`" + a[i] + "`= " + c[i-1] + ","
			}
		}
		if wherename == a[i] {
			z = b[i]
		}
		i++
	}
	x = x[:(len(x) - 1)]
	x += " WHERE `" + wherename + "` = "
	if z == "text" {
		x += "'" + wherevale + "';"
	} else {
		x += wherevale + ";"
	}
	fmt.Println(x)
	_, err := db.Exec(x)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(x)
}

func deleterecord(tname string, a []string, b []string, wherevale string, wherename string) {
	x := "DELETE FROM `go`.`" + tname + "`"
	i := 1
	z := ""
	for i < len(b) {
		if wherename == a[i] {
			z = b[i]
		}
		i++
	}
	x += " WHERE `" + wherename + "` = "
	if z == "text" {
		x += "'" + wherevale + "';"
	} else {
		x += wherevale + ";"
	}
	_, err := db.Exec(x)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(x)
}