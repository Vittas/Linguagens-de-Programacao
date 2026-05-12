package function

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

// func main() {
// 	http.HandleFunc("/", Animais)
// 	http.ListenAndServe(":8080", nil)
// }

type Animal struct {
	Nome          string         `json:"nome"`
	Unidade       string         `json:"unidade"`
	Raca          string         `json:"raca"`
	Tipo          string         `json:"tipo"`
	NomeCondomino string         `json:"nome_condomino"`
	Foto          sql.NullString `json:"foto"`
}

func Animais(w http.ResponseWriter, r *http.Request) {

	allowedOrigin := "https://portaria.condlink.com.br"

	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-GUARDIAN-KEY")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)

	// Define o content type da resposta para json
	w.Header().Set("Content-Type", "application/json")

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

	// Tenta abrir uma conexão com o banco de dados
	db, err := connectToDatabase()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"error": "Erro ao conectar ao banco de dados"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Fecha a conexão com o banco de dados no final da execução da function
	defer db.Close()

	var requestBody RequestBody
	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Erro ao decodificar o JSON do corpo da solicitação", http.StatusBadRequest)
		return
	}

	var animais []Animal

	animais, err = getAnimais(db, requestBody.IdCondominio, requestBody.Query)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(animais)

}

func getAnimais(db *sql.DB, idCondominio, query string) ([]Animal, error) {

	query = "%"+query+"%"

	// query := `select c.unidade, c.nome, a.tipo,a.raca,a.nome, a.foto from animais_condominos a inner join condomino c on a.id_condomino= c.id where c.id_condominio=?;`
	search := `SELECT 
    c.unidade, 
    c.nome, 
    a.tipo, 
    a.raca, 
    a.nome, 
    a.foto 
FROM 
    animais_condominos a 
INNER JOIN 
    condomino c 
ON 
    a.id_condomino = c.id 
WHERE 
    c.id_condominio = ? 
    AND (
        a.nome LIKE ? 
        OR c.unidade LIKE ? 
        OR a.raca LIKE ? 
        OR a.tipo LIKE ?
    );
`

	rows, err := db.Query(search, idCondominio, query, query, query, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var animais []Animal
	for rows.Next() {
		var animal Animal
		err := rows.Scan(&animal.Unidade, &animal.NomeCondomino, &animal.Tipo, &animal.Raca, &animal.Nome, &animal.Foto)
		if err != nil {
			return nil, err
		}
		animais = append(animais, animal)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return animais, err

}

func connectToDatabase() (*sql.DB, error) {

	var (
		dbUser    = "joao"
		dbPwd     = "xxx"
		dbName    = "multicondlink"
		dbPort    = "3306"
		dbTCPHost = "10.5.16.3"
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPwd, dbTCPHost, dbPort, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error opening the database: ", err)
		return nil, fmt.Errorf("failed to open the database: %v", err)
	}

	// check if the connection is actually working by pinging it
	err = db.Ping()
	/* PING
	Confere se a conexão com o servidor está funcionando ou não.
	Se ela tiver caído, tenta realizar uma conexão automática.
	Esta função pode ser usada em scripts que permanecem inativos por um longo tempo,
	para verificar se o servidor fechou ou não a conexão e reconectar se necessário.
	*/

	if err != nil {
		log.Printf("Error pinging the database: %v\n", err)
		return nil, fmt.Errorf("failed to ping the database: %v", err)
	}

	fmt.Println("Conectado ao banco de dados")

	return db, nil
}

type RequestBody struct {
	IdCondominio string `json:"id_condominio"`
	Query        string `json:"query"`
}

func (t Animal) MarshalJSON() ([]byte, error) {
	type Alias Animal
	aux := struct {
		Foto string `json:"foto,omitempty"`
		Alias
	}{
		Alias: Alias(t),
	}

	// Trata campos NULL

	if t.Foto.Valid {
		aux.Foto = t.Foto.String
	}

	return json.Marshal(aux)
}
