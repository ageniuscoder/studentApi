package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/ageniouscoder/student-api/internal/types"
	"github.com/ageniouscoder/student-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Creating student api")
		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GenralError(fmt.Errorf("empty body")))
			return
		}
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GenralError(err))
		}

		//request validation

		if e := validator.New().Struct(student); e != nil {
			errMsg := e.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(errMsg))
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]string{"succes": "OK"})
	}
}
