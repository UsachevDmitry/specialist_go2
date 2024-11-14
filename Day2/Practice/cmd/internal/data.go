package internal

import (
	//"time"
	"slices"
	"database/sql"
	"fmt"
	"os"
	_ "modernc.org/sqlite"
	"log"
)

var DB []Task

type Task struct {
	ID int `json:"id"`
	Text string `json:"text"`
	Tags string `json:"tags"`
	Due string `json:"due"`
}

func init() {
	task := Task{
		ID: 1,
		Text: "Task 1",
		Tags: "Tag1",
		Due: "2006-01-02 15:00:00",
	}
	DB = append(DB, task)
}

func FindTaskByID(id int) (Task, bool) {
	var task Task
	var found bool
	for _, t := range DB {
		if t.ID == id{
			task = t
			found = true
			break
		}
	}
	return task, found
}

func DeleteTaskById(id int) bool {
	for idx, t := range DB {
		if t.ID == id {
			DB = slices.Delete(DB, idx, idx+1)
			return true
		}
	}
	return false
}

func DeleteTasks() {
			DB = nil
}


var Qdb *sql.DB = nil
func Sqlite() {

	// Remove "todo.db" if file exists
	os.Remove("./todo.db")

	// Open database connection
	// github.com/mattn/go-sqlite3
	// db, err := sql.Open("sqlite3", "./todo.db")

	// modernc.org/sqlite
	Qdb, err := sql.Open("sqlite", "./todo.db")

	// check connection result
	if err != nil {
		log.Fatal(err)
	}

	// clone connection
	defer Qdb.Close()

	{ // Create table block
		// SQL statement to create table
		sqlStmt := `
		CREATE TABLE IF NOT EXISTS task (id INTEGER PRIMARY KEY AUTOINCREMENT, text TEXT, tags TEXT, due TEXT);
		`
		// Execute SQL statement
		_, err = Qdb.Exec(sqlStmt)
		if err != nil {
			GlobalSugar.Infoln("%q: %s\n", err, sqlStmt)
		}
	}

	{	// Create records block
		// Begin transaction
		tx, err := Qdb.Begin()
		if err != nil {
			GlobalSugar.Fatalw(err.Error())
		}
		// Prepare SQL statement than can be reused. Char "?" for SQLite, char "%" for MySQL, PostgreSQL
		stmt, err := tx.Prepare("INSERT INTO task(id, text, tags, due) VALUES(?, ?, ?, ?)")
		if err != nil {
			GlobalSugar.Fatalw(err.Error())
		}
		// close prepared statement before exiting program
		defer stmt.Close()

		// Create empty slice to store our todos
		tasks := []*Task{}
		// Create tasks
		tasks = append(tasks, &Task{ID: 1, Text: "Learn REST API", Tags: "teacher", Due: "0"})
		//tasks = append(tasks, &Task{ID: 2, Text: "Make practice", Tags: "students", Due: "0"})

		for i := range tasks {
			// Insert records
			// Execute statement for each task
			_, err = stmt.Exec(tasks[i].ID, tasks[i].Text, tasks[i].Tags, tasks[i].Due)
			if err != nil {
				GlobalSugar.Fatalw(err.Error())
			}
		}
		// Commit the transaction
		if err := tx.Commit(); err != nil {
			GlobalSugar.Fatalw(err.Error())
		}
	}

	{ // Read records block
		stmt, err := Qdb.Prepare("SELECT id, text, tags, due FROM task")
		if err != nil {
			GlobalSugar.Fatalw(err.Error())
		}
		defer stmt.Close()

		rows, err := stmt.Query(0)
		if err != nil {
			GlobalSugar.Fatalw(err.Error())
		}
		// Rows must be closed
		defer rows.Close()

		for rows.Next() {
			var id int
			var text string
			var tags string
			var due string
			// use pointers to get data
			err = rows.Scan(&id, &text, &tags, &due)
			if err != nil {
				GlobalSugar.Fatalw(err.Error())
			}
			fmt.Println(id, text, tags, due)
		}
		// Check error, that can occur during iteration
		err = rows.Err()
		if err != nil {
			GlobalSugar.Fatalw(err.Error())
		}

	}

	{ 
		// Update block through UPDATE SQL statement 	
	}
	{ 
		// Delete block through DELETE SQL statement 	
	}


}

func AddTaskInDB(Qdb *sql.DB, task Task) error {
	query := `INSERT INTO task (id, text, tags, due) VALUES (?, ?, ?, ?)`
	stmt, err := Qdb.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(task.ID, task.Text, task.Tags, task.Due)
	if err != nil {
		return err
	}

	return nil
}