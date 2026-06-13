package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	art()
}

func art() {
	var art = `
 ███  █████ █████ █   █ █████ ████  
█   █ █       █   █   █ █     █   █ 
█████ ████    █   █████ ████  ████  
█   █ █       █   █   █ █     █  █  
█   █ █████   █   █   █ █████ █   █ 
	`
	fmt.Print(art)
	time.Sleep(time.Second * 1)
	hello()
}

func hello() {
	fmt.Println("Привет! Это релиз (v1.0.0) чат-игры 'угадай число'")
	time.Sleep(time.Second * 2)
	fmt.Println("Если хотите начать игру - нажмите Y, если хотите закрыть окно - нажмите N.")
	time.Sleep(time.Second * 1)
	fmt.Println("Есть третий варинат - нажмине E чтобы посмотреть лог обновления (что было сюда добавлено)")

	for {
		var helloread string
		fmt.Scan(&helloread)
		if helloread == "Y" {
			game()
			break
		} else if helloread == "N" {
			fmt.Println("Окно будет закрыто через 3 секунды")
			time.Sleep(time.Second * 3)
			os.Exit(0)
		} else if helloread == "E" {
			log()
			break
		}
	}
}

func game() {
	fmt.Println("Добро пожаловать в игру 'угадай число'")
	time.Sleep(time.Second * 2)
	fmt.Println("Правила игры крайне просты - нужно угадать число от 1 до 100")
	fmt.Println("Чтобы начать игру нажмите Enter. Чтобы выйти из игры - нажмите Ctrl + C")
	var ready string
	fmt.Scan(&ready)
	time.Sleep(time.Second * 1)
	fmt.Println("Загадываю число, не мешай...")

	secret := rand.Intn(100) + 1
	atten := 0
	fmt.Println("Я загадал число... А не скажу. Отгадывай")

	for {
		var guess int
		fmt.Scan(&guess)
		atten++

		fmt.Println("Ты ввел(а) число:", guess)

		if guess > secret {
			fmt.Println("Загаданное число меньше")
		} else if guess < secret {
			fmt.Println("Загаданное число больше")
		} else if guess == secret {
			fmt.Println("Правильно! Ты угадал число за", atten, "попыток")
			time.Sleep(time.Second * 2)
			fmt.Println("Переношу вас в режим выбора дальнейшего действия...")
			time.Sleep(time.Second * 2)
			orno()
		}
	}
}

func log() {
	fmt.Println("Версия программы: v1.0.0")
	fmt.Println("Версия языка программирования: Go 1.18.3")
	fmt.Println("Версия компилятора: gcc 9.3.0")
	fmt.Println("Хотите посмотреть полный список функционала? (Y/N)")
	var logready string
	fmt.Scan(&logready)
	for {
		if logready == "Y" {
			fmt.Println("Проверяю базу данных...")
			time.Sleep(time.Second * 1)
			fmt.Println("Список функционала:")
			time.Sleep(time.Second * 2)
			fmt.Println("1. Добавление простого арта")
			fmt.Println("2. Добавление генератора случайных чисел (1-100)")
			fmt.Println("3. оптимизация кода (распределения функционала по циклам для дальнейшей ротации в новых проектах)")
			time.Sleep(time.Second * 3)
			fmt.Println("Переношу на выбор дальнейшего действия...")
			time.Sleep(time.Second * 2)
			orno()
			break
		}
		if logready == "N" {
			orno()
			break
		}
	}
}

func orno() {
	fmt.Println("Куда хотите перейти? (введите цифру в завиимости от выбора)")
	fmt.Println("1. В начальное меню (приветственное сообщение с артом)")
	fmt.Println("2. В игру")
	fmt.Println("3. Выйти из прогруммы")

	var orno string
	fmt.Scan(&orno)
	for {
		if orno == "1" {
			hello()
			break
		} else if orno == "2" {
			game()
			break
		} else if orno == "3" {
			fmt.Println("Окно будет закрыто через 3 секунды")
			time.Sleep(time.Second * 3)
			os.Exit(0)
		}
	}
}
