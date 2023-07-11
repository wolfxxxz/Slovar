package model

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

// Unmarshal (Декодирование)
func (SliceWord *Slovarick) Decode(file string) {
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
// Marshal Кодирование
func (s *Slovarick) Encode(file string) {

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

// Если хочется поднять первую букву
func (s *Slovarick) ReSavejson(file string) {
	for i, v := range s.Words {
		s.Words[i].English = capitalizeFirstRune(v.English)
		s.Words[i].Russian = capitalizeFirstRune(v.Russian)
		s.Words[i].Theme = capitalizeFirstRune(v.Theme)
	}

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
func TakejsonMap(slice []*Word) (mapString map[string]string) {
	for _, v := range slice {
		mapString[v.English] = v.Russian
	}

	return
}*/
