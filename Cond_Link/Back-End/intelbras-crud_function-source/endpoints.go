package function

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func _downloadAndConvertToBase64(url string) (string, error) {
	const maxSizeKB = 100.0

	// Fazendo a requisição HTTP
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("erro ao baixar a imagem: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("falha ao baixar a imagem: status code %d", resp.StatusCode)
	}

	// Lendo os dados da imagem
	imageData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("erro ao ler os dados da imagem: %v", err)
	}

	// Calculando o tamanho da imagem em KB
	imageSizeKB := float64(len(imageData)) / 1024.0

	// Verificando se o tamanho da imagem excede o limite
	if imageSizeKB > maxSizeKB {
		return "", fmt.Errorf("tamanho da imagem excede o limite de %.2f KB: %.2f KB", maxSizeKB, imageSizeKB)
	}

	// Convertendo para base64
	base64Image := base64.StdEncoding.EncodeToString(imageData)
	return base64Image, nil
}

func _convertStringToMap(input string, sep string) map[string]interface{} {
	// Replace \r\n with actual newlines
	processedInput := strings.ReplaceAll(input, `\r\n`, "\n")

	// Split the string by newlines
	lines := strings.Split(processedInput, "\n")

	// Create a map to hold the key-value pairs
	data := make(map[string]interface{})

	// Process each line
	for _, line := range lines {
		if line == "" {
			continue
		}
		// Split by the first '=' character
		parts := strings.SplitN(line, sep, 2)
		if len(parts) == 2 {
			// Remove any trailing \r from the value part
			key := parts[0]
			value := strings.TrimSuffix(parts[1], "\r")
			if sep == ":" {
				key = strings.TrimSpace(key)
				// Usar strings.Map para remover os caracteres de escape
				key = strings.Map(func(r rune) rune {
					if r == '\\' || r == '"' {
						return -1 // Remove o caractere
					}
					return r
				}, key)
				value = strings.ReplaceAll(value, ",", "")
				value = strings.TrimSpace(value)
			}
			data[key] = value
		}
	}

	return data
}

// Desconpacta o result e retorna o host, username, passwrod e idMiddleware
func _config(result map[string]interface{}) (string, string, string, string, error) {
	// variables
	var username string
	var password string
	var host string
	var idMiddleware string

	// username
	username, ok := result["username"].(string)
	if !ok {
		return "", "", "", "", fmt.Errorf("username not found")
	}

	password, ok = result["senha"].(string)
	if !ok {
		return "", "", "", "", fmt.Errorf("senha not found")
	}

	host, ok = result["host"].(string)
	if !ok {
		return "", "", "", "", fmt.Errorf("host not found")
	}

	idMiddleware, ok = result["idMiddleware"].(string)
	if !ok {
		return "", "", "", "", fmt.Errorf("idMiddleware not found")
	}
	return host, username, password, idMiddleware, nil
}

// ConvertInterfaceToIntSlice converte uma interface para um slice de inteiros.
func _convertInterfaceToIntSlice(input interface{}) ([]int, error) {
	var result []int

	switch v := input.(type) {
	case string:
		// Remove os colchetes e espaços extras da string
		cleaned := strings.Trim(v, "[]")
		cleaned = strings.TrimSpace(cleaned)

		if cleaned == "" {
			return []int{}, nil
		}

		// Divide a string em partes
		strValues := strings.Split(cleaned, ",")
		for _, strVal := range strValues {
			strVal = strings.TrimSpace(strVal)
			intVal, err := strconv.Atoi(strVal)
			if err != nil {
				return nil, fmt.Errorf("falha ao converter '%s' para int: %v", strVal, err)
			}
			result = append(result, intVal)
		}

	case []interface{}:
		for _, item := range v {
			switch t := item.(type) {
			case string:
				intVal, err := strconv.Atoi(t)
				if err != nil {
					return nil, fmt.Errorf("falha ao converter '%s' para int: %v", t, err)
				}
				result = append(result, intVal)
			case float64:
				// Converte float64 para int
				result = append(result, int(t))
			case int:
				// Adiciona o inteiro na lista
				result = append(result, t)
			default:
				return nil, fmt.Errorf("tipo não suportado: %T", item)
			}
		}

	default:
		return nil, fmt.Errorf("tipo de entrada não suportado: %T", v)
	}

	return result, nil
}

