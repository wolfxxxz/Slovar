package model

import (
	"fmt"
	"strings"
)

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

func (old *Slovarick) CheckAndDelDublikats(new *Slovarick) {
	var count = 0
	var withoutDublicat []*Word //:= s.Words
	for i := 0; i <= len(new.Words)-1; i++ {
		for _, v := range old.Words {
			if strings.EqualFold(v.English, new.Words[i].English) {
				count++
			}
		}
		if count == 0 {
			withoutDublicat = append(withoutDublicat, new.Words[i])
		} else {
			count = 0
		}
	}

	fmt.Println(len(withoutDublicat))

	*new = *NewSlovarick(withoutDublicat)
}
