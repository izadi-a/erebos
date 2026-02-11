package http

import (
	"encoding/json"
	"net/http"

	appUserCommand "erebos/internal/application/command/user"
	appUserQuery "erebos/internal/application/query/user"
)

type UserHandler struct {
	createUC   *appUserCommand.CreateUserUseCase
	findByIDUC *appUserQuery.FindUserByIDUseCase
	findAllUC  *appUserQuery.FindAllUsersUseCase
	deleteUC   *appUserCommand.DeleteUserUseCase
}

func NewUserHandler(
	createUC *appUserCommand.CreateUserUseCase,
	findByIDUC *appUserQuery.FindUserByIDUseCase,
	findAllUC *appUserQuery.FindAllUsersUseCase,
	deleteUC *appUserCommand.DeleteUserUseCase,
) *UserHandler {
	return &UserHandler{
		createUC:   createUC,
		findByIDUC: findByIDUC,
		findAllUC:  findAllUC,
		deleteUC:   deleteUC,
	}
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UsersResponse struct {
	Users []UserResponse `json:"users"`
}

// @Summary Create a new user
// @Description Create a user with name, email, password
// @Tags user
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User info"
// @Success 201 {object} UserResponse "Created user"
// @Failure 400 {object} ErrorResponse "Invalid request"
// @Router /user [post]
func (handler *UserHandler) Create(w http.ResponseWriter, request *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	u, err := handler.createUC.Execute(request.Context(), &appUserCommand.CreateUserCommandDTO{Name: req.Name, Email: req.Email, Password: req.Password})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

// @Summary Create a new user
// @Description Fetch all users
// @Tags user
// @Accept json
// @Produce json
// @Success 201 {object} UsersResponse "Got users"
// @Failure 400 {object} ErrorResponse "Invalid request"
// @Router /users [get]
func (handler *UserHandler) FindAll(w http.ResponseWriter, request *http.Request) {
	users, err := handler.findAllUC.Execute(request.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, users)
}

// @Summary Create a new user
// @Description Fetch a user
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 201 {object} UserResponse "Got user"
// @Failure 400 {object} ErrorResponse "Invalid request"
// @Router /user/{id} [get]
func (handler *UserHandler) FindByID(w http.ResponseWriter, request *http.Request, id string) {
	u, err := handler.findByIDUC.Execute(request.Context(), appUserQuery.FindUserByIdQueryDTO{ID: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, http.StatusOK, u)
}

// @Summary Delete a new user
// @Description Delete a user
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 201 {object} nil "Deleted user"
// @Failure 400 {object} ErrorResponse "Invalid request"
// @Router /user/{id} [delete]
func (handler *UserHandler) Delete(w http.ResponseWriter, request *http.Request, id string) {
	if err := handler.deleteUC.Execute(request.Context(), &appUserCommand.DeleteUserCommandDTO{ID: id}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
