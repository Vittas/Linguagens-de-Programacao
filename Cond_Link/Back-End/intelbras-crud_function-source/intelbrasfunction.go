package function

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func _parseUsers(data string) ([]User, error) {
	var users []User
	var currentUser User
	records := make(map[string]string)

	lines := strings.Split(data, "\r\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "found=") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) < 2 {
			continue
		}

		key := parts[0]
		value := parts[1]

		records[key] = value

		if strings.Contains(key, "RecNo") {
			if currentUser.RecNo != 0 {
				users = append(users, currentUser)
			}
			currentUser = User{}
		}

		switch {
		case strings.Contains(key, "CardName"):
			currentUser.CardName = value
		case strings.Contains(key, "CardNo"):
			currentUser.CardNo = value
		case strings.Contains(key, "CardStatus"):
			currentUser.CardStatus = value
		case strings.Contains(key, "CardType"):
			currentUser.CardType = value
		case strings.Contains(key, "CitizenIDNo"):
			currentUser.CitizenIDNo = value
		case strings.Contains(key, "Doors"):
			currentUser.Doors = append(currentUser.Doors, value)
		case strings.Contains(key, "DynamicCheckCode"):
			currentUser.DynamicCheckCode = value
		case strings.Contains(key, "FirstEnter"):
			currentUser.FirstEnter = value == "true"
		case strings.Contains(key, "Handicap"):
			currentUser.Handicap = value == "true"
		case strings.Contains(key, "IsValid"):
			currentUser.IsValid = value == "true"
		case strings.Contains(key, "Password"):
			currentUser.Password = value
		case strings.Contains(key, "RecNo"):
			currentUser.RecNo, _ = strconv.Atoi(value)
		case strings.Contains(key, "RepeatEnterRouteTimeout"):
			currentUser.RepeatEnterRouteTimeout, _ = strconv.ParseUint(value, 10, 64)
		case strings.Contains(key, "TimeSections"):
			section, _ := strconv.Atoi(value)
			currentUser.TimeSections = append(currentUser.TimeSections, section)
		case strings.Contains(key, "UseTime"):
			currentUser.UseTime, _ = strconv.Atoi(value)
		case strings.Contains(key, "UserID"):
			currentUser.UserID, _ = strconv.Atoi(value)
		case strings.Contains(key, "UserType"):
			currentUser.UserType, _ = strconv.Atoi(value)
		case strings.Contains(key, "VTOPosition"):
			currentUser.VTOPosition = value
		case strings.Contains(key, "ValidDateEnd"):
			currentUser.ValidDateEnd = value
		case strings.Contains(key, "ValidDateStart"):
			currentUser.ValidDateStart = value
		}
	}

	if currentUser.RecNo != 0 {
		users = append(users, currentUser)
	}

	return users, nil
}

// This convertes every message of intelbras with separater \r\n into map[string]interface{}
func parseKeyValue(data string) map[string]interface{} {
	result := make(map[string]interface{})

	lines := strings.Split(data, "\r\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			result[key] = value
		}
	}

	return result
}

func _getTimes(years int) (string, string) {
	// Obtém a data atual
	currentTime := time.Now()

	// Formata a data no formato desejado
	formattedCurrentTime := currentTime.Format("2006-01-02 15:04:05")

	// Adiciona 10 anos à data atual
	tenYearsLater := currentTime.AddDate(years, 0, 0)
	formattedYearsLater := tenYearsLater.Format("2006-01-02 15:04:05")
	return formattedCurrentTime, formattedYearsLater
}

func _convertURLToBASE64(url string) (string, error) {

	// Fazer o download da imagem
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erro ao fazer o download da imagem:", err)
		return "", err
	}
	defer resp.Body.Close()

	// Ler o corpo da resposta
	imageData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler o corpo da resposta:", err)
		return "", err
	}

	// Codificar a imagem para base64
	base64Image := base64.StdEncoding.EncodeToString(imageData)
	return base64Image, nil
}

