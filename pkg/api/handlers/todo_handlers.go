package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	req "github.com/leoguilen/simple-go-api/pkg/api/models/requests"
	res "github.com/leoguilen/simple-go-api/pkg/api/models/responses"
	"github.com/leoguilen/simple-go-api/pkg/core/models"
	"github.com/leoguilen/simple-go-api/pkg/core/services"
	service_impl "github.com/leoguilen/simple-go-api/pkg/core/services/impl"
)

var (
	todoService services.ITodoService = service_impl.NewTodoService()
)

// swagger:route POST /api/v1/todos Todos createTodo
//
// Create new todo item.
//
// This will create todo item.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Deprecated: false
//
//     Parameters:
//       + name: requestBody
//         in: body
//         description: Request body format to create todo item.
//         required: true
//         type: CreateTodoRequest
//
//     Responses:
//       201: _ Todo created successfully.
//       400: ErrorResponse Validation error.
//       404: ErrorResponse Not found.
//       500: ErrorResponse Unexpected error.
func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorResponse := res.NewErrorResponseFromError(err)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}
	defer r.Body.Close()

	var req req.CreateTodoRequest
	if err := json.Unmarshal(body, &req); err != nil {
		errorResponse := res.NewErrorResponseFromError(err)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		errorResponse := res.NewErrorResponseFromValidationErrors(err.(validator.ValidationErrors))

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	if err := todoService.CreateNew(r.Context(), req); err != nil {
		errorResponse := res.NewErrorResponseFromError(err)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// swagger:route PUT /api/v1/todos/{id} Todos updateTodo
//
// Update todo item.
//
// This will update todo item by id. If exists.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Deprecated: false
//
//     Parameters:
//       + name: id
//         in: path
//         description: The Todo unique identifier.
//         required: true
//         type: string
//         format: uuid
//		   example: f4926cc8-e100-4f5e-9b73-e443f70daa98
//       + name: requestBody
//         in: body
//         description: Request body format to update todo item.
//         required: true
//         type: UpdateTodoRequest
//
//     Responses:
//       200: TodoResponse Todo updated successfully.
//       400: ErrorResponse Validation error.
//       404: ErrorResponse Not found.
//       500: ErrorResponse Unexpected error.
func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId, err := uuid.Parse(vars["todoId"])
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res.NewErrorResponseFromBadRequest(err))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res.NewErrorResponseFromError(err))
		return
	}
	defer r.Body.Close()

	var req req.UpdateTodoRequest
	if err := json.Unmarshal(body, &req); err != nil {
		errorResponse := res.NewErrorResponseFromError(err)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		errorResponse := res.NewErrorResponseFromValidationErrors(err.(validator.ValidationErrors))

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	todo, err := todoService.Update(r.Context(), todoId.String(), req)
	if err != nil {
		switch {
		case err == models.ErrTodoNotFound:
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(res.NewErrorResponseFromNotFoundResource())
			return
		default:
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(res.NewErrorResponseFromError(err))
			return
		}
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
}

// swagger:route PATCH /api/v1/todos/{id}/new-status Todos setTodoStatus
//
// Set new status to todo item.
//
// This will change todo status by id. If exists.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Deprecated: false
//
//     Parameters:
//       + name: id
//         in: path
//         description: The Todo unique identifier.
//         required: true
//         type: string
//         format: uuid
//		   example: f4926cc8-e100-4f5e-9b73-e443f70daa98
//       + name: requestBody
//         in: body
//         description: Request body format to set new todo status.
//         required: true
//         type: SetTodoStatusRequest
//
//     Responses:
//       200: _ Todo status updated successfully.
//       400: ErrorResponse Validation error.
//       404: ErrorResponse Not found.
//       422: ErrorResponse Business requirement error.
//       500: ErrorResponse Unexpected error.
func SetTodoStatusHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId, err := uuid.Parse(vars["todoId"])
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res.NewErrorResponseFromBadRequest(err))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res.NewErrorResponseFromError(err))
		return
	}
	defer r.Body.Close()

	var req req.SetTodoStatusRequest
	if err := json.Unmarshal(body, &req); err != nil {
		errorResponse := res.NewErrorResponseFromError(err)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		errorResponse := res.NewErrorResponseFromValidationErrors(err.(validator.ValidationErrors))

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	todo, err := todoService.SetStatus(r.Context(), todoId.String(), req)
	if err != nil {
		switch {
		case err == models.ErrTodoNotFound:
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(res.NewErrorResponseFromNotFoundResource())
			return
		case err == models.ErrTodoClosed:
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(res.NewErrorResponse(err.Error(), http.StatusUnprocessableEntity))
			return
		default:
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(res.NewErrorResponseFromError(err))
			return
		}
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"todoId":    todo.Id,
		"newStatus": todo.Status,
	})
}

// swagger:route GET /api/v1/todos/{id} Todos getTodoById
//
// Get single todo by id.
//
// This will get single todo by id. If exists.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Deprecated: false
//
//     Parameters:
//       + name: id
//         in: path
//         description: The Todo unique identifier.
//         required: true
//         type: string
//         format: uuid
//		   example: f4926cc8-e100-4f5e-9b73-e443f70daa98
//
//     Responses:
//       200: TodoResponse Todo item response.
//       404: ErrorResponse Not found todo item.
//       500: ErrorResponse Unexpected error.
func GetTodoByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId, err := uuid.Parse(vars["todoId"])
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res.NewErrorResponseFromBadRequest(err))
		return
	}

	todo, err := todoService.GetDetails(r.Context(), todoId.String())
	if err != nil {
		switch {
		case err == models.ErrTodoNotFound:
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(res.NewErrorResponseFromNotFoundResource())
			return
		default:
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(res.NewErrorResponseFromError(err))
			return
		}
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
}

// swagger:route GET /api/v1/todos Todos getAllTodos
//
// Get list of todos.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Deprecated: false
//
//     Responses:
//       200: []TodoResponse List of todo items response.
//       500: ErrorResponse Unexpected error.
func GetAllTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := todoService.ListAll(r.Context())
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res.NewErrorResponseFromError(err))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

// swagger:route DELETE /api/v1/todos/{id} Todos deleteTodo
//
// Delete todo item by id.
//
// This will delete todo item by id. If exists.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Deprecated: false
//
//     Parameters:
//       + name: id
//         in: path
//         description: The Todo unique identifier.
//         required: true
//         type: string
//         format: uuid
//		   example: f4926cc8-e100-4f5e-9b73-e443f70daa98
//
//     Responses:
//       204: _ Todo item deleted successfully.
//       404: ErrorResponse Not found todo item.
//       500: ErrorResponse Unexpected error.
func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId, err := uuid.Parse(vars["todoId"])
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res.NewErrorResponseFromBadRequest(err))
		return
	}

	if err := todoService.Delete(r.Context(), todoId.String()); err != nil {
		switch {
		case err == models.ErrTodoNotFound:
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(res.NewErrorResponseFromNotFoundResource())
			return
		default:
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(res.NewErrorResponseFromError(err))
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
