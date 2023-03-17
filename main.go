package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"math/rand"
)

func main() {
	// テキストファイルから単語を読み込む
	wordsFile, err := os.Open("words.txt")
	if err != nil {
		fmt.Println("単語ファイルを開けませんでした。")
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

	fmt.Println("タイピングゲームを開始します。60秒間でどれだけ正確に入力できるか挑戦してください！")

gameLoop:
	for {
		word := words[rand.Intn(len(words))]
		fmt.Printf("タイプしてください: %s\n", word)

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
				fmt.Println("正解！")
			} else {
				fmt.Println("不正解...")
			}
		}
	}

	fmt.Printf("時間切れ！あなたのスコアは %d 点です。\n", score)
}

