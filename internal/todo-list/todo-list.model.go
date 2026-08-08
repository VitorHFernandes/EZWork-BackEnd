package todolist

type Todo struct {
	ID          uint
	UserID      uint
	Title       string
	Description string
	DtInit      string
	DtEnd       string
	IsCompleted bool
}