// Returns the users registered
func handlerGetAllUsers(devID string) (map[string]interface{}, int) {
	result, err := getInfoHardwareByDevID(devID)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, http.StatusInternalServerError
	}
	host, username, password, idMiddleware, err := _config(result)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, http.StatusInternalServerError
	}

	// make some request because we have some informations
	fmt.Println(username, password, host, idMiddleware)
	returned, err := getAllUsers(host, idMiddleware, username, password)

	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, http.StatusInternalServerError
	}

	map2Return := _convertStringToMap(returned, "=")

	return map2Return, http.StatusOK
}

// Returns the message to create user
func handlerInsertUser(devID string, user map[string]interface{}) (map[string]interface{}, int) {
	result, err := getInfoHardwareByDevID(devID)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, http.StatusInternalServerError
	}
	host, username, password, idMiddleware, err := _config(result)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, http.StatusInternalServerError
	}

	// make some request because we have some informations
	fmt.Println(username, password, host, idMiddleware)

	// treatment of user
	// Tratamento da lista de usuários
	userList, ok := user["UserList"].([]interface{})
	if !ok {
		fmt.Println("Campo 'UserList' não é do tipo []interface{}")
		return map[string]interface{}{
			"error": "field user has no UserList like []interface{}",
		}, http.StatusInternalServerError
	}
	for index, uss := range userList {
		ussMap, ok := uss.(map[string]interface{})
		if !ok {
			fmt.Printf("Usuário %d não é do tipo map[string]interface{}\n", index)
			continue
		}
		if doors, ok := ussMap["Doors"]; ok {
			convertedDoors, err := _convertInterfaceToIntSlice(doors)
			if err != nil {
				fmt.Printf("Erro ao converter 'Doors' para usuário %d: %v\n", index, err)
			} else {
				ussMap["Doors"] = convertedDoors
			}
		}

		if timeSections, ok := ussMap["TimeSections"]; ok {
			convertedTimeSections, err := _convertInterfaceToIntSlice(timeSections)
			if err != nil {
				fmt.Printf("Erro ao converter 'TimeSections' para usuário %d: %v\n", index, err)
			} else {
				ussMap["TimeSections"] = convertedTimeSections
			}
		}

		userList[index] = ussMap
	}
	user["UserList"] = userList
	returned, err := registerUser("alarm.condlink.com.br:8852", host, idMiddleware, username, password, user, devID)

	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, http.StatusInternalServerError
	}

	return map[string]interface{}{
		"message": returned,
	}, http.StatusOK
}

// Returns the message to create some face to some user
func handlerInsertFaceUser(devID string, user map[string]interface{}) (map[string]interface{}, int) {
	result, err := getInfoHardwareByDevID(devID)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, http.StatusInternalServerError
	}
	host, username, password, idMiddleware, err := _config(result)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, http.StatusInternalServerError
	}

	// make some request because we have some informations
	fmt.Println(username, password, host, idMiddleware)
	returned, err := registerFace("alarm.condlink.com.br:8852", host, idMiddleware, username, password, user, devID)

	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, http.StatusInternalServerError
	}

	return map[string]interface{}{
		"message": returned,
	}, http.StatusOK
}

