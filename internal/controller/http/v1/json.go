package v1

import (
	"encoding/json"
	"github.com/prok05/gophermart/internal/controller/http/response"
	"io"
	"net/http"
	"strings"
)

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_048_578 // 1mb
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

func readPlain(r *http.Request) (string, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	return strings.TrimSpace(string(body)), nil
}

func writeJSONError(w http.ResponseWriter, status int, message string) error {
	return writeJSON(w, status, &response.Error{Error: message})
}

func (h *V1) jsonResponse(w http.ResponseWriter, status int, data any) error {
	return writeJSON(w, status, &response.JSON{Data: data})
}