// Returns all users with photo and card
func getAllUsers(host string, middlewareID string, username string, password string) (string, error) {
	// URL do servidor WebSocket
	u := url.URL{Scheme: "ws", Host: "35.202.130.68:9123", Path: "/wsfront"}

	// Criação do cliente WebSocket
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Erro ao conectar ao servidor:", err)
	}
	defer c.Close()

	// Função para enviar mensagem e aguardar resposta recursivamente

	// Objeto JSON a ser enviado
	// Gerar um UUID aleatório
	uuid := uuid.New().String()
	jsonObject := map[string]interface{}{
		"id": uuid,
	}

	// Envio do JSON
	err = c.WriteJSON(jsonObject)
	if err != nil {
		log.Println("Erro ao enviar mensagem:", err)
		return "", err
	}

	// Esperar pelo retorno
	_, message, err := c.ReadMessage()
	if err != nil {
		log.Println("Erro ao ler mensagem:", err)
		return "", err
	}

	fmt.Println("message", message)

	// Mandar o login
	// Objeto JSON a ser enviado
	jsonObject = map[string]interface{}{
		"url":      "http://" + host + "/cgi-bin/recordFinder.cgi?action=doSeekFind&name=AccessControlCard&count=10000",
		"Method":   "POST",
		"auth":     "Digest",
		"username": username,
		"password": password,
		"idTask":   "task_orangepi01_100",
		"id":       middlewareID,
	}
	// Envio do JSON
	err = c.WriteJSON(jsonObject)
	if err != nil {
		log.Println("Erro ao enviar mensagem:", err)
		return "", err
	}

	// Esperar pelo retorno
	_, message, err = c.ReadMessage()
	if err != nil {
		log.Println("Erro ao ler mensagem:", err)
		return "", err
	}

	// Decodificar a mensagem como JSON
	var messageImage map[string]interface{}
	err = json.Unmarshal(message, &messageImage)
	if err != nil {
		fmt.Printf("Erro ao decodificar JSON de: %v\n", err)
		return "", err
	}

	// Payload
	messageTwo, ok := messageImage["message"].(string)
	if !ok {
		return "", fmt.Errorf("error to convert message to string %v", messageImage)
	}
	return messageTwo, nil
}

// Register user
func registerUser(hostUrl string, host string, middlewareID string, username string, password string, user map[string]interface{}, devID string) (string, error) {

	// Mandar o login
	// Objeto JSON a ser enviado
	jsonObject := map[string]interface{}{
		"url":      "http://" + host + "/cgi-bin/AccessUser.cgi?action=insertMulti",
		"Method":   "POST",
		"auth":     "Digest",
		"username": username,
		"password": password,
		"idTask":   "_register_user",
		"devID":    devID,
		"json":     user,
	}
	// Envio do JSON
	retorno, erro := sendMessageToMiddlewareSocketIO(hostUrl, middlewareID, jsonObject)
	if erro != nil {
		return "", erro
	}

	fmt.Println(retorno)

	if response, ok := retorno["response"].(map[string]interface{}); ok {
		if statusCode, ok := response["statusCode"].(float64); ok && statusCode != 200 {
			return "", fmt.Errorf("error to create user: %v", statusCode)
		}
	}
	// Payload
	return "user created with success", nil
}

// Register face
func registerFace(hostUrl, host string, middlewareID string, username string, password string, user map[string]interface{}, devID string) (string, error) {
	// Mandar o login
	// Objeto JSON a ser enviado
	jsonObject := map[string]interface{}{
		"url":      "http://" + host + "/cgi-bin/AccessFace.cgi?action=insertMulti",
		"Method":   "POST",
		"auth":     "Digest",
		"username": username,
		"password": password,
		"idTask":   "_register_face",
		"devID":    devID,
		"json":     user,
	}

	// Envio do JSON
	retorno, erro := sendMessageToMiddlewareSocketIO(hostUrl, middlewareID, jsonObject)
	if erro != nil {
		return "", erro
	}

	fmt.Println(retorno)

	return "image register with success", nil
}

