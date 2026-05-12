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
	if r.Method != http.MethodPost {
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

	if r.URL.Path == "/createUser" {
		// Get hik id
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

		url_photo, ok_foto := jsonData["url_photo"].(string)
		if !ok_foto {
			url_photo, ok_foto = jsonData["photo"].(string)
		}
		card, ok_card := jsonData["card"].(string)
		tag, ok_tag := jsonData["tag"].(string)
		qrcode, ok_qrcode := jsonData["qrcode"].(string)
		if !ok_foto && !ok_card && !ok_tag && !ok_qrcode {
			returnMessage(w, "abstence of parameters url_photo or card or qrcode", http.StatusBadRequest)
			return
		}

		// Decodifica o JSON recebido
		idGroup, ok := jsonData["idGroup"].(string)
		if !ok || idGroup == "" || idGroup == "-1" {
			idGroup = "1" // Padrão do control id é o id = 1
		}

		// Get Return
		returned, status := handlerCreateUserHardware(devID, idUser, name, url_photo, card, idGroup, tag, qrcode)

		w.WriteHeader(status)
		json.NewEncoder(w).Encode(returned)
		return
	}

	if r.URL.Path == "/deleteUser" {
		// Get hik id
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
		returned, status := handlerDeleteUserHardware(devID, idUser)

		w.WriteHeader(status)
		json.NewEncoder(w).Encode(returned)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "method not implemented",
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
