package model

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

// Unmarshal (decode) раскодировать
func (SliceWord *Slovarick) TakeGob(file string) {
	jsonData, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("TakeGob", err)
	}

	buf := bytes.NewBuffer(jsonData)
	dec := gob.NewDecoder(buf)

	if err := dec.Decode(SliceWord); err != nil {
		log.Fatal(err)
	}
}

// Сохранить в json file
// Marshal (закодировать) encode
func (s *Slovarick) SaveGob(file string) {

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(s); err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(byteArr))             0664
	err := os.WriteFile(file, buf.Bytes(), 0666) //-rw-rw-rw-
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