// Register face
func registerCard(host string, middlewareID string, username string, password string, cards map[string]interface{}) (string, error) {
	// URL do servidor WebSocket
	u := url.URL{Scheme: "ws", Host: "35.202.130.68:9123", Path: "/wsfront"}

	// Criação do cliente WebSocket
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Erro ao conectar ao servidor:", err)
	}
	defer c.Close()

	// Função para enviar mensagem e aguardar resposta recursivamente

	// Objeto JSON a ser enviado
	// Gerar um UUID aleatório
	uuid := uuid.New().String()
	jsonObject := map[string]interface{}{
		"id": uuid,
	}

	// Envio do JSON
	err = c.WriteJSON(jsonObject)
	if err != nil {
		log.Println("Erro ao enviar mensagem:", err)
		return "", err
	}

	// Esperar pelo retorno
	_, message, err := c.ReadMessage()
	if err != nil {
		log.Println("Erro ao ler mensagem:", err)
		return "", err
	}
	fmt.Println("message", message)

	// Mandar o login
	// Objeto JSON a ser enviado
	jsonObject = map[string]interface{}{
		"url":      "http://" + host + "/cgi-bin/AccessCard.cgi?action=insertMulti",
		"Method":   "POST",
		"auth":     "Digest",
		"username": username,
		"password": password,
		"idTask":   "task_orangepi01_100",
		"id":       middlewareID,
		"data":     cards,
	}
	// Envio do JSON
	err = c.WriteJSON(jsonObject)
	if err != nil {
		log.Println("Erro ao enviar mensagem:", err)
		return "", err
	}

	// Esperar pelo retorno
	_, message, err = c.ReadMessage()
	if err != nil {
		log.Println("Erro ao ler mensagem:", err)
		return "", err
	}

	// Decodificar a mensagem como JSON
	var messageImage map[string]interface{}
	err = json.Unmarshal(message, &messageImage)
	if err != nil {
		fmt.Printf("Erro ao decodificar JSON de: %v\n", err)
		return "", err
	}

	// Payload
	messageTwo, ok := messageImage["message"].(string)
	if !ok {
		return "", fmt.Errorf("error to convert message to string %v", messageImage)
	}
	return messageTwo, nil
}

// delete user
func deleteUser(host string, middlewareID string, username string, password string, userID string) (string, error) {
	// URL do servidor WebSocket
	u := url.URL{Scheme: "ws", Host: "35.202.130.68:9123", Path: "/wsfront"}

	// Criação do cliente WebSocket
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Erro ao conectar ao servidor:", err)
	}
	defer c.Close()

	// Função para enviar mensagem e aguardar resposta recursivamente

	// Objeto JSON a ser enviado
	// Gerar um UUID aleatório
	uuid := uuid.New().String()
	jsonObject := map[string]interface{}{
		"id": uuid,
	}

	// Envio do JSON
	err = c.WriteJSON(jsonObject)
	if err != nil {
		log.Println("Erro ao enviar mensagem:", err)
		return "", err
	}

	// Esperar pelo retorno
	_, message, err := c.ReadMessage()
	if err != nil {
		log.Println("Erro ao ler mensagem:", err)
		return "", err
	}
	fmt.Println("message", message)

	// Mandar o login
	// Objeto JSON a ser enviado
	jsonObject = map[string]interface{}{
		"url":      "http://" + host + "/cgi-bin/AccessUser.cgi?action=removeMulti&UserIDList[0]=" + userID,
		"Method":   "POST",
		"auth":     "Digest",
		"username": username,
		"password": password,
		"idTask":   "task_orangepi01_100",
		"id":       middlewareID,
	}
	// Envio do JSON
	err = c.WriteJSON(jsonObject)
	if err != nil {
		log.Println("Erro ao enviar mensagem:", err)
		return "", err
	}

	// Esperar pelo retorno
	_, message, err = c.ReadMessage()
	if err != nil {
		log.Println("Erro ao ler mensagem:", err)
		return "", err
	}

	// Decodificar a mensagem como JSON
	var messageImage map[string]interface{}
	err = json.Unmarshal(message, &messageImage)
	if err != nil {
		fmt.Printf("Erro ao decodificar JSON de: %v\n", err)
		return "", err
	}

	// Payload
	messageTwo, ok := messageImage["message"].(string)
	if !ok {
		return "", fmt.Errorf("error to convert message to string %v", messageImage)
	}
	return messageTwo, nil
}

