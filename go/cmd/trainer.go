package main

import (
	"fmt"
	"io/ioutil"
	"jpkana"
	"os"
	"strconv"
	"strings"

	"github.com/akamensky/argparse"

	"github.com/rivo/tview"
)

const (
	windowTitle         string = "Kanatrainer"
	hiraganaLabel       string = jpkana.HIRAGANA
	katakanaLabel       string = jpkana.KATAKANA
	kanaLabelText       string = "Choose kana"
	lengthLabelText     string = "Choose length"
	difficultyLabelText string = "Choose difficulty"
	romanjiLabelText    string = "romanji"
	verifyButtonText    string = "verify"
	changeButtonText    string = "change"
	exitButtonText      string = "exit"
	continueButtonText  string = "continue"
	okMessage           string = "This is correct ! Well done !"
	koMessage           string = "Wrong ! %s is %s, you wrote %s"
)

var allowedChar map[rune]struct{} = map[rune]struct{}{
	'a': struct{}{}, 'b': struct{}{}, 'c': struct{}{}, 'd': struct{}{}, 'e': struct{}{},
	'f': struct{}{}, 'g': struct{}{}, 'h': struct{}{}, 'i': struct{}{}, 'j': struct{}{},
	'k': struct{}{}, 'l': struct{}{}, 'm': struct{}{}, 'n': struct{}{}, 'o': struct{}{},
	'p': struct{}{}, 'q': struct{}{}, 'r': struct{}{}, 's': struct{}{}, 't': struct{}{},
	'u': struct{}{}, 'v': struct{}{}, 'w': struct{}{}, 'x': struct{}{}, 'y': struct{}{},
	'z': struct{}{},
}

func main() {
	parser := argparse.NewParser("trainer", "Japanese Kana trainer")
	parser.DisableHelp()

	hiraganaFilename := parser.String("h", "hiragana", &argparse.Options{Required: true, Help: "Hiragana resource"})
	katakanaFilename := parser.String("k", "katakana", &argparse.Options{Required: true, Help: "Katakana resource"})

	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		return
	}

	hbytes, err := ioutil.ReadFile(*hiraganaFilename)
	if err != nil {
		fmt.Println(err)
		return
	}
	kbytes, err := ioutil.ReadFile(*katakanaFilename)
	if err != nil {
		fmt.Println(err)
		return
	}
	jpGen, err := jpkana.New(hbytes, kbytes)
	if err != nil {
		fmt.Println(err)
		return
	}
	var choosenKana string = hiraganaLabel
	var length uint64 = 3
	var difficulty uint64 = 3

	app := tview.NewApplication()

	romanjiEntry := tview.NewInputField().
		SetLabel(romanjiLabelText)
	romanjiEntry.SetAcceptanceFunc(func(value string, lastChar rune) bool {
		_, ok := allowedChar[lastChar]
		return ok
	})

	kanaEntry := tview.NewInputField().
		SetLabel(hiraganaLabel)
	kanaSelector := tview.NewDropDown().
		SetLabel(kanaLabelText).
		SetOptions([]string{hiraganaLabel, katakanaLabel}, func(value string, index int) {
			choosenKana = value
			kanaEntry.SetLabel(choosenKana)
		}).
		SetCurrentOption(0)

	lengthSelector := tview.NewDropDown().
		SetLabel(lengthLabelText).
		SetOptions([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}, func(value string, index int) {
			if length, err = strconv.ParseUint(value, 10, 32); err != nil {
				//TODO error
				fmt.Printf("error length")
				return
			}
		}).
		SetCurrentOption(2)

	difficultySelector := tview.NewDropDown().
		SetLabel(difficultyLabelText).
		SetOptions([]string{"1", "2", "3", "4", "5"}, func(value string, index int) {
			if difficulty, err = strconv.ParseUint(value, 10, 32); err != nil {
				//TODO error
				fmt.Printf("error difficulty")
				return
			}
		}).
		SetCurrentOption(3)

	var kana string
	var roma string

	generate := func() {
		kana, roma = jpGen.Generate(choosenKana, uint(length), uint(difficulty))
		kanaEntry.SetText(kana)
		romanjiEntry.SetText("")
	}
	generate()

	form := tview.NewForm().
		AddFormItem(kanaSelector).
		AddFormItem(lengthSelector).
		AddFormItem(difficultySelector).
		AddFormItem(kanaEntry).
		AddFormItem(romanjiEntry)
	form.
		AddButton(verifyButtonText, func() {
			romanjiEntry.SetText(strings.ToLower(romanjiEntry.GetText()))
			if roma == romanjiEntry.GetText() {
				w := tview.NewModal().
					SetText(okMessage).
					AddButtons([]string{continueButtonText}).
					SetDoneFunc(func(buttonIndex int, buttonLabel string) {
						app.SetRoot(form, true)
					})
				generate()
				app.SetRoot(w, false)
			} else {
				w := tview.NewModal().
					SetText(fmt.Sprintf(koMessage, kana, roma, romanjiEntry.GetText())).
					AddButtons([]string{continueButtonText}).
					SetDoneFunc(func(buttonIndex int, buttonLabel string) {
						app.SetRoot(form, true)
					})
				generate()
				app.SetRoot(w, false)
			}
			generate()
		}).
		AddButton(changeButtonText, func() {
			generate()
		}).
		AddButton(exitButtonText, func() {
			app.Stop()
		})

	app.SetRoot(form, true).EnableMouse(true).SetFocus(romanjiEntry).Run()
}