// Returns the message to create some face to some user
func handlerInsertFaceUserWithURL(devID string, user map[string]interface{}) (map[string]interface{}, int) {
	result, err := getInfoHardwareByDevID(devID)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, http.StatusInternalServerError
	}
	host, username, password, idMiddleware, err := _config(result)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, http.StatusInternalServerError
	}

	// make some request because we have some informations
	fmt.Println(username, password, host, idMiddleware)

	// Handler with url
	faceList, ok := user["FaceList"].([]interface{})
	if !ok {
		return map[string]interface{}{
			"error": "FaceList is not an array of interface",
		}, http.StatusInternalServerError
	}

	var arrayFacesBase64 []map[string]interface{}
	var faces_with_problem []interface{}

	// Iterar sobre o FaceList
	for _, item := range faceList {
		face, ok := item.(map[string]interface{})
		if !ok {
			fmt.Println("Erro: item do 'FaceList' não é um objeto")
			continue
		}

		// Obter UserID e verificar tipo
		userID, ok := face["UserID"].(string)
		if !ok {
			fmt.Println("Erro: 'UserID' não é uma string")
			continue
		}

		// Obter PhotoData e verificar tipo
		photoData, ok := face["PhotoData"].([]interface{})
		if !ok {
			fmt.Println("Erro: 'PhotoData' não é uma lista de interface")
			continue
		}

		var base64s []string
		// Iterar sobre as URLs em PhotoData
		for _, urlItem := range photoData {
			url, ok := urlItem.(string)
			if !ok {
				fmt.Println("Erro: item do 'PhotoData' não é uma string")
				continue
			}

			// Fazer download e converter para base64
			base64Image, err := _downloadAndConvertToBase64(url)
			if err != nil {
				fmt.Println("Erro ao converter a imagem:", err)
				faces_with_problem = append(faces_with_problem, "Image:"+url+" error: "+err.Error())
				continue
			}
			base64s = append(base64s, base64Image)
		}

		arrayFacesBase64 = append(arrayFacesBase64, map[string]interface{}{
			"UserID":    userID,
			"PhotoData": base64s,
		})
	}

	mapUserNew := map[string]interface{}{
		"FaceList": arrayFacesBase64,
	}

	returned, err := registerFace("alarm.condlink.com.br:8852", host, idMiddleware, username, password, mapUserNew, devID)

	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, http.StatusInternalServerError
	}

	return map[string]interface{}{
		"message":            returned,
		"faces_with_problem": faces_with_problem,
	}, http.StatusOK
}

// Returns the message to create some face to some user
func handlerInsertCardUser(devID string, cards map[string]interface{}) (map[string]interface{}, int) {
	result, err := getInfoHardwareByDevID(devID)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, http.StatusInternalServerError
	}
	host, username, password, idMiddleware, err := _config(result)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, http.StatusInternalServerError
	}

	// make some request because we have some informations
	fmt.Println(username, password, host, idMiddleware)
	returned, err := registerCard(host, idMiddleware, username, password, cards)

	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, http.StatusInternalServerError
	}

	return map[string]interface{}{
		"message": returned,
	}, http.StatusOK
}

// Returns the message to delete some usar
func handlerDeleteUser(devID string, userID string) (map[string]interface{}, int) {
	result, err := getInfoHardwareByDevID(devID)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, http.StatusInternalServerError
	}
	host, username, password, idMiddleware, err := _config(result)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, http.StatusInternalServerError
	}

	// make some request because we have some informations
	fmt.Println(username, password, host, idMiddleware)
	var returned interface{}
	if strings.HasPrefix(idMiddleware, "v2_1_2") {
		returned, err = deleteUserv2("cloudrun.condlink.com.br:8852", userID, idMiddleware, host, username, password)
	} else {
		returned, err = deleteUserv2("alarm.condlink.com.br:8852", userID, idMiddleware, host, username, password)
	}

	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, http.StatusInternalServerError
	}

	return map[string]interface{}{
		"message": returned,
	}, http.StatusOK
}

