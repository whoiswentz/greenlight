package main

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type request struct{}

func (req *request) GetIntParam(r *http.Request, param string) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName(param), 10, 64)
	if err != nil {
		return 0, errors.New("invalid id parameters")
	}

	return id, nil
}
