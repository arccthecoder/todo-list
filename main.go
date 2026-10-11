// реализовать отлов ошибок + паники
// пофиксить пользовательский ввод для методов

package main

import (
	"fmt"
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

func CompleteTask(t *Task) {
	t.Completed = true
}

func DeleteTask(t []Task, i int) {
	Tasks = append(Tasks[:i-1], Tasks[i:]...)
}

func main() {
	defer func() {
		fmt.Println("До свидания!")
	}()

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
			var vopros string
			fmt.Print("Выберите действие: ")
			fmt.Scan(&choose)

			switch choose {
			case 1:
				fmt.Print("Введите задачу: ")
				fmt.Scan(&vopros)
				Tasks = append(Tasks, Task{Title: vopros})
				fmt.Println("Задача успешно добавлена")
			case 2:
				if len(Tasks) == 0 {
					fmt.Println("Список задач пуст")
				} else {
					fmt.Println("Все задачи:")
					ShowTasks()
				}
			case 3:
				var numb int
				fmt.Print("Введи номер задачи: ")
				fmt.Scan(&numb)
				CompleteTask(&Tasks[numb-1])
			case 4:
				var numb int
				fmt.Print("Введи номер задачи для удаленя: ")
				fmt.Scan(&numb)
				DeleteTask(Tasks, numb)
			case 5:
				return
			default:
				fmt.Println("Такой задачи нет")
			}
			break

		}
	}
}
