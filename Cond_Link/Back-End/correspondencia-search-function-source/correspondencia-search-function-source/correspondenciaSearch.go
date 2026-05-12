package function

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Tickets struct {
	Id              sql.NullString `json:"id"`
	Nome            sql.NullString `json:"nome"`
	Unidade         sql.NullString `json:"unidade"`
	ArquivoAnexo    sql.NullString `json:"arquivo_anexo"`
	DataInicio      sql.NullString `json:"data_inicio"`
	DataAtualizacao sql.NullString `json:"data_atualizacao"`
	RemetenteNome   sql.NullString `json:"remetente_nome"`
	Status          sql.NullString `json:"status"`
	IsLido          sql.NullString `json:"is_lido"`
	TipoRemetente   sql.NullString `json:"tipo_remetente"`
	IdRemetente     sql.NullInt64  `json:"id_remetente"`
	IdDestinatario  sql.NullString `json:"id_destinatario"`
	Assunto         sql.NullString `json:"assunto"`
	Mensagem        sql.NullString `json:"mensagem"`
	CampoExtra1     sql.NullString `json:"campo_extra1"`
}

// func main() {
// 	http.HandleFunc("/", Correspondencias)
// 	http.ListenAndServe(":8080", nil)
// }

// Correspondencias é o ponto de entrada da Cloud Function :D
func Correspondencias(w http.ResponseWriter, r *http.Request) {

	allowedOrigin := "https://portaria.condlink.com.br"
	origin := r.Header.Get("Origin")

	// Verifica se a origem é permitida
	if origin == allowedOrigin || strings.HasPrefix(origin, "http://localhost:") || strings.HasPrefix(origin, "http://127.0.0.1:") {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	} else {
		w.Header().Set("Access-Control-Allow-Origin", "")
	}

	if r.Method == http.MethodOptions {
		if origin == allowedOrigin || strings.HasPrefix(origin, "http://localhost:") || strings.HasPrefix(origin, "http://127.0.0.1:") {
			w.Header().Set("Access-Control-Allow-Methods", "POST")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-GUARDIAN-KEY")
			w.Header().Set("Access-Control-Max-Age", "3600")
			w.WriteHeader(http.StatusNoContent)
			return
		} else {
			w.WriteHeader(http.StatusForbidden)
			return
		}
	}

	// Define o Content-Type
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

	// ------ AQUI COMEÇA A LÓGICA DA FUNCTION ------

	var requestJson RequestJson
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&requestJson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var query string = requestJson.Query
	var status string = requestJson.Status
	var idCondominium string = requestJson.IdCondominio
	var dataInicio string = requestJson.DataInicio
	var dataFim string = requestJson.DataFim

	var tickets []Tickets

	if requestJson.Status == "" {
		tickets, err = getTickets(db, idCondominium, query, dataInicio, dataFim)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"message": err.Error(),
			})
			return
		}

		// verifica se retornou 1 ticket
		if len(tickets) == 1 {
			ticket := tickets[0]
			// get unidade
			unidade := ticket.Unidade.String

			// getTickets where unidade is the query
			tickets, err = getTickets(db, idCondominium, unidade, dataInicio, dataFim)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"success": false,
					"message": err.Error(),
				})
				return
			}

			
		}
	} else {
		tickets, err = getTicketsByStatus(db, idCondominium, query, status, dataInicio, dataFim)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"message": err.Error(),
			})
			return
		}

		// verifica se retornou 1 ticket
		if len(tickets) == 1 {
			ticket := tickets[0]
			// get unidade
			unidade := ticket.Unidade.String

			// getTickets where unidade is the query
			tickets, err = getTickets(db, idCondominium, unidade, dataInicio, dataFim)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"success": false,
					"message": err.Error(),
				})
				return
			}
		}


	}

	if len(tickets) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Nenhum registro encontrado",
		})
		return
	}

	json.NewEncoder(w).Encode(tickets)

}

func getTickets(db *sql.DB, idCondominium, query, dataInicio, dataFim string) ([]Tickets, error) {

	query = query + "%"
	dataFim = dataFim + " 23:59:59"

	
	search := `SELECT t.id, c.nome, c.unidade, mensagem, t.arquivo_anexo, data_inicio, remetente_nome, t.status, t.isLido, t.id_remetente, t.id_destinatario,t.campo_extra1, t.data_atualizacao
    FROM ticket t 
    INNER JOIN condomino c ON t.id_destinatario = c.id 
    WHERE t.id_condominio = ? 
        
        AND tipo_destinatario = "condomino" 
        AND tipo_servico_remetente = "Aviso de Correspondência" 
        AND (c.nome LIKE ? OR c.unidade LIKE ? OR RIGHT(CAST(t.id AS CHAR), 5) LIKE ? OR c.endereco LIKE ?)
		AND data_inicio BETWEEN ? AND ? 
    ORDER BY t.data_inicio DESC;`

	rows, err := db.Query(search, idCondominium, query, query, query, query, dataInicio, dataFim)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tickets := []Tickets{}
	for rows.Next() {
		ticket := Tickets{}
		err := rows.Scan(&ticket.Id, &ticket.Nome, &ticket.Unidade,
			&ticket.Mensagem, &ticket.ArquivoAnexo, &ticket.DataInicio, &ticket.RemetenteNome, &ticket.Status, &ticket.IsLido, &ticket.IdRemetente, &ticket.IdDestinatario, &ticket.CampoExtra1, &ticket.DataAtualizacao)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, ticket)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tickets, nil
}

