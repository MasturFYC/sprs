package controller

import (
	"database/sql"
	"fmt"
	"log"

	"fyc.com/sprs/models"

	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

// func GetTasks(c *gin.Context) {
//

// 	tasks, err := getAllTasks()

// 	if err != nil {
// 		log.Fatalf("Unable to get all tasks. %v", err)
// 	}

// 	c.JSON(http.StatusOK, &tasks)
// }

func GetTask(c *gin.Context) {

	// id = order id
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		//log.Fatalf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	tasks, err := getTask(db, &id)

	if err != nil {
		//		log.Fatalf("Unable to get task. %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &tasks)
}

func DeleteTask(c *gin.Context) {

	// id = order id
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		//	log.Fatalf("Unable to convert the string into int.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	deletedRows := deleteTask(db, &id)

	msg := fmt.Sprintf("Task deleted successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := Response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
}

func CreateTask(c *gin.Context) {

	var task models.Task

	err := c.BindJSON(&task)

	if err != nil {
		//log.Printf("Unable to decode the request body.  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	rowAffected, err := createTask(db, &task)

	if err != nil {
		//log.Printf("Nama tasks tidak boleh sama.  %v", err)
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf("Task created successfully. Total rows/record affected %v", rowAffected)

	// format the reponse message
	res := Response{
		ID:      rowAffected,
		Message: msg,
	}

	c.JSON(http.StatusOK, &res)
}

func UpdateTask(c *gin.Context) {

	// create the postgres db connection

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var task models.Task

	err := c.BindJSON(&task)

	if err != nil {
		log.Printf("Unable to decode the request body.  %v", err)
		c.JSON(http.StatusRequestedRangeNotSatisfiable, gin.H{"error": err.Error()})
		return
	}

	db := c.Keys["db"].(*sql.DB)
	updatedRows, err := updateTask(db, &id, &task)

	if err != nil {
		log.Printf("Unable to update task.  %v", err)
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf("Task updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := Response{
		ID:      updatedRows,
		Message: msg,
	}

	// send the response
	c.JSON(http.StatusOK, res)
}

func getTask(db *sql.DB, id *int64) (models.Task, error) {

	var task models.Task

	var sqlStatement = `SELECT 
		order_id, descriptions, period_from, period_to,
		recipient_name, recipient_position,
		giver_name, giver_position
	FROM tasks
	WHERE order_id=$1`

	rs := db.QueryRow(sqlStatement, id)

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

// 	rs, err := db.Query(sqlStatement)

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

func deleteTask(db *sql.DB, id *int64) int64 {
	// create the delete sql query
	sqlStatement := `DELETE FROM tasks WHERE order_id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

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

func createTask(db *sql.DB, t *models.Task) (int64, error) {

	sqlStatement := `INSERT INTO tasks (
		order_id, descriptions, period_from, period_to,
		recipient_name, recipient_position,
		giver_name, giver_position
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	res, err := db.Exec(sqlStatement,
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
		log.Printf("Unable to create task. %v", err)
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()

	return rowsAffected, err
}

func updateTask(db *sql.DB, id *int64, t *models.Task) (int64, error) {

	sqlStatement := `UPDATE tasks SET
		descriptions=$2, period_from=$3, period_to=$4,
		recipient_name=$5, recipient_position=$6,
		giver_name=$7, giver_position=$8
	WHERE order_id=$1`

	res, err := db.Exec(sqlStatement,
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
		log.Printf("Unable to update task. %v", err)
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	return rowsAffected, err
}
