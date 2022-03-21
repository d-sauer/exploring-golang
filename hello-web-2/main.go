package main

import (
	"encoding/json"
	"fmt"
	"github.com/moogar0880/problems"
	"log"
	"net/http"
	"time"
)

type data struct {
	Name     string `json:"name"`
	DateTime string `json:"date-time"`
}

func main() {
	fmt.Printf("Starting..")
	http.HandleFunc("/", HelloHandler)

	fmt.Printf("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	newData := data{
		Name:     "Test data",
		DateTime: time.Now().Format(time.RFC3339),
	}

	jsonData, err := json.Marshal(newData)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		httpError := problems.NewDetailedProblem(http.StatusInternalServerError, "Server error")
		jsonHttpError, _ := json.Marshal(httpError)
		w.Write(jsonHttpError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Printf("Json data: \n %s \n", newData)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
