package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var words [5757]string
var letterScores = make(map[string]int)
 
func main() {
    processWordFile()
    highScore := 0
    bestWord := ""
    
    // find the "best" starting word
    for _, word := range words {
        score := calcScore(word)
        if score > highScore {
            highScore = score
            bestWord = word
        }
    }

    fmt.Println(bestWord)
    fmt.Println(highScore)
}

func processWordFile() {
    file, err := os.Open("words.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()
 
    scanner := bufio.NewScanner(file)
    for i:=0; scanner.Scan(); i++ {
        words[i] = scanner.Text();
        for _, c := range words[i] {
            letter := string(c)
            value, mapContainsLetter := letterScores[letter]
            if mapContainsLetter {
                letterScores[letter] = value + 1
            } else {
                letterScores[letter] = 1
            }
        }
    }
 
    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}

// Function to calculate the "score" of a given word
// Takes into account the popularity of letters and 
// prefers words with non-repeating letters
func calcScore(word string) (int) {
    score := 0
    for _, c := range word {
        letter := string(c)
        score += letterScores[letter]
        n1 := strings.Count(word, letter)
        if n1 == 1 {
            score += 750
        }
    }
    return score
}