package todolist

type TodoRepository interface {
	GetByUserID(userID uint) ([]Todo, error)
}
