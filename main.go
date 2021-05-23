package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)


const (
	tok_eof = -1

	// команды (ключевые слова)
	tok_def    = -2
	tok_extern = -3

	// операнды (идентификаторы, числа)
	tok_identifier = -4
	tok_number     = -5
)

var (
	IdentifierStr string
	NumVal        float64
)

func checkError(err error) {
	if err != nil {
		log.Panicf("err: %s", err)
	}
}

func getchar() rune {
	 
	_, err:=
	if err == io.EOF {
		return 0
	}
	checkError(err)
	return b[0]
}

func gettok() int {
	LastChar := ' '

	// Пропускаем пробелы.
	for LastChar == ' ' {
		LastChar = getchar()
	}
}

func main() {

}
