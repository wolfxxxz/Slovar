package model

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
	return &Slovarick{a}
}

func (s *Slovarick) AppendWord(w *Word) {
	s.Words = append(s.Words, w)
}
func (s *Slovarick) Preppend(w *Word) {
	sliceWords := []*Word{w}
	s.Words = append(sliceWords, s.Words...)
}

//type Words []*Word
