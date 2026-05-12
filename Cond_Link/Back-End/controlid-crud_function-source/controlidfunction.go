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

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// This functions returns a connection of websocket
func _open_connection_websocket() (*websocket.Conn, error) {
	u := url.URL{Scheme: "ws", Host: "35.202.130.68:9123", Path: "/wsfront"}

	// Criação do cliente WebSocket
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Erro ao conectar ao servidor:", err)
	}

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
	_, _, err = c.ReadMessage()
	if err != nil {
		log.Println("Erro ao ler mensagem:", err)
		return nil, err
	}
	return c, nil
}

// This function sends a message to connection websocket
func _send_message_to_websocket(connection *websocket.Conn, jsonString map[string]interface{}) (map[string]interface{}, error) {

	// Envio do JSON
	err := connection.WriteJSON(jsonString)
	if err != nil {
		return nil, fmt.Errorf("Error: send message to websocket " + err.Error())
	}

	// Esperar pelo retorno
	_, message, err := connection.ReadMessage()
	if err != nil {
		return nil, fmt.Errorf("Error: read message to websocket " + err.Error())
	}

	// Decodificar a mensagem como JSON
	var messageMap map[string]interface{}
	err = json.Unmarshal(message, &messageMap)
	if err != nil {
		return nil, fmt.Errorf("Error: cannot decode message from websocket " + err.Error())
	}

	if payload, ok := messageMap["payload"].(map[string]interface{}); ok {
		return payload, nil
	}

	return messageMap, nil
}

// Function to create user with Image
func createUser(imageUrl string, card string, idUser int64, nome string, middlewareID string, host string, username string, senha string, groupId int64) ([]interface{}, error) {

	connection, err := _open_connection_websocket()
	if err != nil {
		return nil, err
	}
	defer connection.Close()

	// Mandar o login
	// Objeto JSON a ser enviado
	jsonObject := map[string]interface{}{
		"url":    "http://" + host + "/login.fcgi",
		"Method": "POST",
		"idTask": "task_orangepi01_100",
		"id":     middlewareID,
		"json": map[string]string{
			"login":    username,
			"password": senha,
		},
	}

	// Envio do JSON
	receivedMessage, err := _send_message_to_websocket(connection, jsonObject)
	if err != nil {
		log.Println("Erro ao enviar mensagem:", err)
		return nil, fmt.Errorf("error to write from websocket " + err.Error())
	}

	fmt.Println(receivedMessage)
	sessionValue := receivedMessage["session"].(string)

	// Imprimindo o valor de "session"
	fmt.Println("Valor da session:", sessionValue)

	// Criando um usuário
	jsonObject = map[string]interface{}{
		"url":    "http://" + host + "/create_objects.fcgi?session=" + sessionValue,
		"Method": "POST",
		"idTask": "task_orangepi01_101",
		"id":     middlewareID, // Definido como um inteiro
		"json": map[string]interface{}{
			"object": "users",
			"values": []map[string]interface{}{
				{
					"name":         nome,
					"registration": "",
					"password":     "1234", //TODO REMOVE THAT
					"salt":         "",
					"id":           idUser, // Definido como um inteiro
					"begin_time":   0,
					"end_time":     0,
					"expires":      0,
				},
			},
		},
	}

	// Envio do JSON
	_, err = _send_message_to_websocket(connection, jsonObject)
	if err != nil {
		log.Println("Erro ao enviar mensagem:", err)
		return nil, fmt.Errorf("error to write from websocket " + err.Error())
	}

	// Put user on contingency group
	// ================= Salva User on Group ===================
	jsonObject = map[string]interface{}{
		"url":    "http://" + host + "/create_objects.fcgi?session=" + sessionValue,
		"Method": "POST",
		"idTask": "task_orangepi01_101",
		"id":     middlewareID, // Definido como um inteiro
		"json": map[string]interface{}{
			"object": "user_groups",
			"values": []map[string]interface{}{
				{
					"user_id":  idUser,
					"group_id": groupId,
				},
			},
		},
	}

	// Envio do JSON
	_, err = _send_message_to_websocket(connection, jsonObject)
	if err != nil {
		log.Println("Erro ao enviar mensagem:", err)
		return nil, fmt.Errorf("error to write from websocket " + err.Error())
	}

	var retornoInsertUser map[string]interface{}
	// Add image
	if imageUrl != "" {
		var base64Image string
		if strings.HasPrefix(imageUrl, "http") {
			base64Image, err = convertURLToBASE64(imageUrl)
			if err != nil {
				log.Println("Erro ao converter imagem:", err)
				return nil, err
			}

			if len(base64Image) > 1000000 {
				return nil, fmt.Errorf("base64 too big")
			}
			fmt.Println(len(base64Image))
		} else {
			base64Image = imageUrl // imageURL can be a base64
		}

		// Salvando uma imagem
		jsonObject = map[string]interface{}{
			"url":    "http://" + host + "/user_set_image_list.fcgi?session=" + sessionValue,
			"Method": "POST",
			"idTask": "task_orangepi01_102",
			"id":     middlewareID,
			"json": map[string]interface{}{
				"match": false,
				"user_images": []map[string]interface{}{
					{
						"user_id":   idUser,
						"timestamp": 1715782339,
						"image":     base64Image,
					},
				},
			},
		}

		// Envio do JSON
		retornoInsertUser, err = _send_message_to_websocket(connection, jsonObject)
		if err != nil {
			log.Println("Erro ao enviar mensagem:", err)
			return nil, fmt.Errorf("error to write from websocket " + err.Error())
		}

		if returns, ok := retornoInsertUser["results"].([]interface{}); ok && len(returns) > 0 {

			mapp, ok := returns[0].(map[string]interface{})
			if ok {

				erros, ok := mapp["errors"].([]interface{})
				if ok {
					// devo apagar o usuário caso deem algum erro
					return nil, fmt.Errorf("error to create user %v", erros...)
				}
			}

		}
	}

	// Criando um card
	if card != "" {

		cardWithCommmand := _addCommaIfAbsent(card)

		device_id, card_numero, err := splitAndConvert(cardWithCommmand)
		if err != nil {
			return nil, err
		}

		num := calculateCardNumber(device_id, card_numero)

		jsonObject = map[string]interface{}{
			"url":    "http://" + host + "/create_objects.fcgi?session=" + sessionValue,
			"Method": "POST",
			"idTask": "task_orangepi01_101",
			"id":     middlewareID, // Definido como um inteiro
			"json": map[string]interface{}{
				"object": "cards",
				"values": []map[string]interface{}{
					{
						"value":   num,
						"user_id": idUser,
					},
				},
			},
		}

		// Envio do JSON
		_, err = _send_message_to_websocket(connection, jsonObject)
		if err != nil {
			log.Println("Erro ao enviar mensagem:", err)
			return nil, fmt.Errorf("error to write from websocket " + err.Error())
		}
	}

	return []interface{}{"success to create user", retornoInsertUser}, nil
}

