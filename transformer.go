package main

import "strings"

const upperAlphabet = "ABCDEFGHIJKLMNÑOPQRSTUVWXYZ"

//Snakeazer  ...
func Snakeazer(camelString string) string {
	snakeString := ""
	for i := 0; i < len(camelString); i++ {
		s := string(camelString[i])
		if strings.Contains(upperAlphabet, s) && i > 0 {
			snakeString += "_" + strings.ToLower(s)
		} else {
			snakeString += strings.ToLower(s)
		}
	}
	return snakeString
}
