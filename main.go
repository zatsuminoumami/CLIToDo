package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const dataFile = "tasks.json"

func main() {
	tasks, _ := LoadTasks(dataFile)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n1. タスク表示\n2. タスク追加\n3. 完了マーク\n4. 終了")
		fmt.Print("選択: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			for _, t := range tasks {
				status := " "
				if t.Completed {
					status = "✓"
				}
				fmt.Printf("[%s] %d: %s\n", status, t.ID, t.Title)
			}
		case "2":
			fmt.Print("タスク名: ")
			scanner.Scan()
			title := scanner.Text()
			id := len(tasks) + 1
			tasks = append(tasks, Task{ID: id, Title: title, Completed: false})
			SaveTasks(dataFile, tasks)
			fmt.Println("タスクを追加しました")
		case "3":
			fmt.Print("完了したタスクID: ")
			scanner.Scan()
			idStr := scanner.Text()
			id, _ := strconv.Atoi(idStr)
			for i, t := range tasks {
				if t.ID == id {
					tasks[i].Completed = true
					break
				}
			}
			SaveTasks(dataFile, tasks)
			fmt.Println("完了マークをつけました")
		case "4":
			fmt.Println("終了します")
			return
		default:
			fmt.Println("無効な選択です")
		}
	}
}
