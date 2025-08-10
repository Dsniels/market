package pkg

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"
)

func HandleException(w http.ResponseWriter, err interface{}) {

	str := fmt.Sprint(err)
	slog.Error(str)
	arr := strings.Split(str, "`")
	code := arr[0]
	msg := arr[1]

	switch code {
	case "400":
		Response(w, 400, msg)
	case "404":
		Response(w, 404, msg)
	case "409":
		Response(w, 409, msg)
	default:
		Response(w, 500, msg)
	}

}
