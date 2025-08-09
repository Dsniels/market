package pkg

import (
	"fmt"
	"net/http"
	"strings"
)

func PanicException(status int, message string) {
	err := fmt.Errorf("%v`%s", status, message)
	panic(err)
}

func BadRequest(args ...string) {
	msg := "bad request"
	if len(args) > 0 {
		msg = strings.Join(args, "~")
	}
	PanicException(http.StatusBadRequest, msg)
}

func NotFound(args ...string) {
	msg := "not found"
	if len(args) > 0 {
		msg = strings.Join(args, "~")
	}
	PanicException(http.StatusNotFound, msg)
}
