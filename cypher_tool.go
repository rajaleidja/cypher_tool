package main

import (
	"fmt"
	"strings"
)
// Main logic, envoking other functions
func main() {
	toEncrypt, encoding, message := getInput() // gets the variables
	if toEncrypt == true {			// if "toEncrypt" is "yes" we encrypt
		switch encoding {
		case "1":
			encrypt_rot13(message) 		// if "1" we encypt with rot13
		case "2":
			encrypt_reverse(message) 	// if "2" we encypt with reverse
		case "3":
			encrypt_ROTX(message) 		// if "2" we encypt with rot with user input
		}
	} else {
		switch encoding {
		case "1":
			decrypt_rot13(message) 		// if "1" we Decypt with rot13
		case "2":
			decrypt_reverse(message) 	// if "2" we Decypt with reverse
		case "3":
			decrypt_ROTX(message) 		// if "3" we encypt with rot with user input
		}
	}
}
// Get the input data required for the operation
func getInput() (toEncrypt bool, encoding string, message string) {
	var choise int
	var choise1 int
	fmt.Println("Welcome to the Cypher Tool!")
encrypt: //this is goto's label
	choise = 0 		//we make 0 for that moment, when we have done input and going to gotos label, variable "choise" contains previous input
	fmt.Println("Select operation (1/2):\n1. Encrypt.\n2. Decrypt.")
	fmt.Scanln(&choise) // we scan users input
	if choise < 1 || choise > 2 || choise != int(choise) {
		fmt.Println("Invalid input:", choise, "try again!")
		goto encrypt 				// this is "goto" 
	} else {
		if choise == 1 {
			toEncrypt = true // here we making "toEncrypt" true
		} else {
			toEncrypt = false // or false
		} //					dependig on what user chosed in fmt.Scanln
	}
cypherChoice: // goto's label 
	choise1 = 0 //we make 0 for that moment, when we have done input and going to gotos label, variable "choise1" contains previous input
	fmt.Println("Select cypher (1/2/3):\n1. ROT13.\n2. Reverse.\n3. ROT-X with user input")
	fmt.Scanln(&choise1) // scanning the users input
	if choise1 < 1 || choise1 > 3 || choise1 != int(choise1) {
		fmt.Println("Invalid input:", choise1, "try again!")
		goto cypherChoice // goto
	} else {
		if choise1 == 1 { 		//the users choice is 1 we do encoding 1 >> in the function main() in switch "encoding"
			encoding = "1"
		} else if choise1 == 2 { //the users choice is 2 we do encoding 2 >> in the function main() in switch "encoding"
			encoding = "2"
		} else if choise1 == 3 { //the users choice is 2 we do encoding 3 >> in the function main() in switch "encoding"
			encoding = "3"
		} else {
			goto cypherChoice // if something has been done wrong we do goto
		}
	}
entermsg: //goto's label
	fmt.Println("Enter the message:")
	fmt.Scanln(&message) // scanning the users message
	if message == "" || message == " " {
		fmt.Println("Invalid input:", message, "try again!")
		goto entermsg // goto
	}
	return toEncrypt, encoding, message
}

// Encrypt the message with rot13
func encrypt_rot13(s string) string {           // DONE BY OLEH ILCHUK
	result := make([]rune, len(s))				
	for i, char := range s {					
		if 'a' <= char && char <= 'z' {			
			result[i] = 'a' + (char-'a'+13)%26  
		} else if 'A' <= char && char <= 'Z' {  
			result[i] = 'A' + (char-'A'+13)%26  
		} else {								
			result[i] = char					
		}										
	}											
	fmt.Println(string(result))					
	return string(result)						
}

// Encrypt the message with reverse
func encrypt_reverse(input string) string {     								// DONE BY ENAM HOSSAIN
	var result strings.Builder													
	for _, char := range input {												
		if ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') {		
			if 'a' <= char && char <= 'z' {										
				result.WriteRune('z' - (char - 'a'))							
			} else {															
				result.WriteRune('Z' - (char - 'A'))							
			}																	
		} else {																
			result.WriteRune(char) // Non-alphabet characters are left unchanged
		}																		
	}																			
	fmt.Println(result.String())												
	return result.String()														
}

// Decrypt the message with rot13
func decrypt_rot13(s string) string {						// DONE BY OLEH ILCHUK
	result := make([]rune, len(s))							
	for i, char := range s {								
		if 'a' <= char && char <= 'z' {						
			result[i] = 'a' + (char-'a'+13)%26				
		} else if 'A' <= char && char <= 'Z' {				
			result[i] = 'A' + (char-'A'+13)%26				
		} else {											
			result[i] = char								
		}													
	}														
	fmt.Println(string(result))								
	return string(result)									
}															

