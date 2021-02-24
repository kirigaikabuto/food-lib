package food_lib

import (
	"encoding/json"
	"net/http"
)

type HttpEndpointsFactory interface {
	List() func(w http.ResponseWriter, r *http.Request)
	Create() func(w http.ResponseWriter, r *http.Request)
}

type httpEndpointFactory struct {
	foodService FoodService
}

type customError struct {
	Message string `json:"message"`
}

func NewHttpEndpointFactory(foodService FoodService) HttpEndpointsFactory {
	return &httpEndpointFactory{foodService: foodService}
}

func (httpFact *httpEndpointFactory) List() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		listReq := &ListCommand{}
		if r.Header.Get("Content-Type") == "application/json" {
			err := json.NewDecoder(r.Body).Decode(listReq)
			if err != nil {
				respondJSON(w, http.StatusInternalServerError, &customError{err.Error()})
				return
			}
		}
		//count, err := strconv.ParseInt(r.URL.Query().Get("count"), 10, 64)
		//if err != nil {
		//	respondJSON(w, http.StatusInternalServerError, &customError{err.Error()})
		//	return
		//}
		data, err := listReq.Exec(httpFact.foodService)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, &customError{err.Error()})
			return
		}
		respondJSON(w, http.StatusOK, data)
	}
}
func (httpFact *httpEndpointFactory) Create() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		createCmd := &CreateCommand{}
		if r.Header.Get("Content-Type") == "application/json" {
			err := json.NewDecoder(r.Body).Decode(createCmd)
			if err != nil {
				respondJSON(w, http.StatusInternalServerError, &customError{err.Error()})
				return
			}
		}
		data, err := createCmd.Exec(httpFact.foodService)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, &customError{err.Error()})
			return
		}
		respondJSON(w, http.StatusCreated, data)
	}
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}
