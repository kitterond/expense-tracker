package main

import (
	"fmt"
	"time"
)

// summary add delete update
/* что для чего list - это наша история куда будут сохранятся наши траты выраженная слайсом структуры нашей то условно наборчик структур
description - это наш подтип который отвечает за описание тоже для проверки на ошибки нужен и чтобы пользователь объяснил что за трата.
amount - это цена нашей траты или сумма как удобнее
теперь что возвращается.
возвращается первое у нас наш список обновленный или нет взависимости от того ошибка есть там или нет
второе значение отвечает за вывод что именно мы добавили условно вы успешно добавили ваши затраты а именно : ...
последнее это ошибка которая вернется если все было удачно введено или нет под конец, способ протащить ее дальше для проверки к тому же.



*/
func addExpense(list []Expense, description string, amount float64) ([]Expense, Expense, error) {

	if description == "" {
		return list, Expense{}, fmt.Errorf("Описание не может быть пустым")
	}
	if amount <= 0 {
		return list, Expense{}, fmt.Errorf("Не допустимое значение")
	}
	maxID := 0
	for _, e := range list {
		if e.ID > maxID {
			maxID = e.ID
		}
	}
	newExpense := Expense{
		ID:          maxID + 1,
		Description: description,
		Amount:      amount,
		Date:        time.Now(),
	}
	list = append(list, newExpense)
	return list, newExpense, nil
}

func summaryExpenses(list []Expense) float64 {
	var total float64
	for _, e := range list {
		total += e.Amount
	}
	return total
}

func filterByMonth(list []Expense, month int) []Expense {

	m := time.Month(month)
	var result []Expense
	for _, e := range list {
		if e.Date.Month() == m && e.Date.Year() == time.Now().Year() {
			result = append(result, e)
		}
	}
	return result
}

func deleteExpense(list []Expense, ID int) ([]Expense, error) {

	index := -1

	for i, e := range list {
		if e.ID == ID {
			index = i
			break
		}
	}
	if index == -1 {
		return list, fmt.Errorf("Не нашли данную трату")
	}
	list = append(list[:index], list[index+1:]...)
	return list, nil
}

/*
	func updateExpense(list []Expense, ID int, description string, amount float64) ([]Expense, Expense, error) {
		if description == "" {
			return list, Expense{}, fmt.Errorf("Описание не может быть пустым")
		}
		if amount <= 0 {
			return list, Expense{}, fmt.Errorf("Не допустимое значение")
		}
		for i := range list {
			if list[i].ID == ID {
				list[i].Description = description
				list[i].Amount = amount
				return list, list[i], nil
			}
		}
		return list, Expense{}, fmt.Errorf("Не нашли данную трату")
	}
*/
func printExpense(list []Expense) {
	fmt.Print(list)
}
