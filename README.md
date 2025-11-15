# CLI Todo App - A Go Learning Project 🚀

A command-line todo application built with Go to practice and learn fundamental Go programming concepts. This project demonstrates file I/O, data structures, JSON handling, generics, pointers, and building CLI tools.

## 📚 Learning Objectives

This project is designed to help you understand and practice:

- **Structs and Methods**: Custom types with associated functions
- **Pointers vs Values**: Understanding when and why to use pointer receivers
- **Generics**: Using Go's type parameters for reusable code
- **JSON Encoding/Decoding**: Marshaling and unmarshaling data
- **File I/O**: Reading from and writing to files
- **Command-line Flags**: Using the `flag` package for CLI arguments
- **Error Handling**: Proper error checking and propagation
- **Time Handling**: Working with timestamps
- **Third-party Packages**: Integrating external libraries
- **Slices Manipulation**: Appending, deleting, and modifying slice elements

## 🛠️ Tech Stack

- **Language**: Go 1.25.3
- **External Package**: [github.com/aquasecurity/table](https://github.com/aquasecurity/table) - For beautiful table rendering

## 📂 Project Structure

```
CLI-ToDo-App/
│
├── main.go           # Entry point - orchestrates the application flow
├── todo.go           # Todo struct and core business logic (Add, Edit, Delete, Toggle)
├── storage.go        # Generic storage implementation for JSON persistence
├── command.go        # CLI command parsing and execution
├── todos.json        # Data persistence file (auto-generated)
└── go.mod            # Module dependencies
```

## 🧩 Key Go Concepts Demonstrated

### 1. **Structs and Custom Types**
```go
type Todo struct {
    Title       string
    Completed   bool
    CreatedAt   time.Time
    CompletedAt *time.Time  // Pointer to allow nil value
}

type Todos []Todo  // Custom slice type
```

### 2. **Pointer vs Value Receivers**
```go
// Pointer receiver - modifies the original
func (todos *Todos) Add(title string) {
    *todos = append(*todos, newTodo)
}

// Value receiver would only modify a copy!
```

**Why use pointers?**
- Modify the original data structure
- Avoid copying large data
- Memory efficiency

### 3. **Generics (Type Parameters)**
```go
type Storage[T any] struct {
    FileName string
}

func NewStorage[T any](fileName string) *Storage[T] {
    return &Storage[T]{FileName: fileName}
}
```

**What you learn:**
- Creating reusable, type-safe code
- The `any` constraint (equivalent to `interface{}`)
- Generic functions and constructors

### 4. **Error Handling Pattern**
```go
func (todos *Todos) delete(index int) error {
    if err := todos.validateIndex(index); err != nil {
        return err
    }
    // ... rest of logic
    return nil
}
```

**Best practices:**
- Always check errors
- Return errors for caller to handle
- Use descriptive error messages

### 5. **JSON Marshal/Unmarshal**
```go
// Saving to JSON
fileData, err := json.MarshalIndent(data, "", "    ")
os.WriteFile(s.FileName, fileData, 0644)

// Loading from JSON
fileData, err := os.ReadFile(s.FileName)
json.Unmarshal(fileData, data)
```

### 6. **Command-line Flags**
```go
flag.StringVar(&cf.Add, "add", "", "Add a new todo")
flag.IntVar(&cf.Del, "del", -1, "Delete todo by index")
flag.BoolVar(&cf.List, "list", false, "List all todos")
flag.Parse()
```

### 7. **Slice Manipulation**
```go
// Delete element at index
*todos = append(todo[:index], todo[index+1:]...)

// This creates a new slice:
// - todo[:index] = elements before index
// - todo[index+1:] = elements after index
// Combined to skip the element at index
```

### 8. **Working with Time**
```go
CreatedAt: time.Now()
CompletedAt: &completedTime  // Pointer allows nil when not completed

// Formatting
t.CreatedAt.Format(time.RFC1123)
```

## 🚀 Getting Started

### Prerequisites
- Go 1.20 or higher installed
- Basic understanding of command line

### Installation

1. **Clone the repository**
```bash
git clone https://github.com/janak-dev2002/CLI-ToDo-App.git
cd CLI-ToDo-App
```

2. **Install dependencies**
```bash
go mod download
```

3. **Run the application**
```bash
go run ./
```

## 📖 Usage

### View Welcome Screen
```bash
go run ./
```

### Add a New Todo
```bash
go run ./ -add "Learn Go pointers"
go run ./ -add "Build a REST API"
```

### List All Todos
```bash
go run ./ -list
```

### Toggle Todo Completion
```bash
go run ./ -toggle 0
```

### Edit a Todo
```bash
go run ./ -edit 0:"Learn Go pointers and interfaces"
```

### Delete a Todo
```bash
go run ./ -del 0
```

### View Help
```bash
go run ./ -help
```

## 🔨 Building an Executable

Create a standalone executable:

```bash
# For your current OS
go build -o todo

# Run the executable
./todo -list  # Linux/Mac
todo.exe -list  # Windows
```

Cross-compile for different platforms:
```bash
# For Windows
GOOS=windows GOARCH=amd64 go build -o todo.exe

# For Linux
GOOS=linux GOARCH=amd64 go build -o todo

# For macOS
GOOS=darwin GOARCH=amd64 go build -o todo
```

## 🎓 Learning Path

### Beginner Level
1. Understand the basic flow: `main.go` → `command.go` → `todo.go`
2. Study how structs are defined and used
3. Learn about methods and receivers
4. Explore the flag package for CLI parsing

### Intermediate Level
1. Deep dive into pointer vs value receivers
2. Understand slice manipulation and memory implications
3. Study the generic `Storage` implementation
4. Learn JSON encoding/decoding patterns
5. Explore error handling best practices

### Advanced Topics
1. Refactor to use interfaces for better abstraction
2. Add unit tests for each method
3. Implement concurrent operations with goroutines
4. Add database support instead of JSON
5. Create a REST API wrapper around this CLI

## 🧪 Code Exploration Exercises

### Exercise 1: Add a Priority Field
Add a `Priority` field (Low, Medium, High) to the `Todo` struct and update all relevant methods.

### Exercise 2: Implement Filtering
Add a `-filter` flag to list only completed or incomplete todos.

### Exercise 3: Add Due Dates
Include a `DueDate` field and implement sorting by due date.

### Exercise 4: Export to CSV
Create a method to export todos to a CSV file.

### Exercise 5: Unit Testing
Write unit tests for the `Todos` methods using Go's `testing` package.

## 🐛 Common Mistakes to Avoid

1. **Using value receivers when you need to modify data**
   ```go
   // ❌ Wrong - changes won't persist
   func (todos Todos) Add(title string) {
       todos = append(todos, newTodo)
   }
   
   // ✅ Correct - changes persist
   func (todos *Todos) Add(title string) {
       *todos = append(*todos, newTodo)
   }
   ```

2. **Not checking for errors**
   ```go
   // ❌ Wrong
   fileData, _ := os.ReadFile("file.json")
   
   // ✅ Correct
   fileData, err := os.ReadFile("file.json")
   if err != nil {
       return err
   }
   ```

3. **Index out of range**
   - Always validate indices before accessing slice elements
   - The `validateIndex` method demonstrates this pattern

## 📦 Dependencies

- **github.com/aquasecurity/table** - For rendering beautiful tables in the terminal
- Standard library packages:
  - `encoding/json` - JSON marshaling/unmarshaling
  - `flag` - Command-line flag parsing
  - `os` - File operations
  - `time` - Time handling
  - `errors` - Error creation
  - `fmt` - Formatted I/O

## 🤝 Contributing

This is a learning project! Feel free to:
- Add new features
- Improve error handling
- Add tests
- Enhance the UI
- Optimize performance

## 📝 License

This project is open source and available for learning purposes.

## 🙏 Acknowledgments

Special thanks to **Coding with Patrik** for the inspiration and guidance in building this Go learning project!

---

**Happy Learning! 🎉**

*Remember: The best way to learn programming is by building projects. Don't just read the code—modify it, break it, and fix it!*
