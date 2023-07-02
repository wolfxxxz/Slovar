package model

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type Statistick struct {
	Data        string  `json:"data"`
	RightAnswer int     `json:"right"`
	WrongAnswer int     `json:"wrong"`
	Average     float64 `json:"average"`
}

func NewStatistick(NewRight, NewWrong int) *Statistick {
	NewData := time.Stamp
	NewAverage := float64(NewWrong) / float64(NewRight) * 100
	return &Statistick{Data: NewData, RightAnswer: NewRight, WrongAnswer: NewWrong, Average: NewAverage}
}

func (stat Statistick) Println() {
	fmt.Printf("Right Answer: %v/n Wrong Answer: %v/n Your Average is: %v/n", stat.RightAnswer, stat.WrongAnswer, stat.Average)
	fmt.Println(stat.Data)
}

// Open
func (SliceWord *Statistick) Takejson(file string) {
	//1. Создадим файл дескриптор
	filejson, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer filejson.Close()
	//fmt.Println("File descriptor successfully created!")

	//2. Теперь десериализуем содержимое jsonFile в экземпляр Go
	f := make([]byte, 64)
	var data2 string

	for {
		n, err := filejson.Read(f)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		data2 = data2 + string(f[:n])
	}

	data := []byte(data2)

	// Теперь задача - перенести все из data в users - это и есть десериализация!
	json.Unmarshal(data, &SliceWord)

}

// Сохранить в json file
// Marshal
func (s *Statistick) Savejson(file string) {

	byteArr, err := json.MarshalIndent(s, "", "   ")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(byteArr))             0664
	err = os.WriteFile(file, byteArr, 0666) //-rw-rw-rw-
	if err != nil {
		log.Fatal(err)
	}
}

/*
func main() {
	new := NewStatistick(10, 5)
	fmt.Println(new)
	new.Savejson("statistik.json")
}*/
