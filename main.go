package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

var (
	clrError   = color.RGB(255, 80, 80)
	clrSuccess = color.RGB(80, 255, 80)
	clrWarn    = color.RGB(255, 255, 80)
	clrInfo    = color.RGB(180, 180, 255)
)

func main() {
	showGradientArt()
	mainMenu()
}

func showGradientArt() {
	artLines := []string{
		"  /$$$$$$  /$$$$$$$$ /$$$$$$$$ /$$   /$$ /$$$$$$$$ /$$$$$$$ ",
		" /$$__  $$| $$_____/|__  $$__/| $$  | $$| $$_____/| $$__  $$",
		"| $$  \\ $$| $$         | $$   | $$  | $$| $$      | $$  \\ $$",
		"| $$$$$$$$| $$$$$      | $$   | $$$$$$$$| $$$$$   | $$$$$$$/",
		"| $$__  $$| $$__/      | $$   | $$__  $$| $$__/   | $$__  $$",
		"| $$  | $$| $$         | $$   | $$  | $$| $$      | $$  \\ $$",
		"| $$  | $$| $$$$$$$$   | $$   | $$  | $$| $$$$$$$$| $$  | $$",
		"|__/  |__/|________/   |__/   |__/  |__/|________/|__/  |__/",
	}
	for i, line := range artLines {
		t := float64(i) / float64(len(artLines)-1)
		r := uint8(128 + 127*t)
		g := uint8(220 * t)
		b := uint8(255 - 15*t)
		color.RGB(int(r), int(g), int(b)).Println(line)
	}
	time.Sleep(1 * time.Second)
}

func mainMenu() {
	for {
		clrSuccess.Println("\n[СИСТЕМА] Добро пожаловать в игру 'Угадай число'")
		clrInfo.Println("Выберите действие:")
		clrWarn.Println("[1] Начать игру")
		clrWarn.Println("[2] Лог обновлений")
		clrWarn.Println("[3] Выход")

		choice := readChoice("1", "2", "3")
		switch choice {
		case "1":
			playGame()
		case "2":
			showLog()
		case "3":
			clrWarn.Println("Выход через 3 секунды...")
			time.Sleep(3 * time.Second)
			os.Exit(0)
		}
	}
}

func readChoice(allowed ...string) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		for _, a := range allowed {
			if input == a {
				return input
			}
		}
		clrError.Println("Ошибка: нужно ввести " + strings.Join(allowed, " или "))
	}
}

func playGame() {
	clrSuccess.Println("\n[СИСТЕМА] Добро пожаловать в игру 'Угадай число'")
	time.Sleep(1 * time.Second)
	clrInfo.Println("Правила: нужно угадать число от 1 до 100")
	clrInfo.Println("Нажмите Enter, чтобы начать (Ctrl+C для выхода)")
	fmt.Scanln()

	clrInfo.Println("Загадываю число...")
	secret := rand.Intn(100) + 1
	attempts := 0

	clrSuccess.Println("Число загадано. Начинайте отгадывать!")

	for {
		fmt.Print("Ваше число: ")
		guess, ok := readInt()
		if !ok {
			clrError.Println("Ошибка: нужно ввести целое число!")
			continue
		}
		attempts++

		if guess < secret {
			clrWarn.Println("Загаданное число БОЛЬШЕ")
		} else if guess > secret {
			clrWarn.Println("Загаданное число МЕНЬШЕ")
		} else {
			clrSuccess.Printf("Правильно! Ты угадал за %d попыток!\n", attempts)
			time.Sleep(2 * time.Second)
			clrInfo.Println("Возврат в главное меню...")
			time.Sleep(1 * time.Second)
			return
		}
	}
}

func readInt() (int, bool) {
	var val int
	_, err := fmt.Scan(&val)
	if err != nil {
		var discard string
		fmt.Scanln(&discard)
		return 0, false
	}
	return val, true
}

func showLog() {
	clrSuccess.Println("\n[СИСТЕМА] Информация о версии")
	clrInfo.Println("Версия программы: v1.2.0")
	clrInfo.Println("Версия Go: 1.18.3")
	clrInfo.Println("Компилятор: gcc 9.3.0")
	clrWarn.Print("Показать полный список функционала? (Y/N): ")

	answer := readChoice("Y", "N", "y", "n")
	if answer == "Y" || answer == "y" {
		clrSuccess.Println("\nСписок функционала:")
		clrInfo.Println("1. Градиентный ASCII-арт")
		clrInfo.Println("2. Генератор случайных чисел (1-100)")
		clrInfo.Println("3. Цветной вывод (fatih/color)")
		clrInfo.Println("4. Защита от некорректного ввода")
		clrInfo.Println("5. Безопасное меню без рекурсии")
		time.Sleep(3 * time.Second)
	}
	clrInfo.Println("Возврат в главное меню...")
	time.Sleep(1 * time.Second)
}
