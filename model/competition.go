package model

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func (s *Slovarick) WorkTest() *Slovarick {
	startTime := time.Now()

	K := s.CreateAndInitMapWords()
	LearnSlice := NewSlovarick(nil)

	// Scan quantity words for test
	fmt.Println("Количество слов для теста")
	quantity := ScanInt()                       //количество слов для теста
	var capSlovar int = cap(s.Words) - quantity // cap  ёмкость нового []Words
	//Скопировать кусок
	NewWordsu := make([]*Word, capSlovar)
	NewSlovar := NewSlovarick(NewWordsu)
	copy(NewSlovar.Words, s.Words[:quantity])
	//Отрезать этот кусок от оригинала
	s.Words = s.Words[quantity:]
	//Присобачить что то впереди, что то сзади
	fmt.Println("                     START")
	var yes int
	var not int
	for _, v := range NewSlovar.Words {
		if v == nil {
			break
		}
		y, n, _ := Compare(*v, K)
		// тут теряется часть данных

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
		y, _, hey := Compare(*v, K)
		if !hey {
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
func Compare(l Word, mapWord map[string]string) (yes int, not int, hey bool) {
	fmt.Println(l.Russian, " ||Тема: ", l.Theme)
	c := IgnorProbel(l.English)

	a, _ := ScanStringOne()
	if a == "exit" {
		hey = false
		return yes, not, hey
	}
	s := IgnorProbel(a)

	if strings.EqualFold(c, s) {
		yes++
		fmt.Println("Yes")
	} else if ok := CompareWithMap(l.Russian, s, mapWord); ok {
		yes++
		fmt.Println("Не совсем правильно ", l.English)
	} else {
		not++
		fmt.Println("Incorect:", l.English)
		for {
			fmt.Println("Please enter correct: ")
			j, _ := ScanStringOne()
			jj := IgnorProbel(j)
			if strings.EqualFold(c, jj) {
				break
			}
		}
	}
	return yes, not, true
	/* Если захочется игнорировать одну ошибку в слове
	if moreThanOneMistake(c, s) {
		yes++
		fmt.Println("Yes")
	} else {
		not++
		fmt.Println("Incorect:", l.English)
	}
	return yes, not*/
}

func (s *Slovarick) CreateAndInitMapWords() (MapWords map[string]string) {
	MapWords = make(map[string]string)
	for _, v := range s.Words {
		if v == nil {
			break
		}
		MapWords[v.Russian] = v.English // assignment to nil map (SA5000)go-staticcheck field Russian string
	}
	return
}

func IgnorProbel(s string) (c string) {
	for _, v := range s {
		if v != ' ' {
			c = c + string(v)
		}
	}
	return
}

func CompareWithMap(rusian, answer string, mapString map[string]string) (yes bool) {
	if answer == mapString[rusian] {
		yes = true
	} else {
		yes = false
	}
	return
}

/*
// Игнорировать одну ошибку в словах
func moreThanOneMistake(first, second string) bool {
	first = strings.ToLower(first)
	second = strings.ToLower(second)

	lenFirst, lenSecond := len(first), len(second)
	if strings.EqualFold(first, second) {
		//fmt.Println("100% duplicates")
		return true
	}

	if (lenFirst-lenSecond) >= 2 || (lenSecond-lenFirst) >= 2 {
		//fmt.Println("try more, over time")
		return false
	}

	if (lenSecond - lenFirst) == 1 {
		return quantityMistakes(first, second)
	}

	if (lenFirst - lenSecond) == 1 {
		return quantityMistakes(second, first)
	}

	return true
}

func quantityMistakes(a, b string) bool {
	sizeMistake := 0
	for i, v := range a {
		if v != rune(b[i+sizeMistake]) {
			sizeMistake++
			if sizeMistake >= 2 {
				return false
			}
		}
	}

	return true
}*/

func ScanTime(a *string) {
	fmt.Print("    ")
	in := bufio.NewScanner(os.Stdin)
	if in.Scan() {
		*a = in.Text()
	}
}

// Сравнение строк /
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

}

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
