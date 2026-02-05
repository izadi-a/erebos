package http

import (
	"encoding/json"
	"net/http"

	app "erebos/internal/application/user"
)

type UserHandler struct {
	create *app.CreateUser
}

func NewUserHandler(create *app.CreateUser) *UserHandler {
	return &UserHandler{create: create}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	u, err := h.create.Execute(r.Context(), req.Email, req.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(u)
}
