package v1

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/projectmedicloud/MediCloud-Go-Backend/internal/db"
)

// TestTableData represents the structure of your test table data
type TestTableData struct {
	ID        int    `json:"id"`
	TestValue string `json:"test_value"`
}

func GetTestTableData(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling GetTestTableData request")

	database, err := db.Connect()
	if err != nil {
		handleError(w, err, "Error connecting to database")
		return
	}
	defer database.Close()

	log.Println("Successfully connected to the database")

	rows, err := database.Query("SELECT id, test_value FROM test_table")
	if err != nil {
		handleError(w, err, "Error querying test_table")
		return
	}
	defer rows.Close()

	var testData []TestTableData
	for rows.Next() {
		var t TestTableData
		if err := rows.Scan(&t.ID, &t.TestValue); err != nil {
			log.Printf("Error scanning row: %v", err)
			http.Error(w, "Internal Server Error\n"+err.Error(), http.StatusInternalServerError)
			return
		}
		testData = append(testData, t)
	}

	if err != nil {
		handleError(w, err, "Error scanning rows")
		return
	}

	log.Println("Successfully retrieved test data")

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(testData); err != nil {
		handleError(w, err, "Error encoding test data to JSON")
		return
	}
}

func handleError(w http.ResponseWriter, err error, message string) {
	log.Printf("%s: %v", message, err)
	http.Error(w, "Internal Server Error\n"+err.Error(), http.StatusInternalServerError)
}
