package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

/*
	Задача 1 состоит в том чтобы читать файл и читать CLI. У нас даны требования  $ expense-tracker add --description "Lunch" --amount 20

То есть работаем с ними.
Первое что требуется читать CLI. Потом мы разбиваем флаги и что за что отвечает int string в нашем случае еще могут быть типы, этого точно не знаю
дальше мы разбиваем логику на switch case и парсим все это дело то есть у нас в add будет лежать внутри addExpense и по аналогии проделываем
это с остальными тремя. Дальше что у нас получилось мы сохраняем, каким образом. loadExpenses с проверкой на ошибки и дальше saveExpenses это файлы json
для истории наших трат. Вроде все
*/
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Не достаточно данных, используйте программу: expense-tracker <команда> [флаги]")
		os.Exit(1)
	}
	expenses, err := loadExpenses()
	if err != nil {
		log.Fatalf("Не удалось загрузить расходы: %v", err)
	}
	switch os.Args[1] {
	default:
		fmt.Println("неизвестная команда:", os.Args[1])
		os.Exit(1)

	case "add":
		var addedExpense Expense
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		description := addCmd.String("description", "", "описание расхода")
		amount := addCmd.Float64("amount", 0, "сумма траты")
		addCmd.Parse(os.Args[2:])
		if *description == "" || *amount <= 0 {
			fmt.Println("Проверьте формат траты на ее или ее значение.")
			os.Exit(1)
		}
		expenses, addedExpense, err = addExpense(expenses, *description, *amount)
		if err != nil {
			log.Fatalf("Не удалось добавить расходы")
		}
		err = saveExpenses(expenses)
		if err != nil {
			log.Fatal()
		}
		fmt.Println("Трата:", addedExpense, "была успешно добавлена")
	}

}
