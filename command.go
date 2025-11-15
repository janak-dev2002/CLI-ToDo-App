package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
	Help   bool
}

func NewCmdFlags() *CmdFlags {

	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo specify title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index & specify a new title. id:new_title")
	flag.IntVar(&cf.Del, "del", -1, "Specify todo by index to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify todo by index to toggle complete true/false")
	flag.BoolVar(&cf.List, "list", false, "List all todos")
	flag.BoolVar(&cf.Help, "help", false, "List all commands")

	flag.Parse()

	return &cf
}

func printIntro() {
	fmt.Println("\n╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║            CLI Todo App - Task Management Tool            ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
	fmt.Println("\nWelcome to CLI Todo App! A simple and efficient command-line")
	fmt.Println("tool to manage your daily tasks and stay organized.")
	fmt.Println("\nHow to Use:")
	fmt.Println("-----------")
	fmt.Println("Run the app with one of the following commands:")
	fmt.Println()
	fmt.Println("  go run ./ -add \"Buy a coffee\"")
	fmt.Println("  go run ./ -list")
	fmt.Println("  go run ./ -toggle 0")
	fmt.Println("  go run ./ -edit 0:\"Buy groceries and cook dinner\"")
	fmt.Println("  go run ./ -del 0")
	fmt.Println("\nFor a full list of commands, run: go run ./ -help")
	fmt.Println()
}

func printCommands() {
	fmt.Println("\n╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║                    Available Commands                     ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
	fmt.Println("\n-add <title>        : Add a new todo with the specified title")
	fmt.Println("-edit <id:title>    : Edit a todo by index and specify a new title (format: id:new_title)")
	fmt.Println("-del <id>           : Delete a todo by index")
	fmt.Println("-toggle <id>        : Toggle todo completion status by index")
	fmt.Println("-list               : List all todos")
	fmt.Println("-help               : Display this help message")
	fmt.Println("\nExamples:")
	fmt.Println("---------")
	fmt.Println("  go run ./ -add \"Complete project report\"")
	fmt.Println("  go run ./ -list")
	fmt.Println("  go run ./ -toggle 0")
	fmt.Println("  go run ./ -edit 0:\"Finish project report\"")
	fmt.Println("  go run ./ -del 0")
	fmt.Println()
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.Help:
		printCommands()
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.Add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error: Invalid format for edit. Please use index:new_title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: Invalid index for edit.")
			os.Exit(1)

		}
		todos.edit(index, parts[1])
	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)

	case cf.Del != -1:
		todos.delete(cf.Del)

	default:
		printIntro()
	}
}
