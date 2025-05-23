package utils

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetIDParam(r *http.Request) (int, error) {
	params := mux.Vars(r)
	return strconv.Atoi(params["id"])
}
