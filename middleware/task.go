package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"fyc.com/sprs/models"

	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

// func GetTasks(w http.ResponseWriter, r *http.Request) {
// 	EnableCors(&w)

// 	tasks, err := getAllTasks()

// 	if err != nil {
// 		log.Fatalf("Unable to get all tasks. %v", err)
// 	}

// 	json.NewEncoder(w).Encode(&tasks)
// }

func GetTask(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	params := mux.Vars(r)

	// id = order id
	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	tasks, err := getTask(&id)

	if err != nil {
		log.Fatalf("Unable to get task. %v", err)
	}

	json.NewEncoder(w).Encode(&tasks)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	// id = order id
	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRows := deleteTask(&id)

	msg := fmt.Sprintf("Task deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var task models.Task

	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	rowAffected, err := createTask(&task)

	if err != nil {
		log.Fatalf("Nama tasks tidak boleh sama.  %v", err)
	}

	msg := fmt.Sprintf("Task created successfully. Total rows/record affected %v", rowAffected)

	// format the reponse message
	res := Response{
		ID:      rowAffected,
		Message: msg,
	}

	json.NewEncoder(w).Encode(&res)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {

	EnableCors(&w)

	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// create the postgres db connection

	params := mux.Vars(r)

	id, _ := strconv.ParseInt(params["id"], 10, 64)

	var task models.Task

	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	updatedRows := updateTask(&id, &task)

	msg := fmt.Sprintf("Task updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func getTask(id *int64) (models.Task, error) {

	var task models.Task

	var sqlStatement = `SELECT 
		descriptions, period_from, period_to,
		recipient_name, recipient_position,
		giver_name, giver_position
	FROM tasks
	WHERE order_id=$1`

	rs := Sql().QueryRow(sqlStatement, id)

	err := rs.Scan(
		&task.OrderID,
		&task.Descriptions,
		&task.PeriodFrom,
		&task.PeriodTo,
		&task.RecipientName,
		&task.RecipientPosition,
		&task.GiverName,
		&task.GiverPosition,
	)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return task, nil
	case nil:
		return task, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return task, err
}

// func getAllTasks() ([]models.Task, error) {

// 	var tasks []models.Task

// 	var sqlStatement = `SELECT
// 		descriptions, period_from, period_to,
// 		recipient_name, recipient_position,
// 		giver_name, giver_position
// 	FROM tasks`

// 	rs, err := Sql().Query(sqlStatement)

// 	if err != nil {
// 		log.Fatalf("Unable to execute tasks query %v", err)
// 	}

// 	defer rs.Close()

// 	for rs.Next() {
// 		var task models.Task

// 		err := rs.Scan(
// 			&task.OrderID,
// 			&task.Descriptions,
// 			&task.PeriodFrom,
// 			&task.PeriodTo,
// 			&task.RecipientName,
// 			&task.RecipientPosition,
// 			&task.GiverName,
// 			&task.GiverPosition,
// 		)

// 		if err != nil {
// 			log.Fatalf("Unable to scan the row. %v", err)
// 		}

// 		tasks = append(tasks, task)
// 	}

// 	return tasks, err
// }

func deleteTask(id *int64) int64 {
	// create the delete sql query
	sqlStatement := `DELETE FROM tasks WHERE order_id=$1`

	// execute the sql statement
	res, err := Sql().Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to delete task. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	return rowsAffected
}

func createTask(t *models.Task) (int64, error) {

	sqlStatement := `INSERT INTO tasks (
		order_id, descriptions, period_from, period_to,
		recipient_name, recipient_position,
		giver_name, giver_position
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	res, err := Sql().Exec(sqlStatement,
		t.OrderID,
		t.Descriptions,
		t.PeriodFrom,
		t.PeriodTo,
		t.RecipientName,
		t.RecipientPosition,
		t.GiverName,
		t.GiverPosition,
	)

	if err != nil {
		log.Fatalf("Unable to create task. %v", err)
	}

	if err != nil {
		log.Fatalf("Unable to create customer. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Unable to create customer. %v", err)
	}

	return rowsAffected, err
}

func updateTask(id *int64, t *models.Task) int64 {

	sqlStatement := `UPDATE tasks SET
		descriptions=$2, period_from=$3, period_to=$4,
		recipient_name=$5, recipient_position=$6,
		giver_name=$7, giver_position=$8
	WHERE order_id=$1`

	res, err := Sql().Exec(sqlStatement,
		id,
		t.Descriptions,
		t.PeriodFrom,
		t.PeriodTo,
		t.RecipientName,
		t.RecipientPosition,
		t.GiverName,
		t.GiverPosition,
	)

	if err != nil {
		log.Fatalf("Unable to update task. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while updating task. %v", err)
	}

	return rowsAffected
}
