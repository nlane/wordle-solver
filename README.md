# Wordle Solver

A Go program to solve Wordle puzzles. It guesses words that:

- Have the most popular letters
- Have a diverse set of letters
- Follow all the given rules learned from previous guesses

## How to Use: Solver Mode

In solver mode, you must run the program with the `-a` flag and supply the intended answer such as: `go run . -a <answer-here>`.

The program will keep spitting out guesses and checking those guesses against the given answer until it gets it right.

## How to Use: Interactive Mode

In interactive mode, you must input the results for each of the program's guesses. This simulates playing the daily Wordle where you don't know the answer and need inspiration for guesses based on the feedback from your last guess.

The program will ouput its guess and ask for feedback such as:

```
My guess is: thing
How was my guess?
```

You input a 5 letter string using the characters:

- Y - meaning it guessed the right letter for that spot
- M - meaning it guessed the right letter but in the wrong spot
- N - meaning it guessed a letter that does not appear in the word OR has already been accounted for

Going back to the previous example, if your answer is "tints" then you would input:

```
My guess is: thing
How was my guess? YNMMN
```

Since the "t" was correct, "h" and "g" were wrong, and "i" and "n" were correct letters but in the wrong spots.

To run in interactive mode just run the program without any flags: `go run .`

## Flags

All flags are optional. There are 2 flags:

- `-a` - specifies the answer to use when running in "Solver" mode
- `-s` - specifies the word to start with (defaults to "thing")
