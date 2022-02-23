package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var words [5757]string
var letterScores = make(map[string]int)
var wordScores = make(map[string]int)
var wordMap = make(map[string]*[5][]string)
var currGuess string
var answer string

func main() {

    flag.StringVar(&currGuess, "s", "thing", "Specify starting word. Default is 'thing'.")
    flag.StringVar(&answer, "a", "", "Specify answer. Default mode runs without answer.")

    flag.Parse()

    processWordFile()
    createWordMap()

    currWordBank := words[:];
    guessCount := 1

    if answer != "" {
        for {
            fmt.Println("My guess is: " + currGuess)
            res := checkGuess(currGuess, answer)
            if res == "YYYYY" {
                fmt.Println("Got the right answer in " + strconv.Itoa(guessCount) + " guesses")
                break
            }
            currWordBank = filterWords(currWordBank, currGuess, res)
            currGuess = pickBestWord(currWordBank)
            guessCount++
        }
    } else {
        for {
            fmt.Println("My guess is: " + currGuess);
            res := askResult()
            if res == "YYYYY" {
                fmt.Println("I won in " + strconv.Itoa(guessCount) + " guesses")
                break
            }
            currWordBank = filterWords(currWordBank, currGuess, res)
            currGuess = pickBestWord(currWordBank)
            guessCount++
        }
    }
}

func checkGuess(guess string, ans string) string {
    res := ""
    for i, l := range guess {
        letter := string(l)
        answerL := string(ans[i])
        if letter == answerL {
            res += "Y"
        } else if !strings.Contains(ans, letter){
            res += "N"
        } else {
            res += "M"
        }
    }
    guessMap := make(map[string][]int)
    for i, c := range guess {
        _, contains := guessMap[string(c)]
        if !contains && string(res[i]) == "Y" {
            guessMap[string(c)] = []int{i}
        } else if string(res[i]) == "Y"  {
            guessMap[string(c)] = append(guessMap[string(c)], i)
        }
    }

    newRes := ""
    for i, c := range res {
        if string(res[i]) == "M" {
            if strings.Count(ans, string(guess[i])) == len(guessMap[string(guess[i])]) {
                newRes += "N"
                continue
            }
        }
        newRes += string(c)
    }
    return newRes
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
        // n1 := strings.Count(word, letter)
        // if n1 == 1 {
        //     score += 750
        // }
    }
    return score
}

func initializeWordMap() {
    for c := range letterScores {
        letter := string(c)
        wordMap[letter] = &[5][]string{}
    }
}

func createWordMap(){
    initializeWordMap()
    for _, word := range words {
        wordScores[word] = calcScore(word)
        for idx, c := range word {
            letter := string(c)
            wordMap[letter][idx] = append(wordMap[letter][idx], word)
        }
    }
}

func filterWords(wordBank []string, prevWord string, clue string) []string {
    for i, l := range prevWord {
        letter := string(l)
        rule := string(clue[i])
        if rule == "Y" {
            wordBank = intersection(wordBank, wordMap[letter][i])
        } 
        if rule == "M" {
            // add all the sets that include this letter except the one at this spot
            allWords := []string{}
            for j, _ := range wordMap[letter] {
                if j == i {
                    continue
                }
                allWords = append(allWords, wordMap[letter][j]...)
            }
            wordBank = intersection(wordBank, allWords)
        }
        if rule == "N" {
            for j, c := range prevWord {
                prevL := string(c)
                if prevL == letter && string(clue[j]) == "Y" {
                    continue
                }
                wordBank = filterOut(wordBank, wordMap[letter][j])
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

func askResult() string {
    var s string
    r := bufio.NewReader(os.Stdin)
    for {
        fmt.Fprint(os.Stderr, "How was my guess? ")
        s, _ = r.ReadString('\n')
        if s != "" {
            break
        }
    }
    return strings.TrimSpace(s)
}


