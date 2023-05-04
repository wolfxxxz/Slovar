package main

import (
	"fmt"
	"time"

	"github.com/Wolfxxxz/SlovarNV/model"
)

var LibraryNewWords model.Slovarick

var NewWords string = "save/newWords.txt"
var LibraryTXT string = "save/library.txt"
var LibraryJson string = "save/library.json"
var LearnWords string = "save/learning.txt"

func main() {

	for {
		fmt.Println("  Update Library: `update`")
		time.Sleep(1 * time.Second)
		fmt.Println("    Test knowlige: `test'")
		time.Sleep(1 * time.Second)
		fmt.Println("      Sort library: `sort`")
		time.Sleep(1 * time.Second)
		fmt.Println("      Learn Words: `learn`")
		time.Sleep(1 * time.Second)
		fmt.Println("  Exit: `exit`")
		command, _ := model.ScanStringOne()
		switch command {
		case "update":
			var LibraryWords model.Slovarick
			//var NewWordsSlice model.Slovarick
			LibraryWords.Takejson(LibraryJson)
			fmt.Println("Takejson", len(LibraryWords.Words))
			LibraryWords.UpdateLibrary(NewWords)
			fmt.Println("UpdateLibrary", len(LibraryWords.Words))

			//NewWordsSlice.TakeTXT(NewWords)
			//fmt.Println(NewWordsSlice.Words[0])
			//LibraryWords.Words = append(NewWordsSlice.Words, LibraryWords.Words...)

			LibraryWords.DelDublikat()
			fmt.Println("DelDublikat", len(LibraryWords.Words))

			LibraryWords.Savejson(LibraryJson)
			LibraryWords.SaveTXT(LibraryTXT)
		case "test":
			var LibraryWords model.Slovarick
			LibraryWords.Takejson(LibraryJson)
			LibraryWords.WorkTest(LearnWords)
			LibraryWords.Savejson(LibraryJson)
			LibraryWords.SaveTXT(LibraryTXT)
		case "sort":
			var LibraryWords model.Slovarick
			LibraryWords.Takejson(LibraryJson)
			LibraryWords.SortLibrary()
			LibraryWords.Savejson(LibraryJson)
			LibraryWords.SaveTXT(LibraryTXT)
		case "learn":
			var LibraryLearnWords model.Slovarick
			LibraryLearnWords.TakeTXT(LearnWords)
			LibraryLearnWords.LearnWords()
		case "exit":
			fmt.Println("You have to do it, your dream wait")
			return
		}
	}

}
