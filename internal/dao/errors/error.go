package errors

import (
	"encoding/json"
)

type DBError struct {
	Code string
	Msg  string
}

func (e *DBError) Error() string {
	err, _ := json.Marshal(e)
	return string(err)
}

func NewDBError(code, msg string) *DBError {
	return &DBError{
		Code: code,
		Msg:  msg,
	}
}
