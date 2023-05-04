package model

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Достать новые слова из .txt файла
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
		file.WriteString("\n")
	}
}

func (s *Slovarick) SaveForLearningTxt(files string) {
	file, err := os.Create(files)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
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
}

// Часть 1 Добавление новых слов в библиотеку Загрузкой с файла txt
func (s *Slovarick) TakeTXT(filetxt string) {
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

// Часть 2 Добавление новых слов в библиотеку Загрузкой с файла txt
func (oldWords *Slovarick) UpdateLibrary(filetxt string) {
	var NewWords Slovarick
	NewWords.TakeTXT(filetxt)
	fmt.Println(NewWords.Words[0])
	c := len(oldWords.Words)
	//--------Соединяем два среза в один--------------
	oldWords.Words = append(NewWords.Words, oldWords.Words...)
	//NewWords.Words = append(NewWords.Words, oldWords.Words...)

	d := len(NewWords.Words)
	// Записать в filetxt пустой
	var ZeroWords Slovarick
	NewsZero := []Word{}
	NewsZero = append(NewsZero, Word{English: ""})
	ZeroWords.Words = append(ZeroWords.Words, &NewsZero[0])
	ZeroWords.SaveTXT(filetxt)

	if d != c {
		fmt.Println("                   New Words Add:", c-d)
	} else {
		fmt.Println("Для загрузки слов списком необходимо упорядочить и вставить слова в файл `save/newWords.txt`")
		fmt.Println("english - перевод - тема")
		fmt.Println("в конце оставить пустую строчку")
		fmt.Println("I believe in you!!!")
	}
}

// Удалить дубликаты
func (s *Slovarick) DelDublikat() {
	var count = 0
	var count2 = 0
	for i := 0; i <= len(s.Words)-1; i++ {
		for _, v := range s.Words {
			if strings.EqualFold(v.English, s.Words[i].English) {
				count++

			}
			if count == 2 {

				s.Words[i] = v
				count = 0
			}
		}
	}

	s.ReverseSlice()
	var withoutDublicat []*Word //:= s.Words
	for ii, v := range s.Words {
		var count1 int
		for i := ii; i <= len(s.Words)-1; i++ {
			if strings.EqualFold(v.English, s.Words[i].English) {
				count1++
			}
		}
		if count1 == 1 {

			withoutDublicat = append(withoutDublicat, v)
			count1 = 0

		} else {
			count2++

			count1 = 0
		}
	}

	*s = *NewSlovarick(withoutDublicat)
	s.ReverseSlice()

	fmt.Println(len(withoutDublicat))
	fmt.Println(count2)
}

// bufio scaner по сути
func ScanStringOne() (string, error) {
	fmt.Print("    ...")
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
