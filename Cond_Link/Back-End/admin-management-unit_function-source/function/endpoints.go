package function

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"fmt"
)

/*
Handles with /v1/test/
*/
func handleControlViewRoutes(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	switch {
	case strings.HasSuffix(r.URL.Path, "/v1/unidades/listar") && r.Method == "GET":
		queryParams := r.URL.Query()
		idCondominio := queryParams.Get("id_condominio")
		page := queryParams.Get("page")
		limit := queryParams.Get("limit")
		filters := queryParams.Get("filters")

		if idCondominio == "" {
			returnMessage(w, "id_condominio é obrigatório", http.StatusBadRequest)
			return
		}

		pageInt, err := strconv.Atoi(page)
		if err != nil || pageInt < 1 {
			pageInt = 1
		}

		limitInt, err := strconv.Atoi(limit)
		if err != nil || limitInt < 1 {
			limitInt = 5
		}

		var filterMap map[string]interface{}
		if filters != "" {
			if err := json.Unmarshal([]byte(filters), &filterMap); err != nil {
				returnMessage(w, "Formato de filters inválido", http.StatusBadRequest)
				return
			}
			for field := range filterMap {
        		if field != "numero" && field != "unidade_sem_condomino" {
            		returnMessage(w, fmt.Sprintf("Campo de filtro inválido: %s", field), http.StatusBadRequest)
            		return
        		}
    		}
		}
		result, err := ListUnits(db, idCondominio, pageInt, limitInt, filterMap)
		if err != nil {
			returnMessage(w, "Erro ao buscar unidade: "+err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(result)

		return
	}

}
