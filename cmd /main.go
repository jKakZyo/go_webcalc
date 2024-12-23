package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/jKakZyo/go_webcalc/internal/calculator"
)

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result interface{} `json:"result,omitempty"`
	Error  string      `json:"error,omitempty"`
}

func handleCalculate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Error: "Invalid JSON format"})
		return
	}

	if req.Expression == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(Response{Error: "Expression is not valid"})
		return
	}

	result, err := calculator.Calc(req.Expression)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(Response{Error: "Expression is not valid"})
		return
	}

	if result == float64(int(result)) {
		json.NewEncoder(w).Encode(Response{Result: int(result)})
	} else {
		json.NewEncoder(w).Encode(Response{Result: result})
	}
}

func main() {

	http.HandleFunc("/api/v1/calculate", handleCalculate)

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
