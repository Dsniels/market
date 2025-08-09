package pkg

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Body[T any] struct {
	Data T `json:"data"`
}

func Response[T any](w http.ResponseWriter, code int, data T) {

	body := &Body[T]{
		Data: data,
	}
	b, _ := json.Marshal(body)
	w.WriteHeader(code)
	w.Header().Set("content-type", "application/json")
	w.Write(b)
}

func GetIDFromUrl[T string | uint | int64](r *http.Request) T {
	id := r.PathValue("id")

	var t T
	switch any(t).(type) {
	case string:
		return any(id).(T)
	default:
		i, _ := strconv.Atoi(id)
		return any(i).(T)

	}
}
