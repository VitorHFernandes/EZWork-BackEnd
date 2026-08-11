package todolist

type TodoService struct {
	repository TodoRepository
}

func NewTodoService(repository TodoRepository) *TodoService {
	return &TodoService{
		repository: repository,
	}
}

func (s *TodoService) GetTodoList(userID uint) ([]Todo, error) {
	todoList, err := s.repository.GetByUserID(userID)
	if err != nil {
		return nil, err
	}
	return todoList, nil
}
