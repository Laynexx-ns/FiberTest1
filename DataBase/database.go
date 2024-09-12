package DataBase

import (
	"log"
)

var baseJsonNameText = map[string]string{}

func AddBaseJsonNameText(name, text string) {
	baseJsonNameText[name] = text
	log.Println("added new data to: baseJsonNameText")
}

func GetBaseJsonNameText() map[string]string {
	return baseJsonNameText
}
