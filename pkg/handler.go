package pkg

import (
	"fmt"
	"net/http"
	"strings"
)

func HandleException(w http.ResponseWriter, err interface{}) {

	str := fmt.Sprint(err)
	arr := strings.Split(str, "`")
	code := arr[0]
	msg := arr[1]

	switch code {
	case "400":
		Response(w, 400, msg)
	case "404":
		Response(w, 404, msg)
	default:
		Response(w, 500, msg)
	}

}
