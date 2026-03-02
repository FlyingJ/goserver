package util

import (
	"strings"
)

var Profanities = []string{
	"kerfuffle",
	"sharbert",
	"fornax",
}

func Censor(filthyString string) string {
	cleanString := ""
	var profane bool
	for _, word := range strings.Split(filthyString, " ") {
		lowerWord := strings.ToLower(word)
		profane = false
		for _, profanity := range Profanities {
			if lowerWord == profanity {
				profane = true
				break
			}
		}
		var addend string
		if profane {
			addend = "****"
		} else {
			addend = word
		}
		var joiner string
		if cleanString == "" {
			joiner = ""
		} else {
			joiner = " "
		}
		cleanString = strings.Join([]string{cleanString, addend}, joiner)
	}
    return cleanString
}