// Returns the user of ID
func handlerGetUser(devID string, userID string) (map[string]interface{}, int) {
	result, err := getInfoHardwareByDevID(devID)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, http.StatusInternalServerError
	}
	host, username, password, idMiddleware, err := _config(result)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, http.StatusInternalServerError
	}

	// make some request because we have some informations
	fmt.Println(username, password, host, idMiddleware)
	returned, err := getUser(host, idMiddleware, username, password, userID)

	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}, http.StatusInternalServerError
	}

	return map[string]interface{}{
		"message": returned,
	}, http.StatusOK
}

// ===================  V2  ==============================
// Create a single user
func handlerCreateUserHardware(devID string, idUser string, name string, photo string, card string, idGroup string) (map[string]interface{}, int) {
	result, err := getInfoHardwareByDevID(devID)
	if err != nil {
		return map[string]interface{}{
			"error": "get info about hardware returns with error " + err.Error(),
		}, http.StatusBadRequest
	}

	// variables

	// idMiddleware
	idMiddleware, ok := result["idMiddleware"].(string)
	if !ok {
		return map[string]interface{}{
			"error": "idMiddleware not found",
		}, http.StatusBadRequest
	}

	// host
	host, ok := result["host"].(string)
	if !ok {
		return map[string]interface{}{
			"error": "host not found",
		}, http.StatusBadRequest
	}

	// username
	username, ok := result["username"].(string)
	if !ok {
		return map[string]interface{}{
			"error": "username not found",
		}, http.StatusBadRequest
	}

	if idGroup == "" {
		idGroup, ok = result["group_contigencia"].(string)
		if !ok {
			return map[string]interface{}{
				"error": "group_contigencia not found",
			}, http.StatusBadRequest
		}
	}

	idGroupInt, err := strconv.ParseInt(idGroup, 10, 64)
	if err != nil {
		return map[string]interface{}{
			"error": "Error to convert id of user:" + err.Error(),
		}, http.StatusBadRequest
	}

	// password
	password, ok := result["senha"].(string)
	if !ok {
		return map[string]interface{}{
			"error": "password not found",
		}, http.StatusBadRequest
	}

	// Create user on alarm.condlink.com.br
	retorno, err := createUserv2("alarm.condlink.com.br:8852", photo, card, idUser, name, idMiddleware, host, username, password, idGroupInt)

	if err != nil {
		return map[string]interface{}{
			"error": "Error to create user: " + err.Error(),
		}, http.StatusBadRequest

	}

	return map[string]interface{}{
		"idMiddleware": idMiddleware,
		"host":         host,
		"username":     username,
		"name":         name,
		"card":         card,
		"idGroup":      idGroup,
		"message": "user created with success",
		"infos":      retorno,
	}, http.StatusOK
}

// Delete a single user
func handlerDeleteHardware(devID string, idUser string) (map[string]interface{}, int) {
	result, err := getInfoHardwareByDevID(devID)
	if err != nil {
		return map[string]interface{}{
			"error": "get info about hardware returns with error " + err.Error(),
		}, http.StatusBadRequest
	}

	// variables
	// idMiddleware
	idMiddleware, ok := result["idMiddleware"].(string)
	if !ok {
		return map[string]interface{}{
			"error": "idMiddleware not found",
		}, http.StatusBadRequest
	}

	// host
	host, ok := result["host"].(string)
	if !ok {
		return map[string]interface{}{
			"error": "host not found",
		}, http.StatusBadRequest
	}

	// username
	username, ok := result["username"].(string)
	if !ok {
		return map[string]interface{}{
			"error": "username not found",
		}, http.StatusBadRequest
	}

	// password
	password, ok := result["senha"].(string)
	if !ok {
		return map[string]interface{}{
			"error": "password not found",
		}, http.StatusBadRequest
	}

	var retorno interface{}
	if strings.HasPrefix(idMiddleware, "v2_1_2_") {
		retorno, err = deleteUserv2("cloudrun.condlink.com.br:8852", idUser, idMiddleware, host, username, password)
	} else {
		retorno, err = deleteUserv2("alarm.condlink.com.br:8852", idUser, idMiddleware, host, username, password)
	}

	if err != nil {
		return map[string]interface{}{
			"error": "Error to create user: " + err.Error(),
		}, http.StatusBadRequest

	}

	return map[string]interface{}{
		"idMiddleware": idMiddleware,
		"host":         host,
		"username":     username,
		"message":      retorno,
	}, http.StatusOK
}

