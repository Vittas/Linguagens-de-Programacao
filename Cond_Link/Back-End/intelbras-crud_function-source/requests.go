package function

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getInfoHardwareByDevID(devID string) (map[string]interface{}, error) {
	url := "https://us-central1-earnest-cosmos-175020.cloudfunctions.net/resynchronize-db-functions/getInfoHardwareByDevID"
	method := "POST"

	// Dados a serem enviados no corpo da solicitação, usando um map
	payload := map[string]string{
		"devID": devID,
	}

	// Converter o map para JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Criar uma nova solicitação HTTP
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Adicionar cabeçalhos à solicitação
	req.Header.Add("X-GUARDIAN-KEY", "12DE-76XD-98HJ-00PO")
	req.Header.Add("Content-Type", "application/json")

	// Executar a solicitação
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	// Ler a resposta
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Imprimir a resposta
	fmt.Println(string(body))

	// Decodificar a resposta JSON para um mapa
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Erro ao decodificar a resposta:", err)
		return nil, err
	}
	return result, nil
}

// Send a payload to server of socket io
func sendMessageToMiddlewareSocketIO(host string, middlewareID string, payload map[string]interface{}) (map[string]interface{}, error) {
	// alarm.condlink.com.br:8852
	url := "http://" + host + "/v1/hardware/" + middlewareID + "/send"
	method := "POST"

	fmt.Println("sendMessageToMiddlewareSocketIO " + url)
	// Converter o map para JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Criar uma nova solicitação HTTP
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Adicionar cabeçalhos à solicitação
	// req.Header.Add("X-GUARDIAN-KEY", "12DE-76XD-98HJ-00PO")
	req.Header.Add("Content-Type", "application/json")

	// Executar a solicitação
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	// Ler a resposta
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Imprimir a resposta
	fmt.Println(string(body))

	// Decodificar a resposta JSON para um mapa
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Erro ao decodificar a resposta:", err)
		return nil, err
	}
	return result, nil
}
