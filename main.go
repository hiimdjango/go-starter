package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hiimdjango/gostarter/models"
)

func main() {
	transactions := `
	[
		{
			"name":      "Salary",
			"amount":    2500,
			"category":  "Other",
			"frequency": "Onetime",
			"type":      "Income",
			"date":      "2022-01-01"
		},
		{
			"name":      "Rent",
			"amount":    1500,
			"category":  "Other",
			"frequency": "MONTHLY",
			"type":      "Expense",
			"date":      "2023-02-01"
		}
	]
	`

	var unmarshalled []models.TransactionUnmarshalled
	err := json.Unmarshal([]byte(transactions), &unmarshalled)

	marshaled, _ := json.Marshal(unmarshalled)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, string(marshaled))
		if err != nil {
			fmt.Println(err)
		}
	})

	http.ListenAndServe(":8080", nil)

}
