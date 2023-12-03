package utils

import (
	"encoding/json"
	"net/http"
	"online-library/model"
	"strconv"
)

// setup response data
func ReturnResponse(w http.ResponseWriter, statusCode int, respMsg string, retData []*model.RetData) {
	respData := &model.Response{
		Status:  strconv.Itoa(statusCode),
		Message: respMsg,
		Data:    retData,
	}

	// convert data into json and send as response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(respData)
}
