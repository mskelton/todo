package storage

import (
	"database/sql"
	"errors"

	"github.com/mskelton/todo/internal/models"
)

func ListTasks(filters []Filter) ([]models.Task, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	tx := WithFilters(db, filters).Find(&tasks)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if len(tasks) == 0 {
		return nil, errors.New("No tasks match filters")
	}

	return tasks, nil
}

func Add(task models.Task) (int64, error) {
	return 0, nil
	// data, err := json.Marshal(task)
	// if err != nil {
	// 	return 0, fmt.Errorf("Failed to add task: %w", err)
	// }
	//
	// conn, err := connect()
	// if err != nil {
	// 	return 0, fmt.Errorf("Failed to add task: %w", err)
	// }
	//
	// _, err = conn.Exec(
	// 	"INSERT INTO tasks (id, template_id, data) VALUES (?, ?, ?)",
	// 	task.Id,
	// 	task.TemplateId,
	// 	data,
	// )
	// if err != nil {
	// 	return 0, fmt.Errorf("Failed to add task: %w", err)
	// }
	//
	// // Add an id assignment for the newly created task
	// res, err := conn.Exec(
	// 	"INSERT INTO assignments VALUES ((select max(id) + 1 from assignments), ?)",
	// 	task.Id,
	// )
	// if err != nil {
	// 	return 0, fmt.Errorf("Failed to add task assignment: %w", err)
	// }
	//
	// id, err := res.LastInsertId()
	// if err != nil {
	// 	return 0, fmt.Errorf("Failed to get last insert id: %w", err)
	// }
	//
	// // Return the id of the newly created task. Thankfully SQLite handles this
	// // automatically with `LastInsertId()` since we are using a numeric id.
	// return id, nil
}

func Count(filters []Filter) (int, error) {
	return 0, nil
	// conn, err := connect()
	// if err != nil {
	// 	return 0, fmt.Errorf("Failed to count tasks: %w", err)
	// }
	//
	// builder := sql_builder.New().
	// 	Select("count(tasks.id)").
	// 	From("tasks").
	// 	Join("assignments", "tasks.id = assignments.task_id")
	//
	// for _, filter := range filters {
	// 	builder.Filter(filter)
	// }
	//
	// debug := os.Getenv("DEBUG") != ""
	// if debug {
	// 	log.Println(builder.SQL())
	// }
	//
	// row := conn.QueryRow(builder.SQL())
	// if row.Err() != nil {
	// 	return 0, fmt.Errorf("Failed to count tasks: %w", row.Err())
	// }
	//
	// var count int
	// err = row.Scan(&count)
	// if row.Err() != nil {
	// 	return 0, fmt.Errorf("Failed to count tasks: %w", row.Err())
	// }
	//
	// return count, nil
}

func getIds(conn *sql.DB, filters []Filter) ([]int, error) {
	// builder := sql_builder.New().
	// 	Select("assignments.id").
	// 	From("tasks").
	// 	Join("assignments", "tasks.id = assignments.task_id")
	//
	// for _, filter := range filters {
	// 	builder.Filter(filter)
	// }
	//
	// debug := os.Getenv("DEBUG") != ""
	// if debug {
	// 	log.Println(builder.SQL())
	// }
	//
	// res, err := conn.Query(builder.SQL())
	// if err != nil {
	// 	return nil, fmt.Errorf("Failed to get task ids: %w", err)
	// }
	//
	// var ids []int
	// for res.Next() {
	// 	var id int
	// 	err = res.Scan(&id)
	//
	// 	if err != nil {
	// 		return nil, fmt.Errorf("Failed to get task id: %w", err)
	// 	}
	//
	// 	ids = append(ids, id)
	// }

	return nil, nil
}

type QueryEdit struct {
	Path  string
	Value string
}

func Edit(filters []Filter, edits []QueryEdit) ([]int, error) {
	return nil, nil
	// conn, err := connect()
	// if err != nil {
	// 	return nil, fmt.Errorf("Failed to edit tasks: %w", err)
	// }
	//
	// builder := sql_builder.New().Update("tasks")
	// var params []any
	//
	// for _, edit := range edits {
	// 	params = append(params, edit.Value)
	// 	builder.Set(fmt.Sprintf("data = json_set(data, '$.%s', ?)", edit.Path))
	// }
	//
	// for _, filter := range filters {
	// 	builder.Filter(filter)
	// }
	//
	// debug := os.Getenv("DEBUG") != ""
	// if debug {
	// 	log.Println(builder.SQL())
	// }
	//
	// _, err = conn.Exec(builder.SQL(), params...)
	// if err != nil {
	// 	return nil, fmt.Errorf("Failed to edit tasks: %w", err)
	// }
	//
	// return getIds(conn, filters)
}

func Delete(filters []Filter) ([]int, error) {
	// conn, err := connect()
	// if err != nil {
	// 	return nil, fmt.Errorf("Failed to delete tasks: %w", err)
	// }
	//
	// builder := sql_builder.New().Delete("tasks")
	//
	// for _, filter := range filters {
	// 	builder.Filter(filter)
	// }
	//
	// debug := os.Getenv("DEBUG") != ""
	// if debug {
	// 	log.Println(builder.SQL())
	// }
	//
	// _, err = conn.Exec(builder.SQL())
	// if err != nil {
	// 	return nil, fmt.Errorf("Failed to delete tasks: %w", err)
	// }
	//
	// return getIds(conn, filters)
	return nil, nil
}