func getPayload(body map[string]interface{}) (map[string]interface{}, bool) {

	response, ok := body["response"].(map[string]interface{})
	if !ok {
		return nil, false
	}

	info, ok := response["payload"].(map[string]interface{})
	if !ok {
		return nil, false
	}

	return info, true
}

func _addCommaIfAbsent(s string) string {
	// Check if the string already contains a comma
	s = strings.Replace(s, ".", ",", -1)
	if strings.Contains(s, ",") {
		return s
	}

	// If the string is shorter than 5 characters, return it as is
	if len(s) <= 5 {
		return s
	}

	// Insert a comma 5 characters from the end
	return s[:len(s)-5] + "," + s[len(s)-5:]
}

func splitAndConvert(s string) (int64, int64, error) {
	parts := strings.Split(s, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("string does not contain exactly one comma", s)
	}

	firstPart, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("error converting first part to int64: %v", err)
	}

	secondPart, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("error converting second part to int64: %v", err)
	}

	return firstPart, secondPart, nil
}

func calculateCardNumber(firstPart, secondPart int64) int64 {
	return firstPart*4294967296 + secondPart // 2^32 = 4294967296
}

// Function to create user with Card
func deleteUserControlID(middlewareID string, host string, username string, senha string, user int64) (map[string]interface{}, error) {

	// URL do servidor WebSocket
	connection, err := _open_connection_websocket()
	if err != nil {
		return nil, err
	}
	defer connection.Close()

	// Mandar o login
	// Objeto JSON a ser enviado
	jsonObject := map[string]interface{}{
		"url":    "http://" + host + "/login.fcgi",
		"Method": "POST",
		"idTask": "task_orangepi01_100",
		"id":     middlewareID,
		"json": map[string]string{
			"login":    username,
			"password": senha,
		},
	}

	// Envio do JSON

	message, err := _send_message_to_websocket(connection, jsonObject)
	if err != nil {
		fmt.Printf("Erro ao decodificar JSON de: %v\n", err)
		return nil, err
	}

	// Acessando o valor da chave "session" dentro do mapa "payload"
	sessionValue := message["session"].(string)

	// Imprimindo o valor de "session"
	fmt.Println("Valor da session:", sessionValue)

	// Criando um usuário
	jsonObject = map[string]interface{}{
		"url":    "http://" + host + "/destroy_objects.fcgi?session=" + sessionValue,
		"Method": "POST",
		"idTask": "task_orangepi01_101",
		"id":     middlewareID, // Definido como um inteiro
		"json": map[string]interface{}{
			"object": "users",
			"where": map[string]interface{}{
				"users": map[string]interface{}{"id": user},
			},
		},
	}

	// Envio do JSON
	messageCreate, err := _send_message_to_websocket(connection, jsonObject)
	if err != nil {
		fmt.Printf("Erro ao decodificar JSON de: %v\n", err)
		return nil, err
	}

	return messageCreate, nil
}

