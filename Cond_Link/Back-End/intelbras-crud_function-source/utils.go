package function

// This join two maps
func mergeMaps(map1, map2 map[string]interface{}) map[string]interface{} {
	// Cria um novo mapa para armazenar o resultado
	mergedMap := make(map[string]interface{})

	// Adiciona todos os elementos do primeiro mapa
	for key, value := range map1 {
		mergedMap[key] = value
	}

	// Adiciona todos os elementos do segundo mapa
	for key, value := range map2 {
		mergedMap[key] = value
	}

	return mergedMap
}