// get user
func getUser(host string, middlewareID string, username string, password string, userID string) (map[string]interface{}, error) {
	// URL do servidor WebSocket
	u := url.URL{Scheme: "ws", Host: "35.202.130.68:9123", Path: "/wsfront"}

	// Criação do cliente WebSocket
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Erro ao conectar ao servidor:", err)
	}
	defer c.Close()

	// Função para enviar mensagem e aguardar resposta recursivamente

	// Objeto JSON a ser enviado
	// Gerar um UUID aleatório
	uuid := uuid.New().String()
	jsonObject := map[string]interface{}{
		"id": uuid,
	}

	// Envio do JSON
	err = c.WriteJSON(jsonObject)
	if err != nil {
		log.Println("Erro ao enviar mensagem:", err)
		return nil, err
	}

	// Esperar pelo retorno
	_, message, err := c.ReadMessage()
	if err != nil {
		log.Println("Erro ao ler mensagem:", err)
		return nil, err
	}
	fmt.Println("message", message)

	// Mandar o login
	// Objeto JSON a ser enviado
	jsonObject = map[string]interface{}{
		"url":      "http://" + host + "/cgi-bin/AccessUser.cgi?action=list&UserIDList[0]=" + userID,
		"Method":   "POST",
		"auth":     "Digest",
		"username": username,
		"password": password,
		"idTask":   "task_orangepi01_100",
		"id":       middlewareID,
	}
	// Envio do JSON
	err = c.WriteJSON(jsonObject)
	if err != nil {
		log.Println("Erro ao enviar mensagem:", err)
		return nil, err
	}

	// Esperar pelo retorno
	_, message, err = c.ReadMessage()
	if err != nil {
		log.Println("Erro ao ler mensagem:", err)
		return nil, err
	}

	// Decodificar a mensagem como JSON
	var messageImage map[string]interface{}
	err = json.Unmarshal(message, &messageImage)
	if err != nil {
		fmt.Printf("Erro ao decodificar JSON de: %v\n", err)
		return nil, err
	}

	// Payload
	messageTwo, ok := messageImage["message"].(string)
	if !ok {
		return nil, fmt.Errorf("error to convert message to string %v", messageImage)
	}

	// ==================== Get info about image ======================
	// Objeto JSON a ser enviado
	jsonObject = map[string]interface{}{
		"url":      "http://" + host + "/cgi-bin/AccessFace.cgi?action=list&UserIDList[0]=" + userID,
		"Method":   "POST",
		"auth":     "Digest",
		"username": username,
		"password": password,
		"idTask":   "task_orangepi01_100",
		"id":       middlewareID,
	}
	// Envio do JSON
	err = c.WriteJSON(jsonObject)
	if err != nil {
		log.Println("Erro ao enviar mensagem:", err)
		return nil, err
	}

	// Esperar pelo retorno
	_, message, err = c.ReadMessage()
	if err != nil {
		log.Println("Erro ao ler mensagem:", err)
		return nil, err
	}

	// Decodificar a mensagem como JSON
	var imageData map[string]interface{}
	err = json.Unmarshal(message, &imageData)
	if err != nil {
		fmt.Printf("Erro ao decodificar JSON de: %v\n", err)
		return nil, err
	}

	// Payload
	messageData, ok := imageData["message"].(string)
	if !ok {
		return nil, fmt.Errorf("error to convert message to string %v", imageData)
	}

	// ==================== Get info about cartao ======================
	// Objeto JSON a ser enviado
	jsonObject = map[string]interface{}{
		"url":      "http://" + host + "/cgi-bin/AccessCard.cgi?action=startFind&Condition.UserID=" + userID,
		"Method":   "POST",
		"auth":     "Digest",
		"username": username,
		"password": password,
		"idTask":   "task_orangepi01_100",
		"id":       middlewareID,
	}
	// Envio do JSON
	err = c.WriteJSON(jsonObject)
	if err != nil {
		log.Println("Erro ao enviar mensagem:", err)
		return nil, err
	}

	// Esperar pelo retorno
	_, message, err = c.ReadMessage()
	if err != nil {
		log.Println("Erro ao ler mensagem:", err)
		return nil, err
	}

	// Decodificar a mensagem como JSON
	var cartaoData map[string]interface{}
	err = json.Unmarshal(message, &cartaoData)
	if err != nil {
		fmt.Printf("Erro ao decodificar JSON de: %v\n", err)
		return nil, err
	}

	// Payload
	messageCartao, ok := cartaoData["message"].(string)
	if !ok {
		return nil, fmt.Errorf("error to convert message to string %v", cartaoData)
	}

	// ==================== Get cartao info ======================
	var cartao map[string]interface{}
	cartaoinfo := _convertStringToMap(messageCartao, ":")
	token, ok := cartaoinfo["Token"].(string)
	if ok {
		total, ok := cartaoinfo["Total"].(string)
		if ok {
			// Objeto JSON a ser enviado
			jsonObject = map[string]interface{}{
				"url":      "http://" + host + "/cgi-bin/AccessCard.cgi?action=doFind&Token=" + token + "&Offset=0&Count=" + total,
				"Method":   "POST",
				"auth":     "Digest",
				"username": username,
				"password": password,
				"idTask":   "task_orangepi01_100",
				"id":       middlewareID,
			}
			// Envio do JSON
			err = c.WriteJSON(jsonObject)
			if err != nil {
				log.Println("Erro ao enviar mensagem:", err)
				return nil, err
			}

			// Esperar pelo retorno
			_, message, err = c.ReadMessage()
			if err != nil {
				log.Println("Erro ao ler mensagem:", err)
				return nil, err
			}

			// Decodificar a mensagem como JSON
			err = json.Unmarshal(message, &cartaoData)
			if err != nil {
				fmt.Printf("Erro ao decodificar JSON de: %v\n", err)
				return nil, err
			}

			// Payload
			cartaoMessageM, ok := cartaoData["message"].(string)
			if !ok {
				return nil, fmt.Errorf("error to convert message to string %v", cartaoData)
			}
			cartao = _convertStringToMap(cartaoMessageM, ":")
		}
	}

	return map[string]interface{}{
		"userInfo":         _convertStringToMap(messageTwo, "="),
		"imageInfo":        _convertStringToMap(messageData, "="),
		"configCartaoInfo": _convertStringToMap(messageCartao, ":"),
		"cartaoInfo":       cartao,
	}, nil
}

