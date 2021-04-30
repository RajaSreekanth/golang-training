package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "Velit officia dolor excepteur eu Lorem deserunt excepteur anim ullamco. Sunt ex tempor sit ullamco do irure irure pariatur nisi. Ipsum ullamco id ex elit ut labore dolor amet Lorem. Anim non aliquip cupidatat duis. Sint commodo sit esse ipsum sunt in ea excepteur enim fugiat. Sunt tempor cillum laboris enim elit sint tempor ullamco culpa ut ea. Eu elit do est voluptate deserunt exercitation aliquip labore anim duis velit eiusmod laborum irure. Amet cupidatat ad veniam in aliqua velit pariatur elit incididunt esse veniam nulla. In irure sunt aliqua minim ipsum nostrud sit nulla dolore sit."
	sentence = sentence[0 : len(sentence)-1]
	fmt.Println(sentence)
	words := strings.Split(sentence, " ")

	myMap := map[int]int{
		0: 0,
	}
	for i := 0; i < len(words); i++ {
		size := len(words[i])
		if length, exists := myMap[size]; exists {
			myMap[size] = length + 1
		} else {
			myMap[size] = 1
		}
	}

	fmt.Println(myMap)

	// find highestOccurances value in the keys
	var highestOccurances, maxOccuredWordLength int = 0, 0

	for key, value := range myMap {
		if value > highestOccurances {
			highestOccurances = value
			maxOccuredWordLength = key
		}
	}

	fmt.Printf("the most occurances of a word is %d and length of it is %d", highestOccurances, maxOccuredWordLength)

}
