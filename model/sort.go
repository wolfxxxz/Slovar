package model

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

// Sortlibrary includes 5 Option
func (l Slovarick) SortLibrary() {
	fmt.Println("sort Theme 1 || sort Englisch 2 || sort Russian 3 || mix up 4 || sort RightAnswer 5")
	c, _ := ScanStringOne()
	cc, err := strconv.Atoi(c)
	if err != nil {
		fmt.Println("Incorect, please enter number")
	}
	if cc == 1 {
		l.sortLibraryTheme()
	} else if cc == 2 {
		l.sortLibraryEnglisch()
	} else if cc == 3 {
		l.sortLibraryRussian()
	} else if cc == 4 {
		l.mixUp()
	} else if cc == 5 {
		l.sortLibraryAnswer()
	} else {
		fmt.Println("You are lazy")
	}
}

func (l *Slovarick) sortLibraryTheme() {
	sort.SliceStable(l.Words, func(i, j int) bool {
		return l.Words[i].Theme < l.Words[j].Theme
	})

}

func (l *Slovarick) sortLibraryEnglisch() {
	sort.SliceStable(l.Words, func(i, j int) bool {
		return l.Words[i].English < l.Words[j].English
	})

}

func (l *Slovarick) sortLibraryRussian() {
	sort.SliceStable(l.Words, func(i, j int) bool {
		return l.Words[i].Russian < l.Words[j].Russian
	})

}

func (l *Slovarick) mixUp() {
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	perm := randGen.Perm(len(l.Words))
	for i, j := range perm {
		l.Words[i], l.Words[j] = l.Words[j], l.Words[i]
	}

}
func (l *Slovarick) sortLibraryAnswer() {
	sort.SliceStable(l.Words, func(i, j int) bool {
		return l.Words[i].RightAswer < l.Words[j].RightAswer
	})
}

// i file.txt

func (l *Slovarick) MixUpTwo() {
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	perm := randGen.Perm(len(l.Words))
	for i, j := range perm {
		l.Words[i], l.Words[j] = l.Words[j], l.Words[i]
	}
}

func (l *Slovarick) MixUpTime() {
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	perm := randGen.Perm(len(l.Words))
	for i, j := range perm {
		l.Words[i], l.Words[j] = l.Words[j], l.Words[i]
	}
}

func (s *Slovarick) ReverseSlice() {
	for i, j := 0, len(s.Words)-1; i < j; i, j = i+1, j-1 {
		s.Words[i], s.Words[j] = s.Words[j], s.Words[i]
	}
}
