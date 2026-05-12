package function

import (
	"net/http"
	"strconv"
	"strings"
)

// Create a single user
func handlerCreateUserHardware(devID, idUser, name, url_photo, card, idGroup, tag, qrcode string) (map[string]interface{}, int) {
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

	// password
	password, ok := result["senha"].(string)
	if !ok {
		return map[string]interface{}{
			"error": "password not found",
		}, http.StatusBadRequest
	}

	integerNum, err := strconv.ParseInt(idUser, 10, 64)
	if err != nil {
		return map[string]interface{}{
			"error": "Error to convert id of user:" + err.Error(),
		}, http.StatusBadRequest
	}

	idGroupInt, err := strconv.ParseInt(idGroup, 10, 64)
	if err != nil {
		return map[string]interface{}{
			"error": "Error to convert id of user:" + err.Error(),
		}, http.StatusBadRequest
	}

	var retorno interface{}
	retorno = "user created"
	if strings.HasPrefix(idMiddleware, "v2_") {
		err = createUserv2("alarm.condlink.com.br:8852", url_photo, card, integerNum, name, idMiddleware, host, username, password, idGroupInt, tag, qrcode, devID)
	} else {
		retorno, err = createUser(url_photo, card, integerNum, name, idMiddleware, host, username, password, idGroupInt)
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
		"name":         name,
		"card":         card,
		"idGroup":      idGroup,
		"message":      retorno,
	}, http.StatusOK
}

// Delete a single user
func handlerDeleteUserHardware(devID string, idUser string) (map[string]interface{}, int) {
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

	integerNum, err := strconv.ParseInt(idUser, 10, 64)
	if err != nil {
		return map[string]interface{}{
			"error": "Error to convert id of user:" + err.Error(),
		}, http.StatusBadRequest
	}

	var retorno interface{}
	retorno = "user deleted"
	if strings.HasPrefix(idMiddleware, "v2_") {
		err = deleteUserv2("alarm.condlink.com.br:8852", idMiddleware, host, username, password, integerNum)
	} else {
		retorno, err = deleteUserControlID(idMiddleware, host, username, password, integerNum)
	}
	if err != nil {
		return map[string]interface{}{
			"error": "Error to create user:" + err.Error(),
		}, http.StatusBadRequest
	}

	return map[string]interface{}{
		"idMiddleware": idMiddleware,
		"host":         host,
		"message":      retorno,
	}, http.StatusOK
}