func convertURLToBASE64(url string) (string, error) {

	// Fazer o download da imagem
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erro ao fazer o download da imagem:", err)
		return "", nil
	}
	defer resp.Body.Close()

	// Ler o corpo da resposta
	imageData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler o corpo da resposta:", err)
		return "", nil
	}

	// Codificar a imagem para base64
	base64Image := base64.StdEncoding.EncodeToString(imageData)
	return base64Image, nil
}

//  =============================== v2 ===================================

// Function to create user with Image
func createUserv2(hostUrl, imageUrl, card string, idUser int64, nome, middlewareID, host, username, senha string, groupId int64, tag, qrcode, devID string) error {

	// Apagar user
	deleteUserv2(hostUrl, middlewareID, host, username, senha, idUser)

	// Mandar o login
	// Objeto JSON a ser enviado
	jsonObject := map[string]interface{}{
		"url":    "http://" + host + "/login.fcgi",
		"Method": "POST",
		"devID":  devID,
		"idTask": "_login",
		"json": map[string]string{
			"login":    username,
			"password": senha,
		},
	}

	// Envio do JSON
	retorno, err := sendMessageToMiddlewareSocketIO(hostUrl, middlewareID, jsonObject)
	if err != nil {
		return err
	}

	message, ok := getPayload(retorno)
	if !ok {
		return fmt.Errorf("error: payload not send: %v", retorno)
	}

	sessionValue, ok := message["session"].(string)
	if !ok {
		return fmt.Errorf("error: session not send: %v", message)
	}

	// Imprimindo o valor de "session"
	fmt.Println("Valor da session:", sessionValue)

	// Criando um usuário
	jsonObject = map[string]interface{}{
		"url":    "http://" + host + "/create_objects.fcgi?session=" + sessionValue,
		"Method": "POST",
		"idTask": "_create_group",
		"devID":  devID,
		"json": map[string]interface{}{
			"object": "users",
			"values": []map[string]interface{}{
				{
					"name":         nome,
					"registration": "",
					"password":     "1234", //TODO REMOVE THAT
					"salt":         "",
					"id":           idUser, // Definido como um inteiro
					"begin_time":   0,
					"end_time":     0,
					"expires":      0,
				},
			},
		},
	}

	// Envio do JSON
	retorno, err = sendMessageToMiddlewareSocketIO(hostUrl, middlewareID, jsonObject)
	if err != nil {
		return err
	}

	fmt.Println("retorno do create_objects - users: %v", retorno)

	// Put user on contingency group
	// ================= Salva User on Group ===================
	jsonObject = map[string]interface{}{
		"url":    "http://" + host + "/create_objects.fcgi?session=" + sessionValue,
		"Method": "POST",
		"idTask": "_insert_group",
		"devID":  devID,
		"json": map[string]interface{}{
			"object": "user_groups",
			"values": []map[string]interface{}{
				{
					"user_id":  idUser,
					"group_id": groupId,
				},
			},
		},
	}

	// Envio do JSON
	retorno, err = sendMessageToMiddlewareSocketIO(hostUrl, middlewareID, jsonObject)
	if err != nil {
		return err
	}

	fmt.Println("retorno do create_objects - user_groups: %v", retorno)

	// Add image
	if imageUrl != "" {
		var base64Image string
		if strings.HasPrefix(imageUrl, "http") {
			base64Image, err = convertURLToBASE64(imageUrl)
			if err != nil {
				log.Println("Erro ao converter imagem:", err)
				return err
			}

			if len(base64Image) > 1000000 {
				return fmt.Errorf("base64 too big")
			}
			fmt.Println(len(base64Image))
		} else {
			base64Image = imageUrl // imageURL can be a base64
		}

		// Salvando uma imagem
		jsonObject = map[string]interface{}{
			"url":    "http://" + host + "/user_set_image_list.fcgi?session=" + sessionValue,
			"Method": "POST",
			"idTask": "_set_image",
			"devID":  devID,
			"json": map[string]interface{}{
				"match": false,
				"user_images": []map[string]interface{}{
					{
						"user_id":   idUser,
						"timestamp": 1715782339,
						"image":     base64Image,
					},
				},
			},
		}

		// Envio do JSON
		retorno, err = sendMessageToMiddlewareSocketIO(hostUrl, middlewareID, jsonObject)
		if err != nil {
			return err
		}

		fmt.Println("retorno do user_set_image_list: %v", retorno)

		payload, ok := getPayload(retorno)
		if !ok {
			return fmt.Errorf("error: payload not send: %v", retorno)
		}

		if returns, ok := payload["results"].([]interface{}); ok && len(returns) > 0 {

			mapp, ok := returns[0].(map[string]interface{})
			if ok {

				erros, ok := mapp["errors"].([]interface{})
				if ok {
					// devo apagar o usuário caso deem algum erro
					return fmt.Errorf("error to create user %v", erros...)
				}
			}

		}
	}

	// Criando um card
	if card != "" {

		cardWithCommmand := _addCommaIfAbsent(card)

		device_id, card_numero, err := splitAndConvert(cardWithCommmand)
		if err != nil {
			return err
		}

		num := calculateCardNumber(device_id, card_numero)

		jsonObject = map[string]interface{}{
			"url":    "http://" + host + "/create_objects.fcgi?session=" + sessionValue,
			"Method": "POST",
			"idTask": "_create_card",
			"devID":  devID,
			"json": map[string]interface{}{
				"object": "cards",
				"values": []map[string]interface{}{
					{
						"value":   num,
						"user_id": idUser,
					},
				},
			},
		}

		// Envio do JSON
		retorno, err = sendMessageToMiddlewareSocketIO(hostUrl, middlewareID, jsonObject)
		if err != nil {
			return err
		}

		fmt.Println("retorno do user_set_image_list: %v", retorno)
	}

	if tag != "" {
		jsonObject = map[string]interface{}{
			"url":    "http://" + host + "/create_objects.fcgi?session=" + sessionValue,
			"Method": "POST",
			"idTask": "_create_tag",
			"devID":  devID,
			"json": map[string]interface{}{
				"object": "uhf_tags",
				"values": []map[string]interface{}{
					{
						"value":   tag,
						"user_id": idUser,
					},
				},
			},
		}

		// Envio do JSON
		retorno, err = sendMessageToMiddlewareSocketIO(hostUrl, middlewareID, jsonObject)
		if err != nil {
			return err
		}

		fmt.Println("retorno do user_set_image_list: %v", retorno)
	}

	if qrcode != "" {
		jsonObject = map[string]interface{}{
			"url":    "http://" + host + "/create_objects.fcgi?session=" + sessionValue,
			"Method": "POST",
			"idTask": "_create_qrcode",
			"devID":  devID,
			"json": map[string]interface{}{
				"object": "qrcodes",
				"values": []map[string]interface{}{
					{
						"value":   qrcode,
						"user_id": idUser,
					},
				},
			},
		}

		// Envio do JSON
		retorno, err = sendMessageToMiddlewareSocketIO(hostUrl, middlewareID, jsonObject)
		if err != nil {
			return err
		}

		fmt.Println("retorno do _create_qrcode: %v", retorno)
	}

	return nil
}

