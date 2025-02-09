package httperror

import (
	"hot-coffee/pkg/json"
	"hot-coffee/pkg/logger"
	"net/http"
)

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	log := logger.GetLogger()

	m := map[string]interface{}{
		"code":    http.StatusMethodNotAllowed,
		"message": http.StatusText(http.StatusMethodNotAllowed),
	}

	jsonData, err := json.MarshalJson(m)
	if err != nil {
		log.Error(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(m["code"].(int))
	_, err = w.Write(jsonData)

	if err != nil {
		log.Error(err.Error())
	}
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	log := logger.GetLogger()

	m := map[string]interface{}{
		"code":    http.StatusNotFound,
		"message": http.StatusText(http.StatusNotFound),
	}

	jsonData, err := json.MarshalJson(m)
	if err != nil {
		log.Error(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(m["code"].(int))
	_, err = w.Write(jsonData)

	if err != nil {
		log.Error(err.Error())
	}
}
