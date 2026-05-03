package api

import (
	"encoding/json"
	"net/http"
	app "tic-tac-toe/internal/game"

	"github.com/google/uuid"
)

type Handler struct {
	useCase *app.UseCase
}

func NewHandler(useCase *app.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func RegisterRoutes(handler *Handler) {
	http.HandleFunc("/game", handler.CreateGameHandler)
	http.HandleFunc("/game/{id}", handler.UpdateGameHandler)
}

func (h *Handler) CreateGameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "this method cannot be used", http.StatusMethodNotAllowed)
		return
	}
	game, err := h.useCase.CreateGame()
	if err != nil {
		http.Error(w, "game create failed", http.StatusInternalServerError)
		return
	}
	dto := ToDTO(game)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto)
}

func (h *Handler) UpdateGameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "this method cannot be used", http.StatusMethodNotAllowed)
		return
	}
	uuidStr := r.PathValue("id")
	gameID, err := uuid.Parse(uuidStr)
	if err != nil {
		http.Error(w, "invalid game id", http.StatusBadRequest)
		return
	}

	var dto GameDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	game, err := h.useCase.MakeMove(gameID, ToDomain(dto).Board)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ToDTO(game))
}
