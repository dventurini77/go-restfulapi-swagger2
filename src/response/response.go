package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type JSONSuccessResult struct {
	Code    int         `json:"code" example:"200"`
	Message string      `json:"message" example:"Success"`
	Data    interface{} `json:"data"`
}

type JSONBadRequest struct {
	Code    int         `json:"code" example:"400"`
	Message string      `json:"message" example:"Wrong parameter"`
	Data    interface{} `json:"data"`
}

type JSONIntServerErrReqResult struct {
	Code    int         `json:"code" example:"500"`
	Message string      `json:"message" example:"Error Database"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(w http.ResponseWriter, data interface{}) {

	body, err := json.Marshal(JSONSuccessResult{Code: 200, Message: "Success", Data: data})
	if err != nil {
		log.Printf("Failed to encode a JSON response: %v", err)
		//http.Error(w, err.Error(), 500)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(body)
	if err != nil {
		log.Printf("Failed to write the response body: %v", err)
		return
	}
}

func FailResponse(w http.ResponseWriter, respCode int, message string) {

	if respCode == http.StatusInternalServerError {
		_, err := json.Marshal(JSONSuccessResult{Code: 500, Message: "Internal Server Error", Data: nil})
		if err != nil {
			log.Printf("Failed to encode a JSON response: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	body, err := json.Marshal(JSONSuccessResult{Code: respCode, Message: message, Data: nil})
	if err != nil {
		log.Printf("Failed to encode a JSON response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	_, err = w.Write(body)
	if err != nil {
		log.Printf("Failed to write the response body: %v", err)
		return
	}
}
