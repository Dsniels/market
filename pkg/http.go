package pkg

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"reflect"
	"strconv"
)

type Body[T any] struct {
	Data T `json:"data"`
}

type Params struct {
	Categoria string
	Nombre    string
}

func Response[T any](w http.ResponseWriter, code int, data T) {

	body := &Body[T]{
		Data: data,
	}
	b, _ := json.MarshalIndent(body, "", " ")
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
	case uint:
		i, _ := strconv.Atoi(id)
		return any(uint(i)).(T)
	default:
		i, _ := strconv.Atoi(id)
		return any(i).(T)

	}
}

func GetQuery(r *http.Request, q any) error {
	val := reflect.ValueOf(q)

	if val.Kind() != reflect.Ptr {
		slog.Error("ups, you must pass a pointer to a struct")
		return fmt.Errorf("must pass a pointer to a struct")
	}
	el := val.Elem()
	t := val.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		val := r.URL.Query().Get(field.Name)
		if el.Field(i).CanSet() {
			el.Field(i).SetString(val)
		}
	}
	return nil
}
