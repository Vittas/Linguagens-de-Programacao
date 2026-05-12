package function

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// BootUp é o ponto de entrada da Cloud Function
func BootUp(w http.ResponseWriter, r *http.Request) {

	// Configuração básica de CORS
	setCORS(w);

	// Se for preflight (OPTIONS), só retorna 200
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Permite apenas o metodo POST
	if r.Method != http.MethodPost && r.Method != http.MethodDelete && r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		response := map[string]string{"error": "Método não permitido"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Captura o valor do header X-API-ACCESS-TOKEN
	token := r.Header.Get("X-GUARDIAN-KEY")

	// Valida o token: 12DE-76XD-98HJ-00PO
	if token != "12DE-76XD-98HJ-00PO" {
		w.WriteHeader(http.StatusUnauthorized)
		response := map[string]string{"error": "Token inválido, acesso não autorizado"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// ------ AQUI COMEÇA A LÓGICA DA FUNCTION ------
	// ====================== V1 ====================

	//create
	if r.URL.Path == "/v1/user/create" {
		// Get Hardware Name
		decoder := json.NewDecoder(r.Body)

		// Cria um mapa genérico para armazenar o JSON
		var jsonData map[string]interface{}

		// Decodifica o JSON diretamente para o mapa genérico
		if err := decoder.Decode(&jsonData); err != nil {
			returnMessage(w, "Error to decode "+err.Error(), http.StatusBadRequest)
			return
		}
		// Decodifica o JSON recebido
		devID, ok := jsonData["devID"].(string)
		if !ok {
			returnMessage(w, "abstence of parameters", http.StatusBadRequest)
			return
		}

		// Decodifica o JSON recebido
		idUser, ok := jsonData["idUser"].(string)
		if !ok {
			idUser, ok = jsonData["userID"].(string)
			if !ok {
				returnMessage(w, "abstence of parameters", http.StatusBadRequest)
				return
			}
		}

		// Decodifica o JSON recebido
		name, ok := jsonData["name"].(string)
		if !ok {
			returnMessage(w, "abstence of parameters", http.StatusBadRequest)
			return
		}

		photo, ok_foto := jsonData["photo"].(string)
		card, ok_card := jsonData["card"].(string)
		if !ok_foto && !ok_card {
			returnMessage(w, "abstence of parameters photo or card", http.StatusBadRequest)
			return
		}

		// Decodifica o JSON recebido
		idGroup, ok := jsonData["idGroup"].(string)
		if !ok {
			idGroup = "1" // Padrão do control id é o id = 1
		}

		// Get Return
		returned, status := handlerCreateUserHardware(devID, idUser, name, photo, card, idGroup)

		w.WriteHeader(status)
		json.NewEncoder(w).Encode(returned)
		return
	}

	if r.URL.Path == "/v1/user/delete" && r.Method == http.MethodDelete {
		// Get Hardware Name
		decoder := json.NewDecoder(r.Body)

		// Cria um mapa genérico para armazenar o JSON
		var jsonData map[string]interface{}

		// Decodifica o JSON diretamente para o mapa genérico
		if err := decoder.Decode(&jsonData); err != nil {
			returnMessage(w, "Error to decode "+err.Error(), http.StatusBadRequest)
			return
		}
		// Decodifica o JSON recebido
		devID, ok := jsonData["devID"].(string)
		if !ok {
			returnMessage(w, "abstence of parameters", http.StatusBadRequest)
			return
		}

		// Decodifica o JSON recebido
		idUser, ok := jsonData["idUser"].(string)
		if !ok {
			idUser, ok = jsonData["userID"].(string)
			if !ok {
				returnMessage(w, "abstence of parameters", http.StatusBadRequest)
				return
			}
		}

		// Get Return
		returned, status := handlerDeleteHardware(devID, idUser)

		w.WriteHeader(status)
		json.NewEncoder(w).Encode(returned)
		return
	}

	if r.URL.Path == "/v1/user/list" && r.Method == http.MethodGet {
		// Extrai o parâmetro devID da query string
		devID := r.URL.Query().Get("devID")

		if devID == "" {
			returnMessage(w, "abstence of parameters", http.StatusBadRequest)
			return
		}

		// Get Return
		returned, status := handlerGetAllUsersHardware(devID)

		w.WriteHeader(status)
		json.NewEncoder(w).Encode(returned)
		return
	}

	if r.URL.Path == "/v1/user/retrieve" && r.Method == http.MethodGet {
		// Extrai o parâmetro devID da query string
		devID := r.URL.Query().Get("devID")
		if devID == "" {
			returnMessage(w, "abstence of parameters devID", http.StatusBadRequest)
			return
		}

		idUser := r.URL.Query().Get("idUser")
		if idUser == "" {
			returnMessage(w, "abstence of parameters idUser", http.StatusBadRequest)
			return
		}

		// Get Return
		returned, status := handlerGetUserHardware(devID, idUser)

		w.WriteHeader(status)
		json.NewEncoder(w).Encode(returned)
		return
	}

	// ================ REQUESTS ANTIGOS ===============

	//getSnapshot
	if r.URL.Path == "/getUsers" {
		// Get Hardware Name
		decoder := json.NewDecoder(r.Body)

		// Cria um mapa genérico para armazenar o JSON
		var jsonData map[string]interface{}

		// Decodifica o JSON diretamente para o mapa genérico
		if err := decoder.Decode(&jsonData); err != nil {
			returnMessage(w, "Error to decode "+err.Error(), http.StatusBadRequest)
			return
		}
		// Decodifica o JSON recebido
		devID, ok := jsonData["devID"].(string)
		if !ok {
			returnMessage(w, "abstence of parameters", http.StatusBadRequest)
			return
		}

		// Get Return
		returned, status := handlerGetAllUsers(devID)

		w.WriteHeader(status)
		json.NewEncoder(w).Encode(returned)
		return
	}

	//insertUser
	if r.URL.Path == "/insertUser" {
		// Get Hardware Name
		decoder := json.NewDecoder(r.Body)

		// Cria um mapa genérico para armazenar o JSON
		var jsonData map[string]interface{}

		// Decodifica o JSON diretamente para o mapa genérico
		if err := decoder.Decode(&jsonData); err != nil {
			returnMessage(w, "Error to decode "+err.Error(), http.StatusBadRequest)
			return
		}
		// Decodifica o JSON recebido
		devID, ok := jsonData["devID"].(string)
		if !ok {
			returnMessage(w, "abstence of parameters", http.StatusBadRequest)
			return
		}

		// Decodifica o JSON recebido
		user, ok := jsonData["user"].(map[string]interface{})
		if !ok {
			returnMessage(w, "abstence of parameters", http.StatusBadRequest)
			return
		}

		// Get Return
		returned, status := handlerInsertUser(devID, user)

		w.WriteHeader(status)
		json.NewEncoder(w).Encode(returned)
		return
	}

	//insertFaceUser
	if r.URL.Path == "/insertFaceUser" {
		// Get Hardware Name
		decoder := json.NewDecoder(r.Body)

		// Cria um mapa genérico para armazenar o JSON
		var jsonData map[string]interface{}

		// Decodifica o JSON diretamente para o mapa genérico
		if err := decoder.Decode(&jsonData); err != nil {
			returnMessage(w, "Error to decode "+err.Error(), http.StatusBadRequest)
			return
		}
		// Decodifica o JSON recebido
		devID, ok := jsonData["devID"].(string)
		if !ok {
			returnMessage(w, "abstence of parameters", http.StatusBadRequest)
			return
		}

		// Decodifica o JSON recebido
		faces, ok := jsonData["faces"].(map[string]interface{})
		if !ok {
			returnMessage(w, "abstence of parameters", http.StatusBadRequest)
			return
		}

		// Get Return
		returned, status := handlerInsertFaceUser(devID, faces)

		w.WriteHeader(status)
		json.NewEncoder(w).Encode(returned)
		return
	}

	if r.URL.Path == "/insertFaceUserWithUrl" {
		// Get Hardware Name
		decoder := json.NewDecoder(r.Body)

		// Cria um mapa genérico para armazenar o JSON
		var jsonData map[string]interface{}

		// Decodifica o JSON diretamente para o mapa genérico
		if err := decoder.Decode(&jsonData); err != nil {
			returnMessage(w, "Error to decode "+err.Error(), http.StatusBadRequest)
			return
		}
		// Decodifica o JSON recebido
		devID, ok := jsonData["devID"].(string)
		if !ok {
			returnMessage(w, "abstence of parameters", http.StatusBadRequest)
			return
		}

		// Decodifica o JSON recebido
		faces, ok := jsonData["faces"].(map[string]interface{})
		if !ok {
			returnMessage(w, "abstence of parameters", http.StatusBadRequest)
			return
		}

		// Get Return
		returned, status := handlerInsertFaceUserWithURL(devID, faces)

		w.WriteHeader(status)
		json.NewEncoder(w).Encode(returned)
		return
	}

	//handlerInsertCardUser
	if r.URL.Path == "/insertCardUser" {
		// Get Hardware Name
		decoder := json.NewDecoder(r.Body)

		// Cria um mapa genérico para armazenar o JSON
		var jsonData map[string]interface{}

		// Decodifica o JSON diretamente para o mapa genérico
		if err := decoder.Decode(&jsonData); err != nil {
			returnMessage(w, "Error to decode "+err.Error(), http.StatusBadRequest)
			return
		}
		// Decodifica o JSON recebido
		devID, ok := jsonData["devID"].(string)
		if !ok {
			returnMessage(w, "abstence of parameters", http.StatusBadRequest)
			return
		}

		// Decodifica o JSON recebido
		cards, ok := jsonData["cards"].(map[string]interface{})
		if !ok {
			returnMessage(w, "abstence of parameters", http.StatusBadRequest)
			return
		}

		// Get Return
		returned, status := handlerInsertCardUser(devID, cards)

		w.WriteHeader(status)
		json.NewEncoder(w).Encode(returned)
		return
	}

	//deleteUser
	if r.URL.Path == "/deleteUser" {
		// Get Hardware Name
		decoder := json.NewDecoder(r.Body)

		// Cria um mapa genérico para armazenar o JSON
		var jsonData map[string]interface{}

		// Decodifica o JSON diretamente para o mapa genérico
		if err := decoder.Decode(&jsonData); err != nil {
			returnMessage(w, "Error to decode "+err.Error(), http.StatusBadRequest)
			return
		}
		// Decodifica o JSON recebido
		devID, ok := jsonData["devID"].(string)
		if !ok {
			returnMessage(w, "abstence of parameters", http.StatusBadRequest)
			return
		}

		// Decodifica o JSON recebido
		userID, ok := jsonData["userID"].(string)
		if !ok {
			returnMessage(w, "abstence of parameters", http.StatusBadRequest)
			return
		}

		// Get Return
		returned, status := handlerDeleteUser(devID, userID)

		w.WriteHeader(status)
		json.NewEncoder(w).Encode(returned)
		return
	}

	//getUser
	if r.URL.Path == "/getUser" {
		// Get Hardware Name
		decoder := json.NewDecoder(r.Body)

		// Cria um mapa genérico para armazenar o JSON
		var jsonData map[string]interface{}

		// Decodifica o JSON diretamente para o mapa genérico
		if err := decoder.Decode(&jsonData); err != nil {
			returnMessage(w, "Error to decode "+err.Error(), http.StatusBadRequest)
			return
		}
		// Decodifica o JSON recebido
		devID, ok := jsonData["devID"].(string)
		if !ok {
			returnMessage(w, "abstence of parameters", http.StatusBadRequest)
			return
		}

		// Decodifica o JSON recebido
		userID, ok := jsonData["userID"].(string)
		if !ok {
			returnMessage(w, "abstence of parameters", http.StatusBadRequest)
			return
		}

		// Get Return
		returned, status := handlerGetUser(devID, userID)

		w.WriteHeader(status)
		json.NewEncoder(w).Encode(returned)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "method not found",
	})
}

func returnMessage(w http.ResponseWriter, message string, status int) {
	// Resposta de sucesso
	w.WriteHeader(status)

	response := struct {
		Message string `json:"message"`
	}{
		Message: message,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Erro ao codificar as informações de conexão para JSON", http.StatusInternalServerError)
		return
	}
}

func setCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-GUARDIAN-KEY")
	w.Header().Set("Content-Type", "application/json")
}

