package todolist

import "database/sql"

type MySQLTodoRepository struct {
	db *sql.DB
}

func NewMySQLTodoRepository(db *sql.DB) *MySQLTodoRepository {
	return &MySQLTodoRepository{
		db: db,
	}
}

func (r *MySQLTodoRepository) GetByUserID(userID uint) ([]Todo, error) {
	rows, err := r.db.Query(`
		SELECT 
			* 
		FROM tb_todo_list 
		WHERE userId = ?
	`, userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		var todo Todo

		err := rows.Scan(
			&todo.ID,
			&todo.UserID,
			&todo.Title,
			&todo.Description,
			&todo.DtInit,
			&todo.DtEnd,
			&todo.IsCompleted,
		)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}
