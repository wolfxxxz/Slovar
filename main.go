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
var LibraryGob string = "save/library.gob"
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
			var LibraryNewWords model.Slovarick
			LibraryNewWords.TakeTXT(NewWords)
			LibraryWords.Decode(LibraryJson)
			LibraryWords.UpdateLibrary(NewWords)
			LibraryWords.DelDublikat()
			LibraryWords.Encode(LibraryJson)
			LibraryWords.SaveTXT(LibraryTXT)
		case "test":
			var LibraryWords model.Slovarick
			LibraryWords.Decode(LibraryJson)
			LearnSlise := LibraryWords.WorkTest()
			LearnSlise.SaveForLearningTxt(LearnWords)
			LibraryWords.Encode(LibraryJson)
			LibraryWords.SaveTXT(LibraryTXT)
		case "sort":
			var LibraryWords model.Slovarick
			LibraryWords.Decode(LibraryJson)
			LibraryWords.SortLibrary()
			LibraryWords.Encode(LibraryJson)
			LibraryWords.SaveTXT(LibraryTXT)
		case "learn":
			var LibraryLearnWords model.Slovarick
			LibraryLearnWords.TakeTXT(LearnWords)
			LibraryLearnWords.LearnWords()
		case "prepend":
			var LibraryWords model.Slovarick
			var LibraryNewWords model.Slovarick
			LibraryNewWords.TakeTXT(NewWords)
			fmt.Println("len newWords ", len(LibraryNewWords.Words))
			LibraryWords.Decode(LibraryJson)
			fmt.Println("len libJson", len(LibraryWords.Words))
			LibraryWords.CheckAndDelDublikats(&LibraryNewWords)
			fmt.Println("len newWords delDubl ", len(LibraryNewWords.Words))
			LibraryWords.UpdateLibraryOnlyNewWords(&LibraryNewWords)
			model.SaveEmptyTXT(NewWords, "You need to add your words here")
			LibraryWords.Encode(LibraryJson)
			LibraryWords.SaveTXT(LibraryTXT)
		case "resave":
			var LibraryWords model.Slovarick
			//LibraryWords.TakeTXT(LibraryTXT)
			LibraryWords.Decode(LibraryJson)
			LibraryWords.DelDublikat()
			LibraryWords.Encode(LibraryJson)
			LibraryWords.SaveTXT(LibraryTXT)
			LibraryWords.SaveGob(LibraryGob)
			fmt.Println("len", len(LibraryWords.Words))
		case "testgob":
			var LibraryWords model.Slovarick
			LibraryWords.TakeGob(LibraryGob)
			LibraryWords.WorkTest()
			LibraryWords.SaveGob(LibraryGob)
			LibraryWords.SaveTXT(LibraryTXT)
		case "exit":
			fmt.Println("You have to do it, your dream wait")
			return
		}
	}

}
