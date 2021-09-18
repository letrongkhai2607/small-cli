package heplers

import (
	"log"
)


// HandleError is a reusable error checking function
func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func IsAdult(isAdult bool) string{
	if isAdult == false{
		return " False"
	}
	return " True"
}