// Function to delete user with idUserTerminal
func deleteUserv2(hostUrl string, middlewareID string, host string, username string, senha string, user int64) error {

	// Mandar o login
	// Objeto JSON a ser enviado
	jsonObject := map[string]interface{}{
		"url":    "http://" + host + "/login.fcgi",
		"Method": "POST",
		"idTask": "_login",
		"json": map[string]string{
			"login":    username,
			"password": senha,
		},
	}

	// Envio do JSON
	retorno, err := sendMessageToMiddlewareSocketIO(hostUrl, middlewareID, jsonObject)
	if err != nil {
		return err
	}

	message, ok := getPayload(retorno)
	if !ok {
		return fmt.Errorf("error: payload not send: %v", retorno)
	}

	sessionValue, ok := message["session"].(string)

	// Imprimindo o valor de "session"
	fmt.Println("Valor da session:", sessionValue)

	// Criando um usuário
	jsonObject = map[string]interface{}{
		"url":    "http://" + host + "/destroy_objects.fcgi?session=" + sessionValue,
		"Method": "POST",
		"idTask": "_delete_user",
		"json": map[string]interface{}{
			"object": "users",
			"where": map[string]interface{}{
				"users": map[string]interface{}{"id": user},
			},
		},
	}

	// Envio do JSON
	retorno, err = sendMessageToMiddlewareSocketIO(hostUrl, middlewareID, jsonObject)
	if err != nil {
		return err
	}

	payload, ok := getPayload(retorno)
	if !ok {
		return fmt.Errorf("error: payload not send: %v", retorno)
	}

	if _, ok = payload["changes"].(float64); !ok {
		return fmt.Errorf("error: changes not send: %v", payload)
	}

	return nil
}
