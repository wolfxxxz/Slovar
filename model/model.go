package model

import "strings"

// Word model
type Word struct {
	ID         int    `json:"id"`
	English    string `json:"english"`
	Russian    string `json:"russian"`
	Theme      string `json:"theme"`
	RightAswer int    `json:"rightAnswer"`
}

// Create Library
func NewLibrary(newId int, newEnglish string, newRussian string, newTheme string) *Word {
	d := &Word{ID: newId, English: newEnglish, Russian: newRussian, Theme: newTheme}
	return d
}

type Slovarick struct {
	Words []*Word
}

func NewSlovarick(a []*Word) *Slovarick {
	return &Slovarick{Words: a}
}

func (s *Slovarick) AppendWord(w *Word) {
	s.Words = append(s.Words, w)
}
func (s *Slovarick) Preppend(w *Word) {
	sliceWords := []*Word{w}
	s.Words = append(sliceWords, s.Words...)
}

// Псевдоним типа
type Words []*Word

func (s *Slovarick) CreateAndInitMapWords() *map[string][]string {
	maps := make(map[string][]string)
	for _, w := range s.Words {
		maps[w.Russian] = append(maps[w.Russian], strings.ToLower(w.English))
	}
	return &maps
}

/*
func (s *Slovarick) CreateAndInitMapWordsOld() (MapWords map[string]string) {
	MapWords = make(map[string]string)
	for _, v := range s.Words {
		if v == nil {
			break
		}
		MapWords[v.Russian] = v.English // assignment to nil map (SA5000)go-staticcheck field Russian string
	}
	return
}
*/
