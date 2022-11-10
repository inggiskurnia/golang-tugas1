package main  
  
import (  
 "fmt"  
)  
  
func isVowel(character rune) {  
 if character == 'a' || character == 'e' || character == 'i' || character == 'o' || character == 'u' {  
  fmt.Printf(" %c is vowel\n", character)  
 } else {  
  fmt.Printf(" %c is consonant\n", character)  
 }  
  
}  
func main() {  
 isVowel('a') // vowel  
 isVowel('b') // consonant  
}  