// ======================== V2 ========================
// Register user
func createUserv2(hostUrl string, photo string, card string, idUser string, name string, idMiddleware string, host string, username string, password string, idGroup int64) (interface{}, error) {
	// Mandar o login
	// Objeto JSON a ser enviado
	now, to := _getTimes(10) // get now time e more than 10 years

	var jsonObject map[string]interface{}
	if idGroup == 255{
		jsonObject = map[string]interface{}{
			"url":      "http://" + host + "/cgi-bin/recordUpdater.cgi?action=insert&name=AccessControlCard&CardNo="+ idUser +"&CardStatus=0&CardName=" + name + "&UserID=" + idUser,
			"Method":   "POST",
			"auth":     "Digest",
			"username": username,
			"password": password,
			"idTask":   "_create_user",
		}
	}else{

		jsonObject = map[string]interface{}{
			"url":      "http://" + host + "/cgi-bin/AccessUser.cgi?action=insertMulti",
			"Method":   "POST",
			"auth":     "Digest",
			"username": username,
			"password": password,
			"idTask":   "_create_user",
			"data": map[string]interface{}{
				"UserList": []map[string]interface{}{
					{
						"UserID":       idUser,
						"UserName":     name,
						"UserType":     0,
						"Password":     "",
						"Doors":        []int{0},
						"TimeSections": []int64{idGroup},
						"ValidFrom":    now,
						"ValidTo":      to,
					},
				},
			},
		}
	}

	// Envio do JSON
	retorno, erro := sendMessageToMiddlewareSocketIO(hostUrl, idMiddleware, jsonObject)
	if erro != nil {
		return nil, erro
	}

	// Verifica se deu erro na requisição
	// Inclusive capta o erro de desconexão do middleware
	if err, ok := retorno["error"].(string); ok {
		return nil, fmt.Errorf("error to create user: " + err)
	}

	if response, ok := retorno["response"].(map[string]interface{}); ok {
		if statusCode, ok := response["statusCode"].(float64); ok && statusCode != 200 {
			if message, ok := response["message"].(string); ok{
				return nil, fmt.Errorf("error to create user: %v - " + message, statusCode)
			}
			return nil, fmt.Errorf("error to create user: %v", statusCode)
		}
	}


	var infos = map[string]interface{}{}

	// Envio da Face
	if photo != "" {

		var base64Image string
		if strings.HasPrefix(photo, "http") {
			base64Image, erro = _convertURLToBASE64(photo)
			if erro != nil {
				log.Println("Erro ao converter imagem:", erro)
				return nil, erro
			}

			if len(base64Image) > 1000000 {
				return nil, fmt.Errorf("base64 too big")
			}
		} else {
			base64Image = photo // imageURL can be a base64
		}

		jsonObject = map[string]interface{}{
			"url":      "http://" + host + "/cgi-bin/AccessFace.cgi?action=insertMulti",
			"Method":   "POST",
			"auth":     "Digest",
			"username": username,
			"password": password,
			"idTask":   "_insert_face",
			"data": map[string]interface{}{
				"FaceList": []map[string]interface{}{
					{
						"UserID":    idUser,
						"PhotoData": []string{base64Image},
					},
				},
			},
		}

		retorno, erro := sendMessageToMiddlewareSocketIO(hostUrl, idMiddleware, jsonObject)
		if erro != nil {
			return nil, erro
		}

		// Verifica se deu erro na requisição
		// Inclusive capta o erro de desconexão do middleware
		if err, ok := retorno["error"].(string); ok {
			return nil, fmt.Errorf("error to create user: " + err)
		}

		infos["image"] = retorno

		if response, ok := retorno["response"].(map[string]interface{}); ok {
			if statusCode, ok := response["statusCode"].(float64); ok && statusCode != 200 {
				return nil, fmt.Errorf("error to send image: %v - %v", statusCode, response)
			}

			if message, ok := response["message"].(string); ok && message == "Batch Process Error" {
				return nil, fmt.Errorf("error to send image (problem with image): " + message)
			}

		}

	}

	return infos, nil
}

