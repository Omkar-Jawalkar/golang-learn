package main

import (
	"fmt"
	"time"
	"unicode/utf8"
)


var pl = fmt.Println

// main is the entry point of the program.
func main() {

		vStr := "abcdefg"

		pl("UNI CODE " , utf8.RuneCountInString(vStr))

		for i, val := range vStr {
			fmt.Printf("%d : %#U : %c\n", i , val ,val)
		}



		now := time.Now();
		pl(now.Year(), now.Day(), now.Month())


		// rand
 

	
	
}