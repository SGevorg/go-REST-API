package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type DataEntry struct {
	ID        string `json:"id,omitempty"`
	ColumnOne string `json:"columntone,omitempty"`
	ColumnRef *RefDataEntry
}

type RefDataEntry struct {
	RefColOne string `json:"refcolone,omitempty"`
	RefColTwo string `json:"refcoltwo,omitempty"`
}

var dataPoints []DataEntry

func GetDataPoints(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(dataPoints)
}

func main() {
	router := mux.NewRouter()
	dataPoints = append(dataPoints, DataEntry{ID: "1", ColumnOne: "ColumneOneVal1", ColumnRef: &RefDataEntry{RefColOne: "Col value one Ref 1", RefColTwo: "Col value two Ref 1"}})
	dataPoints = append(dataPoints, DataEntry{ID: "2", ColumnOne: "ColumneOneVal2", ColumnRef: &RefDataEntry{RefColOne: "Col value one Ref 2", RefColTwo: "Col value two Ref 2"}})

	router.HandleFunc("/data-entries", GetDataPoints).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
