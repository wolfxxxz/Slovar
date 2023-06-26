package model

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Записать слова в .txt файл
func (s *Slovarick) SaveTXT(files string) {
	file, err := os.Create(files)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	for _, v := range s.Words {
		file.WriteString(v.English)
		file.WriteString(" - ")
		file.WriteString(v.Russian)
		if v.Theme != "" {
			file.WriteString(" - ")
			file.WriteString(v.Theme)
		}
		file.WriteString("\n")
	}
}

// Записать пустой .txt файл
func SaveEmptyTXT(files string, txt string) {
	file, err := os.Create(files)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(txt)
}

func (s *Slovarick) SaveForLearningTxt(files string) {
	file, err := os.Create(files)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	if len(s.Words) != 0 {
		for _, v := range s.Words {
			file.WriteString(v.English)
			var length = len(v.English)
			if length <= 30 {
				for i := 0; i+length <= 25; i++ {
					file.WriteString(" ")
				}
			}
			file.WriteString(" - ")
			file.WriteString(v.Russian)
			file.WriteString("\n")
		}
	} else {
		fmt.Println("empty learn words")
	}
}

// Прочитать файл txt и серилиазовать
// Пользуемся strings. для расшифровки слов с .txt
func (s *Slovarick) TakeTXT(filetxt string) {
	data, err := os.ReadFile(filetxt)
	if err != nil {
		fmt.Println(err)
		return
	}
	content := string(data)
	//Делим по \n
	//Получаем массив разбитый по enter
	lines := strings.Split(content, "\n")
	//[dsdsd - sdsdsd - sdsdsd]
	for _, line := range lines {
		//вдруг пустая строка пропустить
		if line == "" {
			continue
		}
		//Делим строку по "-"
		words := strings.Split(line, "-")
		if len(words) <= 1 {
			continue
		}
		for i, v := range words {
			//Пробелы и точки...
			words[i] = strings.TrimSpace(v)
			words[i] = strings.ReplaceAll(words[i], ".", "")
			words[i] = capitalizeFirstRune(words[i])
		}
		id := 0
		theme := ""
		if len(words) > 2 {
			theme = words[2]
		}
		word := NewLibrary(id, words[0], words[1], theme)
		s.Words = append(s.Words, word)
	}
}

// bufio scaner по сути
func ScanStringOne() (string, error) {
	fmt.Print("       ...")
	in := bufio.NewScanner(os.Stdin)
	if in.Scan() {
		return in.Text(), nil
	}
	if err := in.Err(); err != nil {
		return "", err
	}
	return "", nil
}

func (l *Slovarick) Print() {
	for _, v := range l.Words {
		fmt.Print(v.English, " - ", v.Russian, " - ", v.Theme)
		fmt.Println()
	}
}
func capitalizeFirstRune(str string) string {
	runes := []rune(str)
	for i, r := range runes {
		if i == 0 /*|| !unicode.IsLetter(runes[i-1]) */ {
			runes[i] = unicode.ToUpper(r)
		}
	}
	return string(runes)
}

// Прочитать файл txt и серилиазовать
func (s *Slovarick) TakeTXTold(filetxt string) {
	//Такой длинный ридер
	/*
		file, err := os.Open(filetxt)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()
		f := make([]byte, 100024) //Длинна строки
		data2 := []byte{}
		for {
			n, err := file.Read(f)
			if err == io.EOF { // если конец файла
				break // выходим из цикла
			}
			data2 = f[:n]
		}*/
	// капец длинный
	data2, err := os.ReadFile(filetxt)
	if err != nil {
		fmt.Println(err)
	}

	sliseString := []string{}
	dbyte := []byte{}
	for i, v := range data2 {
		b := "-"
		if i < len(data2)-1 {
			if v == 32 && data2[i+1] == 32 {
				continue
			}
		}
		if v == 13 {
			continue
		}

		if v == 10 {
			d := string(dbyte) + b
			if d != "-" {
				sliseString = append(sliseString, d)
				dbyte = []byte{}
			}
		}
		if v == 10 {
			continue
		}
		if v == 46 {
			continue
		}
		dbyte = append(dbyte, v)
	}

	for _, vv := range sliseString {
		SliceThreeString := []string{}
		var Str string

		for _, v := range vv {
			if v == '-' && Str != "" {
				strByte := []byte(Str)
				strByte2 := []byte{}
				for i, v := range strByte {
					if i == 0 && v == 32 {
						continue
					} else if i == len(strByte)-1 && v == 32 {
						continue
					} else {
						strByte2 = append(strByte2, v)
					}
				}
				Str = string(strByte2)
				SliceThreeString = append(SliceThreeString, Str)
				Str = ""
			}
			if v == '-' {
				continue
			}

			Str = Str + string(v)
		}
		if len(SliceThreeString) > 3 {
			SliceThreeString = SliceThreeString[:2]
		}
		for i := 0; len(SliceThreeString) == 2; i++ {
			if len(SliceThreeString) <= 2 {
				SliceThreeString = append(SliceThreeString, "")
			}
		}
		id := 0
		a := NewLibrary(id, SliceThreeString[0], SliceThreeString[1], SliceThreeString[2])
		s.Words = append(s.Words, a)
	}
	//return &Slovarick{SliceLib}
	//10 - начало строки
	//13 - enter
	//46 - точка
	//32 - пробел
	//45 - дефис
}
