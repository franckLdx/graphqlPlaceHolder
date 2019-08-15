package httpClient

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	UserID    int    `json:"userId"`
}

const TodoResource Resource = "todos"

func FetchTodos() (*[]Todo, error) {
	var todos []Todo
	err := FetchResources(TodoResource, &todos)
	return &todos, err
}

func FetchTodo(todoId int) (*Todo, error) {
	var todo Todo
	err := FetchResource(TodoResource, todoId, &todo)
	return &todo, err
}
