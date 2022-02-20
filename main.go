package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var words [5757]string
var letterScores = make(map[string]int)
var organizedWords = make(map[string]*[5][]string)
var wordScores = make(map[string]int)
 
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

    initializeOrganizedWords()
    organize()

    opts := createSetOpt(words[:], "treat", "YYYYN")
    fmt.Println(opts)

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

func initializeOrganizedWords() {
    for c := range letterScores {
        letter := string(c)
        organizedWords[letter] = &[5][]string{}
    }
}

func organize(){
    for _, word := range words {
        wordScores[word] = calcScore(word)
        for idx, c := range word {
            letter := string(c)
            organizedWords[letter][idx] = append(organizedWords[letter][idx], word)
        }
    }
}

func createSetOpt(wordBank []string, prevWord string, clue string) []string {
    for i, l := range prevWord {
        letter := string(l)
        rule := string(clue[i])
        if rule == "Y" {
            wordBank = intersection(wordBank, organizedWords[letter][i])
        } 
        if rule == "M" {
            wordBank = filterOut(wordBank, organizedWords[letter][i])
        }
        if rule == "N" {
            for j, c := range prevWord {
                prevL := string(c)
                if prevL == letter && string(clue[j]) == "Y" {
                    continue
                }
                wordBank = filterOut(wordBank, organizedWords[letter][j])
            }
        }
    }
    return wordBank
}

func pickBestWord(wordBank []string) string {
    highScore := 0
    bestWord := ""
    for _, w := range wordBank {
        if wordScores[w] > highScore {
            highScore = wordScores[w]
            bestWord = w
        }
    }
    return bestWord
}