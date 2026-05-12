package function

import (
	"encoding/json"
	"net/http"
)

func returnMessage(w http.ResponseWriter, message string, status int) {
	// Resposta de sucesso
	w.WriteHeader(status)

	response := struct {
		Message string `json:"message"`
	}{
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Erro ao codificar as informações de conexão para JSON", http.StatusInternalServerError)
		return
	}
}