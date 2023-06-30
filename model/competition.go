package model

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/agnivade/levenshtein"
)

func (s *Slovarick) WorkTest() *Slovarick {
	startTime := time.Now()

	//Инициализирую мапу map[rus][]
	K := s.CreateAndInitMapWords()
	//Тут будут слова которые потом нужно учить
	LearnSlice := NewSlovarick(nil)

	fmt.Println("Количество слов для теста")
	//количество слов для теста и скопировать их с оригинала
	quantity := ScanInt()
	var capSlovar int = quantity // cap  ёмкость нового []Words
	NewWordsu := make([]*Word, capSlovar)
	NewSlovar := NewSlovarick(NewWordsu)
	copy(NewSlovar.Words, s.Words[:quantity])

	//Отрезать этот кусок от оригинала
	s.Words = s.Words[quantity:]
	//Присобачить что то впереди, что то сзади
	fmt.Println("                     START")
	var yes int
	var not int
	var exit1 bool
	for _, v := range NewSlovar.Words {
		if v == nil {
			log.Println("WorkTest err")
			break
		}
		if exit1 {
			not++
			s.Preppend(v)
			LearnSlice.Preppend(v)
			continue
		}

		y, n, exit := Compare(*v, K)
		if exit {
			exit1 = true
			n = 1
		}
		if y > 0 && n > 0 {
			yes++
			s.Preppend(v)
			LearnSlice.Preppend(v)
			continue
		}

		if y > 0 {
			yes++
			v.RightAswer += 1
			s.AppendWord(v)
		} else if n > 0 {
			not++
			s.Preppend(v)
			LearnSlice.Preppend(v)
			//
		} else {
			break
		}
	}
	duration := time.Since(startTime)
	PrintTime(duration)

	fmt.Println(yes, not)
	return LearnSlice
}

// Учить слова которые в тесте не смог выучить
func (s Slovarick) LearnWords() {
	// Измерить время выполнения
	startTime := time.Now()

	fmt.Println("                 Learn Words")
	K := s.CreateAndInitMapWords()
	for {
		if len(s.Words) == 0 {
			break
		}
		v := s.Words[0]
		if v == nil {
			break
		}
		y, _, exit := Compare(*v, K)
		if exit {
			break
		}

		if y > 0 && len(s.Words) != 1 {
			s.Words = s.Words[1:]
		} else if y < 1 {
			copy(s.Words, s.Words[1:])
			s.Words[len(s.Words)-1] = v
			PrintXpen(v.English)
		} else if y > 0 && len(s.Words) == 1 {
			break
		}
	}
	duration := time.Since(startTime)
	PrintTime(duration)
}

func ScanInt() (n int) {
	for {
		cc, _ := ScanStringOne()
		i, err := strconv.Atoi(cc)
		if err != nil {
			fmt.Println("Incorect, please enter number")
		} else {
			n = i
			break
		}
	}
	return
}

// Сравнение строк / пробелы между словами "_"
func Compare(l Word, mapWord *map[string][]string) (yes int, not int, exit bool) {
	fmt.Println(l.Russian, " ||Тема: ", l.Theme)
	c := IgnorSpace(l.English)

	a, _ := ScanStringOne()
	if a == "exit" {
		exit = true
		//yes = 1
		return yes, not, exit
	}
	s := IgnorSpace(a)

	if strings.EqualFold(c, s) {
		yes++
		fmt.Println("Yes")
	} else if CompareWithMap(l.Russian, s, mapWord) {
		yes++
		not++
		fmt.Println("Не совсем правильно ", l.English)
	} else if compareStringsLevenshtein(c, s) {
		yes++
		fmt.Println("Не совсем правильно ", l.English)
	} else {
		not++
		fmt.Println("Incorect:", l.English)
		for {
			fmt.Println("Please enter correct: ")
			j, _ := ScanStringOne()
			jj := IgnorSpace(j)
			if strings.EqualFold(c, jj) {
				break
			}
		}
	}
	return yes, not, false
}

func compareStringsLevenshtein(str1, str2 string) bool {
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)
	//number of allowed errors
	mistakes := 1

	if distance := levenshtein.ComputeDistance(str1, str2); distance <= mistakes {
		return true
	} else {
		return false
	}
}

func IgnorSpace(s string) (c string) {
	for _, v := range s {
		if v != ' ' {
			c = c + string(v)
		}
	}
	return
}

// Сравнение Мапы со слайсом значений
func CompareWithMap(russian, answer string, mapWords *map[string][]string) bool {
	englishWords, ok := (*mapWords)[russian]
	if !ok {
		// Русское слово отсутствует в словаре
		return false
	}

	for _, word := range englishWords {
		if answer == word {
			return true // Найдено совпадение
		}
	}

	return false // Не найдено совпадение
}

func ScanTime(a *string) {
	fmt.Print("    ")
	in := bufio.NewScanner(os.Stdin)
	if in.Scan() {
		*a = in.Text()
	}
}

/*
// Сравнение строк
func CompareTime(l Word) (yes int, not int) {
	fmt.Println(l.Russian, " ||Тема: ", l.Theme)
	c := ""
	//Игнорировать пробелы
	for _, v := range l.English {
		if v != ' ' {
			c = c + string(v)
		}
	}
	var a string
	s := ""
	//Mistake-----------------------------------------------------------
	go ScanTime(&a)
	time.Sleep(10 * time.Second)
	for _, v := range a {
		if v != ' ' {
			s = s + string(v)
		}
	}

	if strings.EqualFold(c, s) {
		yes++
		fmt.Println("Yes")
	} else {
		not++
		fmt.Println("Incorect:", l.English)
	}
	return yes, not

}*/

func PrintXpen(s string) {
	var d int
	for i := 0; i <= 15; i++ {

		d++
		for i := 0; i <= d; i++ {
			fmt.Print(" ")
		}
		fmt.Println(s, "   ", s, "   ", s)
	}
}

func PrintTime(duration time.Duration) {
	minutes := int(duration.Minutes())
	seconds := int(duration.Seconds()) % 60
	fmt.Printf("Time: %d minutes %d seconds\n", minutes, seconds)
}