func getTicketsByStatus(db *sql.DB, idCondominium, query, status, dataInicio, dataFim string) ([]Tickets, error) {

	query = query + "%"
	dataFim = dataFim + " 23:59:59"

	fmt.Println(idCondominium, query, status, dataInicio, dataFim)
	search := `SELECT t.id, c.nome, c.unidade, mensagem, t.arquivo_anexo, data_inicio, remetente_nome, t.status, t.isLido, t.id_remetente, t.id_destinatario,t.campo_extra1, t.data_atualizacao
    FROM ticket t 
    INNER JOIN condomino c ON t.id_destinatario = c.id 
    WHERE t.id_condominio = ? 
        
        AND tipo_destinatario = "condomino" 
        AND tipo_servico_remetente = "Aviso de Correspondência" 
        AND (c.nome LIKE ? OR c.unidade LIKE ? OR RIGHT(CAST(t.id AS CHAR), 5) LIKE ? OR c.endereco LIKE ?)
		AND t.status = ?
		AND data_inicio BETWEEN ? AND ? 
    ORDER BY t.data_inicio DESC;`

	rows, err := db.Query(search, idCondominium, query, query, query, query, status, dataInicio, dataFim)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tickets := []Tickets{}
	for rows.Next() {
		ticket := Tickets{}
		err := rows.Scan(&ticket.Id, &ticket.Nome, &ticket.Unidade,
			&ticket.Mensagem, &ticket.ArquivoAnexo, &ticket.DataInicio, &ticket.RemetenteNome, &ticket.Status, &ticket.IsLido, &ticket.IdRemetente, &ticket.IdDestinatario, &ticket.CampoExtra1, &ticket.DataAtualizacao)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, ticket)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tickets, nil
}

func connectToDatabase() (*sql.DB, error) {

	var (
		dbUser    = "joao"
		dbPwd     = "xxx"
		dbName    = "multicondlink"
		dbPort    = "3306"
		dbTCPHost = "10.5.16.3"
		// dbTCPHost = "mysql1.condlink.com.br"
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

func (t Tickets) MarshalJSON() ([]byte, error) {
	type Alias Tickets
	aux := struct {
		Id              string `json:"id,omitempty"`
		Nome            string `json:"nome,omitempty"`
		Unidade         string `json:"unidade,omitempty"`
		ArquivoAnexo    string `json:"arquivo_anexo,omitempty"`
		DataInicio      string `json:"data_inicio,omitempty"`
		DataAtualizacao string `json:"data_atualizacao,omitempty"`
		RemetenteNome   string `json:"remetente_nome,omitempty"`
		Status          string `json:"status,omitempty"`
		IsLido          string `json:"is_lido,omitempty"`
		TipoRemetente   string `json:"tipo_remetente,omitempty"`
		IdRemetente     int64  `json:"id_remetente,omitempty"`
		Assunto         string `json:"assunto,omitempty"`
		Mensagem        string `json:"mensagem,omitempty"`
		IdDestinatario  string `json:"id_destinatario,omitempty"`
		CampoExtra1     string `json:"campo_extra1,omitempty"`
		Alias
	}{
		Alias: Alias(t),
	}

	// Trata campos NULL
	if t.Id.Valid {
		aux.Id = t.Id.String
	}
	if t.Nome.Valid {
		aux.Nome = t.Nome.String
	}
	if t.Unidade.Valid {
		aux.Unidade = t.Unidade.String
	}
	if t.ArquivoAnexo.Valid {
		aux.ArquivoAnexo = t.ArquivoAnexo.String
	}
	if t.DataInicio.Valid {
		aux.DataInicio = t.DataInicio.String
	}
	if t.DataAtualizacao.Valid {
		aux.DataAtualizacao = t.DataAtualizacao.String
	}
	if t.RemetenteNome.Valid {
		aux.RemetenteNome = t.RemetenteNome.String
	}
	if t.Status.Valid {
		aux.Status = t.Status.String
	}
	if t.IsLido.Valid {
		aux.IsLido = t.IsLido.String
	}
	if t.TipoRemetente.Valid {
		aux.TipoRemetente = t.TipoRemetente.String
	}
	if t.IdRemetente.Valid {
		aux.IdRemetente = t.IdRemetente.Int64
	}
	if t.Assunto.Valid {
		aux.Assunto = t.Assunto.String
	}
	if t.Mensagem.Valid {
		aux.Mensagem = t.Mensagem.String
	}
	if t.IdDestinatario.Valid {
		aux.IdDestinatario = t.IdDestinatario.String
	}
	if t.CampoExtra1.Valid {
		aux.CampoExtra1 = t.CampoExtra1.String
	}

	return json.Marshal(aux)
}

type RequestJson struct {
	IdCondominio string `json:"id_condominio"`
	Query        string `json:"query"`
	Status       string `json:"status"`
	DataInicio   string `json:"data_inicio"`
	DataFim      string `json:"data_fim"`
}
