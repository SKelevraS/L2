package main

import (
	"fmt"
	"os"
	


	"github.com/beevik/ntp"
)

/**
1. Создать программу печатающую точное время с использованием NTP -библиотеки. Инициализировать как go module.
Использовать библиотеку github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.
*/

func main() {
	res, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
		return
	}

	fmt.Println(res)


}
