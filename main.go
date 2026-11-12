package main

import (
	"log"
)

func main() {
	expenses, err:= loadExpenses()
	if err != nil {
		log.Fatalf("Не удалось загрузить расходы: %v", err)
	}
  _ = expenses
}