package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	// Read words from text file
	wordsFile, err := os.Open("words.txt")
	if err != nil {
		fmt.Println("Failed to open word file.")
		return
	}
	defer wordsFile.Close()

	var words []string
	scanner := bufio.NewScanner(wordsFile)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	reader := bufio.NewReader(os.Stdin)

	score := 0
	timer := time.NewTimer(60 * time.Second)

	fmt.Println("Starting typing game. Try to accurately type as many words as possible within 60 seconds.")

gameLoop:
	for {
		word := words[rand.Intn(len(words))]
		fmt.Printf("Type this word: %s\n", word)

		inputCh := make(chan string)
		go func() {
			input, _ := reader.ReadString('\n')
			inputCh <- strings.TrimSpace(input)
		}()

		select {
		case <-timer.C:
			break gameLoop
		case input := <-inputCh:
			if input == word {
				score++
				fmt.Println("Correct!")
			} else {
				fmt.Println("Incorrect...")
			}
		}
	}

	fmt.Printf("Time's up! Your score is %d points.\n", score)
}

