package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	lowerChars   = "abcdefghijklmnopqrstuvwxyz"
	upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberChars  = "0123456789"
	specialChars = "!@#$%^&*()-_=+[]{}|;:,.<>?/"
)

func generatePass(length int, useUpper bool, useNumber bool, useSpecial bool) string {
	element := lowerChars // 半角英数字は確定で使用する予定(大体いるでしょう)
	if useUpper {
		element += upperChars
	}
	if useNumber {
		element += numberChars
	}
	if useSpecial {
		element += specialChars
	}

	password := make([]byte, 0, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < length; i++ {
		randIndex := r.Intn(len(element))
		password = append(password, element[randIndex])
	}

	return string(password)
}

func promptForBool(message string) bool {
	for {
		fmt.Print(message)
		var input string
		fmt.Scan(&input)
		input = strings.ToLower(strings.TrimSpace(input))
		if input == "true" || input == "t" || input == "1" {
			return true
		} else if input == "false" || input == "f" || input == "0" {
			return false
		} else {
			fmt.Println("入力が無効です。true または false を入力してください。")
		}
	}
}

func promptForInt(message string) int {
	for {
		fmt.Print(message)
		var input string
		fmt.Scan(&input)
		value, err := strconv.Atoi(strings.TrimSpace(input))
		if err == nil && value > 0 {
			return value
		} else {
			fmt.Println("入力が無効です。正の整数を入力してください。")
		}
	}
}

func main() {
	ascii := []string{
		"##   ##   ####             ######     ##      #####    #####",
		"### ###    ##               ##  ##   ####    ##   ##  ##   ##",
		"#######    ##               ##  ##  ##  ##   #        #",
		"#######    ##     ######    #####   ##  ##    #####    #####",
		"## # ##    ##               ##      ######        ##       ##",
		"##   ##    ##               ##      ##  ##   ##   ##  ##   ##",
		"##   ##   ####             ####     ##  ##    #####    #####",
	}

	for _, line := range ascii {
		fmt.Println(line)
	}

	fmt.Println("****************************")
	fmt.Println(" パスワード生成ツール v1.0")
	fmt.Println(" (c) 2024 powered by Miranon")
	fmt.Println("****************************")

	for {
		length := promptForInt("パスワードの長さを入力してください: ")
		useUpper := promptForBool("大文字を含めますか？ (true/false): ")
		useNumber := promptForBool("数字を含めますか？ (true/false): ")
		useSpecial := promptForBool("特殊文字を含めますか？ (true/false): ")

		password := generatePass(length, useUpper, useNumber, useSpecial)
		fmt.Println("生成されたパスワード:", password)

		fmt.Print("終了しますか？ (y/n): ")
		var exitInput string
		fmt.Scan(&exitInput)
		if strings.ToLower(strings.TrimSpace(exitInput)) == "y" {
			fmt.Println("プログラムを終了します。")
			break
		}
	}
}
