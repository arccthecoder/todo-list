package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	Title     string
	Completed bool
}

func Empty() {
	fmt.Println("")
}

func ShowTasks() {
	for n, v := range Tasks {
		if v.Completed {
			fmt.Println(n+1, v.Title, "✔️")
		} else {
			fmt.Println(n+1, v.Title, "❌")
		}
	}
}

var Tasks = []Task{}

func CompleteTask(t []Task, i int) error {
	if i < 0 || i > len(t) {
		return errors.New("Такой задачи нет!")
	}
	t[i].Completed = true
	return nil
}

func DeleteTask(t []Task, i int) error {
	if i < 0 || i > len(t) {
		return errors.New("Такой задачи нет!")
	}
	Tasks = append(Tasks[:i-1], Tasks[i:]...)
	return nil
}

func main() {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println("Поймал панику:", p)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("=== To-Do List ===")
		Empty()
		fmt.Println("1. Добавить задачу")
		fmt.Println("2. Показать задачи")
		fmt.Println("3. Выполнить задачу")
		fmt.Println("4. Удалить задачу")
		fmt.Println("5. Выйти")

		for {
			var choose int
			fmt.Print("Выберите действие: ")
			fmt.Scan(&choose)

			switch choose {
			case 1:
				fmt.Print("Введите задачу: ")
				if scanner.Scan() {
					input := scanner.Text()
					Tasks = append(Tasks, Task{Title: input})
					fmt.Println("Задача успешно добавлена")
				}
			case 2:
				if len(Tasks) == 0 {
					fmt.Println("Список задач пуст")
				} else {
					fmt.Println("Все задачи:")
					ShowTasks()
				}
			case 3:
				fmt.Print("Введи номер задачи: ")
				if scanner.Scan() {
					input := scanner.Text()
					num, _ := strconv.Atoi(input)
					comp := CompleteTask(Tasks, num-1)
					if comp != nil {
						fmt.Println("Ошибка!", comp.Error())
					} else {
						fmt.Println("Задача выполнена!")
					}
				}
			case 4:
				fmt.Print("Введи номер задачи для удаленя: ")
				if scanner.Scan() {
					input := scanner.Text()
					num, err := strconv.Atoi(input)
					if err != nil {
						fmt.Println("Ошибка: введите корректный номер задачи")
						continue
					}
					del := DeleteTask(Tasks, num)
					if del != nil {
						fmt.Println("Ошибка!", del.Error())
					} else {
						fmt.Println("Задача успешно удалена")
					}
				}
				if err := scanner.Err(); err != nil {
					fmt.Fprintln(os.Stderr, "Критическая ошибка сканера:", err)
				}
			case 5:
				fmt.Println("До свидания!")
				return
			default:
				fmt.Println("Такого действия нет")
			}
			break

		}
	}
}