// Delete user
func deleteUserv2(hostUrl string, idUser string, idMiddleware string, host string, username string, password string) (interface{}, error) {
	// Mandar o login
	jsonObject := map[string]interface{}{
		"url":      "http://" + host + "/cgi-bin/AccessUser.cgi?action=removeMulti&UserIDList[0]=" + idUser,
		"Method":   "POST",
		"auth":     "Digest",
		"username": username,
		"password": password,
		"idTask":   "_delete_user",
	}
	// Envio do JSON
	retorno, erro := sendMessageToMiddlewareSocketIO(hostUrl, idMiddleware, jsonObject)
	if erro != nil {
		return nil, erro
	}

	if err, ok := retorno["error"].(string); ok {
		return nil, fmt.Errorf("error to delete user: " + err)
	}

	if response, ok := retorno["response"].(map[string]interface{}); ok {
		if statusCode, ok := response["statusCode"].(float64); ok && statusCode != 200 {
			return nil, fmt.Errorf("error to delete user: %v", statusCode)
		}
	}

	return "user deleted with success", nil
}

// Get all users
func getUsersv2(hostUrl string, idMiddleware string, host string, username string, password, devID string) (interface{}, error) {
	// Mandar o login
	jsonObject := map[string]interface{}{
		"url":      "http://" + host + "/cgi-bin/recordFinder.cgi?action=doSeekFind&name=AccessControlCard&count=1000",
		"Method":   "POST",
		"auth":     "Digest",
		"username": username,
		"password": password,
		"devID":    devID,
		"timeout":  50,
		"idTask":   "_get_users",
	}
	// Envio do JSON
	retorno, erro := sendMessageToMiddlewareSocketIO(hostUrl, idMiddleware, jsonObject)
	if erro != nil {
		return nil, erro
	}

	if response, ok := retorno["response"].(map[string]interface{}); ok {
		if statusCode, ok := response["statusCode"].(float64); ok && statusCode != 200 {
			return nil, fmt.Errorf("error to create user: %v", statusCode)
		}

		// Retorno igual a 200
		if payload, ok := response["payload"].(string); ok {
			users, erro := _parseUsers(payload)
			if erro != nil {
				return nil, erro
			}
			return users, nil
		}
	}

	return []interface{}{}, nil
}