// Get all users
func handlerGetAllUsersHardware(devID string) (map[string]interface{}, int) {
	result, err := getInfoHardwareByDevID(devID)
	if err != nil {
		return map[string]interface{}{
			"error": "get info about hardware returns with error " + err.Error(),
		}, http.StatusBadRequest
	}

	// variables
	// idMiddleware
	idMiddleware, ok := result["idMiddleware"].(string)
	if !ok {
		return map[string]interface{}{
			"error": "idMiddleware not found",
		}, http.StatusBadRequest
	}

	// host
	host, ok := result["host"].(string)
	if !ok {
		return map[string]interface{}{
			"error": "host not found",
		}, http.StatusBadRequest
	}

	// username
	username, ok := result["username"].(string)
	if !ok {
		return map[string]interface{}{
			"error": "username not found",
		}, http.StatusBadRequest
	}

	// password
	password, ok := result["senha"].(string)
	if !ok {
		return map[string]interface{}{
			"error": "password not found",
		}, http.StatusBadRequest
	}

	var retorno interface{}
	if strings.HasPrefix(idMiddleware, "v2_1_2_") {
		retorno, err = getUsersv2("cloudrun.condlink.com.br:8852", idMiddleware, host, username, password, devID)
	} else {
		retorno, err = getUsersv2("alarm.condlink.com.br:8852", idMiddleware, host, username, password, devID)
	}

	if err != nil {
		return map[string]interface{}{
			"error": "Error to create user: " + err.Error(),
		}, http.StatusBadRequest

	}

	return map[string]interface{}{
		"idMiddleware": idMiddleware,
		"host":         host,
		"username":     username,
		"users":        retorno,
	}, http.StatusOK
}

// Get a single user
func handlerGetUserHardware(devID string, idUser string) (map[string]interface{}, int) {
	result, err := getInfoHardwareByDevID(devID)
	if err != nil {
		return map[string]interface{}{
			"error": "get info about hardware returns with error " + err.Error(),
		}, http.StatusBadRequest
	}

	// variables
	// idMiddleware
	idMiddleware, ok := result["idMiddleware"].(string)
	if !ok {
		return map[string]interface{}{
			"error": "idMiddleware not found",
		}, http.StatusBadRequest
	}

	// host
	host, ok := result["host"].(string)
	if !ok {
		return map[string]interface{}{
			"error": "host not found",
		}, http.StatusBadRequest
	}

	// username
	username, ok := result["username"].(string)
	if !ok {
		return map[string]interface{}{
			"error": "username not found",
		}, http.StatusBadRequest
	}

	// password
	password, ok := result["senha"].(string)
	if !ok {
		return map[string]interface{}{
			"error": "password not found",
		}, http.StatusBadRequest
	}

	var retorno interface{}
	if strings.HasPrefix(idMiddleware, "v2_1_2_") {
		retorno, err = getUserv2("cloudrun.condlink.com.br:8852", idMiddleware, host, username, password, idUser)
	} else {
		retorno, err = getUserv2("alarm.condlink.com.br:8852", idMiddleware, host, username, password, idUser)
	}
	if err != nil {
		return map[string]interface{}{
			"error": "Error to create user: " + err.Error(),
		}, http.StatusBadRequest

	}

	return map[string]interface{}{
		"idMiddleware": idMiddleware,
		"host":         host,
		"username":     username,
		"user":         retorno,
	}, http.StatusOK
}
