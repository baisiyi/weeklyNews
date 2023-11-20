package models

import (
	"encoding/json"
	"net/http"
)

// JSONErrResponse TODO
func JSONErrResponse(r *http.Request, w http.ResponseWriter, errCode string, err error) {
	var rspErr ResponseErr
	rspErr.RequestId = r.FormValue("RequestId")
	rspErr.Error.Code = errCode
	rspErr.Error.Message = err.Error()

	json.NewEncoder(w).Encode(rspErr)
}

type ResponseErr struct {
	Error     APIError
	RequestId string
	Ext       string
}

type APIError struct {
	Code    string
	Message string
}
