package main

import (
	"fmt"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"os"
	"strconv"
)

var MainMenu *ui.Window


func main() {
	if err := ui.Main(menu); err!= nil {
		panic(err)
	}
}

func stringToInt(str string) int {
    i, err := strconv.Atoi(str)
    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }
    return i
}

func menu() {
	MainMenu = ui.NewWindow("Test UI", 300, 500, true)
	MainMenu.OnClosing(func (*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		MainMenu.Destroy()
		return true
	})

	tab := ui.NewTab()

	MainMenu.SetChild(tab)
	MainMenu.SetMargined(true)

	tab.Append("Calculator", CalculatorUI())
	tab.SetMargined(0, true)

	MainMenu.Show()

}

func CalculatorUI() ui.Control {

	var anwserEntry *ui.Entry
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	vbox.Append(hbox, false)

	vbox.Append(ui.NewLabel(""), false)

	group := ui.NewGroup("")
	group.SetMargined(true)
	vbox.Append(group, true)

	group.SetChild(ui.NewNonWrappingMultilineEntry())

	CalForm := ui.NewForm()
	CalForm.SetPadded(true)
	group.SetChild(CalForm)

	entryNum1 := ui.NewEntry()
	entryNum1.OnChanged(func(entry *ui.Entry) {
	})
	CalForm.Append("First Number:", entryNum1, false)

	simbol := ui.NewCombobox()
	simbol.Append("+")
	simbol.Append("-")
	simbol.Append("*")
	simbol.Append("/")
	simbol.OnSelected(func(combobox *ui.Combobox) {
	})
	CalForm.Append("simbol:", simbol, false)

	entryNum2 := ui.NewEntry()
	entryNum2.OnChanged(func(entry *ui.Entry) {
	})
	CalForm.Append("Second Number:", entryNum2, false)

	gimmeAnwser := ui.NewButton("gimme anwser >:D")
	CalForm.Append("", gimmeAnwser, false)
	gimmeAnwser.OnClicked(func(button *ui.Button) {
		number1 := entryNum1.Text()
		number2 := entryNum2.Text()
		var nb1 int = stringToInt(number1)
		var nb2 int = stringToInt(number2)
		anwser := 0
		switch simbol.Selected() {
		case 0:
			anwser = nb1 + nb2
		case 1:
			anwser = nb1 - nb2
		case 2:
			anwser = nb1 * nb2
		case 3:
			anwser = nb1 / nb2
		}
		anwser1 := strconv.Itoa(anwser)
		fmt.Println(anwser)
		anwserEntry.SetText(anwser1)
	})

	anwserEntry = ui.NewEntry()
	anwserEntry.SetReadOnly(true)
	CalForm.Append("Anwser:", anwserEntry, false)

	return vbox
}