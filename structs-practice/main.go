package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

// emebeds displayer and saver interfaces
type outputtable interface {
	Display()
	saver
}

func main() {

	printSomething("Hello, World!")
	printSomething(42)
	printSomething(3.14)
	printSomething(true)
	printSomething([]string{"apple", "banana", "cherry"})

	title, content := getNoteData()
	todoText := getTodoData()

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	printSomething(todo)

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		fmt.Println("Error saving todo:", err)
		return
	}
	err = outputData(userNote)
	if err != nil {
		fmt.Println("Error saving note:", err)
		return
	}
	fmt.Println("All data saved successfully!")

}

// any value allowed
func printSomething(value interface{}) {
	// extract the type of value
	typedValue, ok := value.(int)

	if ok {
		fmt.Println("Integer:", typedValue)
		return
	}

	// switch value.(type) {
	// case string:
	// 	fmt.Println("String:", value)
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// case bool:
	// 	fmt.Println("Boolean:", value)
	// case []string:
	// 	fmt.Println("String slice:", value)
	// default:
	// 	// Use fmt.Printf to print the value with its type
	// 	fmt.Printf("Unknown type %T: %v\n", value, value)
	// }
}

func outputData(o outputtable) error {
	o.Display()
	return saveData(o)
}

func saveData(s saver) error {
	err := s.Save()

	if err != nil {
		fmt.Println("Saving data failed:", err)
		return err
	}

	fmt.Println("Saving data succeeded!")
	return nil
}

func getTodoData() string {
	text := getUserInput("Todo text:")

	if text == "" {
		fmt.Println("Invalid input.")
		return ""
	}

	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
