package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/ageniouscoder/student-api/internal/storage"
	"github.com/ageniouscoder/student-api/internal/types"
	"github.com/ageniouscoder/student-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New(storage storage.Storage) http.HandlerFunc {
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

		//database part

		lastId, err := storage.CreateStudent(student.Name, student.Email, student.Age)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err)
		}
		slog.Info("user created succesfully", slog.String("userId", fmt.Sprint(lastId)))

		//

		response.WriteJson(w, http.StatusCreated, map[string]int64{"id": lastId})
	}
}

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		slog.Info("getting student by Id", slog.String("id", id))

		intId, err := strconv.ParseInt(id, 10, 64)

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GenralError(err))
			return
		}

		student, err := storage.GetStudentById(intId)

		if err != nil {
			slog.Error("error getting student", slog.String("id", id))
			response.WriteJson(w, http.StatusInternalServerError, response.GenralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, student)
	}
}

func GetList(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("getting all student Info")

		students, err := storage.GetStudents()

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err)
		}

		response.WriteJson(w, http.StatusOK, students)
	}
}

func UpdateByid(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GenralError(err))
			return
		}

		var student types.Student
		err = json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GenralError(fmt.Errorf("body is empty")))
			return
		}
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GenralError(err))
			return
		}
		slog.Info("Start updating student with ", slog.Int64("id", id))
		rowAffected, err := storage.UpdateStudent(student.Name, student.Email, student.Age, id)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GenralError(err))
			return
		}
		slog.Info("User updated Succesfully")
		response.WriteJson(w, http.StatusOK, map[string]int64{"No of rows Updated": rowAffected})

	}
}

func DeleteByid(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GenralError(err))
			return
		}
		slog.Info("deleting student with", slog.Int64("id", id))
		affRow, err := storage.DeleteStudent(id)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GenralError(err))
			return
		}
		slog.Info("student deleted success")
		if affRow == 1 {
			response.WriteJson(w, http.StatusOK, map[string]string{"deleted": "success"})
		} else {
			response.WriteJson(w, http.StatusInternalServerError, map[string]string{"deleted": "failed"})
		}
	}
}
