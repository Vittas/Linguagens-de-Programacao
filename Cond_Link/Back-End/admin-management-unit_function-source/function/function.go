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

// BootUp é o ponto de entrada da Cloud Function
func BootUp(w http.ResponseWriter, r *http.Request) {
	// CORS e content-type
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-GUARDIAN-KEY")

	// Se for uma requisição OPTIONS, retorna direto (sem autenticar!)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent) // 204 No Content
		return
	}

	// Verifica o token
	token := r.Header.Get("X-GUARDIAN-KEY")
	if token != "12DE-76XD-98HJ-00PO" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Token inválido"})
		return
	}

	// Conexão com banco
	db, err := connectToDatabase()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Erro ao conectar ao banco: " + err.Error()})
		return
	}
	defer db.Close()

	// Roteamento
	path := r.URL.Path

	if strings.HasPrefix(path, "/v1/unidades/listar") {
		handleControlViewRoutes(w, r, db)
		return
	}

	// 404 se não achou rota
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"message": "Rota não encontrada"})
}

// Function to connect to database
func connectToDatabase() (*sql.DB, error) {

	var (
		dbUser    = "joao"
		dbPwd     = "Ph=f-rF9A{8285&+"
		dbName    = "multicondlink"
		dbPort    = "3306"
		// dbTCPHost = "10.129.0.20"
		dbTCPHost = "mysql1.condlink.com.br"
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