// Get info about user
func getUserv2(hostUrl string, idMiddleware string, host string, username string, password string, idUser string) (interface{}, error) {
	// Mandar o login
	jsonObject := map[string]interface{}{
		"url":      "http://" + host + "/cgi-bin/AccessUser.cgi?action=list&UserIDList[0]=" + idUser,
		"Method":   "POST",
		"auth":     "Digest",
		"username": username,
		"password": password,
		"idTask":   "task_orangepi01_100",
	}
	// Envio do JSON
	retorno, erro := sendMessageToMiddlewareSocketIO(hostUrl, idMiddleware, jsonObject)
	if erro != nil {
		return nil, erro
	}

	var mapToReturn map[string]interface{}
	if response, ok := retorno["response"].(map[string]interface{}); ok {
		if statusCode, ok := response["statusCode"].(float64); ok && statusCode != 200 {
			return nil, fmt.Errorf("error to create user: %v", statusCode)
		}

		if payload, ok := response["payload"].(string); ok {
			mapToReturn = parseKeyValue(payload)
		}
	}

	// Image
	jsonObject = map[string]interface{}{
		"url":      "http://" + host + "/cgi-bin/AccessFace.cgi?action=list&UserIDList[0]=" + idUser,
		"Method":   "POST",
		"auth":     "Digest",
		"username": username,
		"password": password,
		"idTask":   "task_orangepi01_100",
	}

	retorno, erro = sendMessageToMiddlewareSocketIO(hostUrl, idMiddleware, jsonObject)
	if erro != nil {
		return nil, erro
	}

	if response, ok := retorno["response"].(map[string]interface{}); ok {
		if statusCode, ok := response["statusCode"].(float64); ok && statusCode != 200 {
			return mapToReturn, nil
		}

		if payload, ok := response["payload"].(string); ok {
			mapToReturn = mergeMaps(mapToReturn, parseKeyValue(payload))
		}
	}

	return mapToReturn, nil
}
