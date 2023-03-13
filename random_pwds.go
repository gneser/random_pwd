package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	lowerCharSet = "abcdedfghijklmnopqrstuvwxyz"
	upperCharSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberSet    = "0123456789"
	allCharSet   = lowerCharSet + upperCharSet + numberSet
)

func main() {

	//flag.IntVar(&)

	rand.Seed(time.Now().Unix())

	minNum := 1
	minUpperCase := 1
	passwordLength := 16

	f, e := os.Create("passwords.txt")
	if e != nil {
		panic(e)
	}
	defer f.Close()

	// Number of lines
	total := 20

	for j := 0; j < total; j++ {
		fmt.Fprintln(f, generatePassword(passwordLength, minNum, minUpperCase))
	}

}

func generatePassword(passwordLength, minNum, minUpperCase int) string {
	var password strings.Builder

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}
