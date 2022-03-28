package middleware

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type Response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func createConnection() *sql.DB {
	// load .env file
	//err := godotenv.Load("/hdd/go-lang/hello/.env")

	// Open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetMaxIdleConns(25)
	db.SetMaxOpenConns(25)
	db.Ping()

	//defer db.DB().Close()

	return db
}

func EnableCors(w *http.ResponseWriter) {
	//	(*w).Header().Set("Access-Control-Allow-Origin", "http://pixel.id:8080")
	//(*w).Header().Set("Context-Type", "application/x-www-form-urlencoded")
	(*w).Header().Set("Content-Type", "application/json")
	//(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func LoadConnection() func() *sql.DB {
	db := createConnection()

	fmt.Println("Connecting to database...")

	return func() *sql.DB {
		err := (*db).Ping()
		if err != nil {
			db = createConnection()
		}
		return db
	}
}

func CombineString(a string, b string) string {
	if len(b) == 0 {
		return a
	}
	return a + ", " + b
}

var Sql = LoadConnection()

func NestQuerySingle(query string) string {
	return fmt.Sprintf(`(SELECT row_to_json(x) FROM (%s) x)`, query)
}

func NestQuery(query string) string {
	return fmt.Sprintf(`COALESCE((
        SELECT array_to_json(array_agg(row_to_json(x)))
        FROM (%s) x), '[]')`, query)
}

func GetMonthName(id int, isShort bool) string {
	months := [12]string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "Nopember", "Desember"}
	//months2 := [12]string{"Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Agus", "Sep", "Okt", "Nop", "Des"}
	if isShort {
		return months[id][0:3]
	}
	return months[id]
}
