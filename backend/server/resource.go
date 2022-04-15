package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/robertjanetzko/LegendsBrowser2/backend/model"
)

type Info struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func RegisterResource[T model.Named](router *mux.Router, resourceName string, resources map[int]T) {
	list := func(w http.ResponseWriter, r *http.Request) {
		values := make([]Info, 0, len(resources))
		for _, v := range resources {
			values = append(values, Info{Id: v.Id(), Name: v.Name()})

		}
		json.NewEncoder(w).Encode(values)
	}

	get := func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			fmt.Println(err)
		}

		json.NewEncoder(w).Encode(resources[id])
	}

	router.HandleFunc(fmt.Sprintf("/api/%s", resourceName), list).Methods("GET")
	router.HandleFunc(fmt.Sprintf("/api/%s/{id}", resourceName), get).Methods("GET")
}