// Decrypt the message with reverse
func decrypt_reverse(input string) string {										// DONE BY ENAM HOSSAIN
	var result strings.Builder													
	for _, char := range input {												
		if ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') {		
			if 'a' <= char && char <= 'z' {										
				result.WriteRune('z' - (char - 'a'))							
			} else {															
				result.WriteRune('Z' - (char - 'A'))							
			}																	
		} else {																
			result.WriteRune(char) // Non-alphabet characters are left unchanged
		}																		
	}																			
	fmt.Println(result.String())												
	return result.String()														
}																				

// encrypt the message with ROT X -------------------------------------------------------------------------
func encrypt_ROTX(s string) string {
	var step rune
	var runeToStr string = ""
	
back: 					// goto's label
	fmt.Println("Type index of ROT:")
	step = 0  		//we make 0 for that moment, when we have done input and going to gotos label, variable "step" contains previous input
	fmt.Scanln(&step) 		//scanning input
	if step < 1 || step != rune(step) {  // checking for step was lower than 1 OR step dont was step's rune

		fmt.Println("try again, ", step, " is invalid input")
		goto back // goto
	}
	for step > 25 { 		// if step bigger than 25
		step = step % 26 	// we do step - 26 many times, to get step lower than 27
	}						// i writed step%26 because here we can get same anvser 

	fmt.Printf("ROT%v is DECRYPTING word \"%s\" With step %v\n", step, s, step)

	runes := []rune(s) 				// we slice the string S to the symbols
	for i := 0; i < len(s); i++ { 	
		if runes[i] == 32 {
			runeToStr += " "
			i++
		}	
	ostav: // gotos label
		if runes[i] < 97 && runes[i] > 'Z' || runes[i] < 'A' || runes[i] > 122 { // check that symbol i was lower 97(decimal of small a) AND bigger than Z decimal(122) OR symbol lower than A decimal(65) OR symbol was bigger than 122(decimal of small z)
			runeToStr += string(runes[i]) // writing symbol that we DONT change to variable runeToStr
			i++ 			// skip this symbol
			goto ostav // goto
		}
		ch := runes[i] + rune(int(step)) 		// we do symbol + step like a+3 = e
		if ch > 122 || ch > 'Z' && ch < 97 { // if ch bigger than 122(decimal of small z) OR ch bigger than Z decima(122) AND ch smaller than 97(decimal of small a)
			ch = ch - 26 
		}
		runeToStr += string(ch) 	// writing changed symbol to the variable
	}
	fmt.Println("Thank you for using our programm, \nhere is your Decrypted message:", runeToStr)
	return runeToStr
}

// Decrypt the message with ROT X -------------------------------------------------------------------------

func decrypt_ROTX(s string) string {
	var step rune
	var runeToStr string = ""
	
back: 					// goto's label
	fmt.Println("Type index of ROT:")
	step = 0  		//we make 0 for that moment, when we have done input and going to gotos label, variable "step" contains previous input
	fmt.Scanln(&step) 		//scanning input
	if step < 1 || step != rune(step) {  // checking for step was lower than 1 OR step dont was step's rune

		fmt.Println("try again, ", step, " is invalid input")
		goto back // goto
	}
	for step > 25 { 		// if step bigger than 25
		step = step % 26 	// we do step - 26 many times, to get step lower than 27
	}						// i writed step%26 because here we can get same anvser 

	fmt.Printf("ROT%v is DECRYPTING word \"%s\" With step %v\n", step, s, step)

	runes := []rune(s) 				// we slice the rune S to the symbols
	for i := 0; i < len(s); i++ { 	
		if runes[i] == 32 {
			runeToStr += " "
			i++
		}	
	ostav: // gotos label
		if runes[i] < 97 && runes[i] > 'Z' || runes[i] < 'A' || runes[i] > 122 { // check that symbol i was lower 97(decimal of small a) AND bigger than Z decimal(122) OR symbol lower than A decimal(65) OR symbol was bigger than 122(decimal of small z)
			runeToStr += string(runes[i]) // writing symbol that we DONT change to variable runeToStr
			i++ 			// skip this symbol
			goto ostav // goto
		}
		ch := runes[i] + rune(int(step)) 		// we do symbol + step like a+3 = e
		if ch > 122 || ch > 'Z' && ch < 97 { // if ch bigger than 122(decimal of small z) OR ch bigger than Z decima(122) AND ch smaller than 97(decimal of small a)
			ch = ch - 26 
		}
		runeToStr += string(ch) 	// writing changed symbol to the variable
	}
	fmt.Println("Thank you for using our programm, \nhere is your Decrypted message:", runeToStr)
	return runeToStr
}