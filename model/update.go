package model

import "fmt"

// Соединяем два среза в один
func (oldWords *Slovarick) UpdateLibrary(filetxt string) {
	var NewWords Slovarick
	NewWords.TakeTXT(filetxt)

	c := len(oldWords.Words)
	//--------Соединяем два среза в один--------------
	oldWords.Words = append(NewWords.Words, oldWords.Words...)
	//NewWords.Words = append(NewWords.Words, oldWords.Words...)

	d := len(oldWords.Words)
	// Записать в filetxt пустой
	SaveEmptyTXT(filetxt, "You need to add your words here")

	if d != c {
		fmt.Println("                   New Words Add:", d-c)
	} else {
		fmt.Println("Для загрузки слов списком необходимо упорядочить и вставить слова в файл `save/newWords.txt`")
		fmt.Println("english - перевод - тема")
		fmt.Println("в конце оставить пустую строчку")
		fmt.Println("I believe in you!!!")
	}
}

func (oldWords *Slovarick) UpdateLibraryOnlyNewWords(NewWords *Slovarick) {

	c := len(oldWords.Words)
	// Проверить на дубликаты

	//--------Соединяем два среза в один--------------
	oldWords.Words = append(NewWords.Words, oldWords.Words...)
	//NewWords.Words = append(NewWords.Words, oldWords.Words...)

	d := len(oldWords.Words)
	// Записать в filetxt пустой

	if d != c {
		fmt.Println("                   New Words Add:", d-c)
	} else {
		fmt.Println("Для загрузки слов списком необходимо упорядочить и вставить слова в файл `save/newWords.txt`")
		fmt.Println("english - перевод - тема")
		fmt.Println("в конце оставить пустую строчку")
		fmt.Println("I believe in you!!!")
	}
}
