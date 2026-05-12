package function

import (
	"database/sql"
	"strings"
)

type Unidade struct {
	ID          string `json:"id"`
	Numero      *string `json:"numero,omitempty"`
    Apartamento *string `json:"apartamento,omitempty"`
    Bloco       *string `json:"bloco,omitempty"`
    Ramal       *string `json:"ramal,omitempty"`
}

func ListUnits(db *sql.DB, idCondominio string, page, limit int, filters map[string]interface{}) (map[string]interface{}, error) {
	offset := (page - 1) * limit

	query := `SELECT unidade.id, unidade.numero, unidade.apartamento, unidade.bloco, unidade.ramal FROM unidade WHERE unidade.id_condominio = ?`
	countQuery := `SELECT COUNT(*) FROM unidade WHERE id_condominio = ?`
	
	countArguments := []interface{}{idCondominio}
	arguments := []interface{}{idCondominio}

	if filters != nil {
		if numero, ok := filters["numero"]; ok {
            query += " AND unidade.numero = ?"
            countQuery += " AND unidade.numero = ?"
            arguments = append(arguments, numero)
            countArguments = append(countArguments, numero)
        }
        if unidadeSemCondomino, ok := filters["unidade_sem_condomino"]; ok {
            if isUnidadeSemCondomino, isBool := unidadeSemCondomino.(bool); isBool && isUnidadeSemCondomino {
                query = strings.Replace(query, "FROM unidade", "FROM unidade LEFT JOIN condomino_unidade ON unidade.id = condomino_unidade.id_unidade", 1)
                query += " AND condomino_unidade.id_unidade IS NULL"
                countQuery = strings.Replace(countQuery, "FROM unidade", "FROM unidade LEFT JOIN condomino_unidade ON unidade.id = condomino_unidade.id_unidade", 1)
                countQuery += " AND condomino_unidade.id_unidade IS NULL"
            }
        }
	}

	query += " LIMIT ? OFFSET ?;"
	arguments = append(arguments, limit, offset)

	rows, err := db.Query(query, arguments...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var unidades []Unidade
	for rows.Next() {
    	var u Unidade
    	var numero, apartamento, bloco, ramal sql.NullString
    
    	err := rows.Scan(
        	&u.ID,
	        &numero,
    	    &apartamento,
        	&bloco,
        	&ramal,
    	)
    	if err != nil {
        	return nil, err
    	}
	    if numero.Valid {
        	u.Numero = &numero.String
    	}
	    if apartamento.Valid {
        	u.Apartamento = &apartamento.String
    	}
	    if bloco.Valid {
        	u.Bloco = &bloco.String
    	}
	    if ramal.Valid {
        	u.Ramal = &ramal.String
    	}
    	unidades = append(unidades, u)
	}

	var total int
	err = db.QueryRow(countQuery, countArguments...).Scan(&total)
	if err != nil {
		return nil, err
	}
	result := map[string]interface{}{
		"total": total,
		"page":  page,
		"limit": limit,
		"data":  unidades,
	}
	
	return result, nil
}
