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

	case "add":
		var addedExpense Expense

		addCmd := flag.NewFlagSet("add", flag.ExitOnError)

		description := addCmd.String("description", "", "описание расхода")

		amount := addCmd.Float64("amount", 0, "сумма траты")

		addCmd.Parse(os.Args[2:])

		if *description == "" || *amount <= 0 {
			fmt.Println("Трата не должна быть пустой.")
			os.Exit(1)
		}

		expenses, addedExpense, err = addExpense(expenses, *description, *amount)

		if err != nil {
			log.Fatalf("Не удалось добавить расходы: %v", err)
		}

		err = saveExpenses(expenses)

		if err != nil {
			log.Fatalf("Не удалось сохранить: %v", err)
		}

		fmt.Println("Трата:", addedExpense, "была успешно добавлена")

	case "delete":

		deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)

		ID := deleteCmd.Int("id", 0, "expense's ID")

		deleteCmd.Parse(os.Args[2:])

		if *ID <= 0 {
			fmt.Println("Неверный ID.")
			os.Exit(1)
		}
		expenses, err = deleteExpense(expenses, *ID)

		if err != nil {
			log.Fatalf("Не смогли удалить данную трату:%v", err)
		}
		err = saveExpenses(expenses)
		if err != nil {
			log.Fatalf("Не смогли сохранить: %v", err)
		}
		fmt.Println("Трата была успешно удалена")

	case "summary":
		summaryCmd := flag.NewFlagSet("summary", flag.ExitOnError)
		month := summaryCmd.Int("month", 0, "номер месяца (1-12)")
		summaryCmd.Parse(os.Args[2:])

		if *month == 0 {
			fmt.Printf("Total expenses: $%.2f\n", summaryExpenses(expenses))
		} else if *month < 1 || *month > 12 {
			fmt.Println("Неправильный месяц: нужно число от 1 до 12.")
			os.Exit(1)
		} else {
			total := summaryExpenses(filterByMonth(expenses, *month))
			fmt.Printf("Total expenses for month %d: $%.2f\n", *month, total)
		}
	case "list":
		printExpense(expenses)

	default:
		fmt.Println("неизвестная команда:", os.Args[1])
		os.Exit(1)
	}
}
