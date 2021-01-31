package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"wallesterTestApp/models"

	_ "github.com/lib/pq" // postgres golang driver
)

// response format
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "martin"
	password = "qweasdzxc"
	dbname   = "application"
)

// create connection with postgres db
func createConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Open the connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// Check the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}

// GetAllCustomers will return all the customers
func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// get all the users in the db
	customers, err := getAllCustomers()
	if err != nil {
		log.Fatalf("Unable to get all customers. %v", err)
	}

	// send all the users as response
	json.NewEncoder(w).Encode(customers)
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ts, err := template.ParseFiles("./templates/index.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

//private
func getAllCustomers() ([]models.Customer, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	var customers []models.Customer

	// create the select sql query
	sqlStatement := `SELECT * FROM customers`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var customer models.Customer

		// unmarshal the row object to customer
		err = rows.Scan(
			&customer.ID,
			&customer.FirstName,
			&customer.LastName,
			&customer.BirthDate,
			&customer.Gender,
			&customer.Email,
			&customer.Address,
		)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the customer in the customers slice
		customers = append(customers, customer)

	}

	return customers, err
}
