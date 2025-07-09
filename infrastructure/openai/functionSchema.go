package openai

func GetColmeiaFunctionSchemas() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"name":        "list_colmeia",
			"description": "List all bee hives",
			"parameters": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"status": map[string]interface{}{
						"type":        "string",
						"description": "Status of bee hive",
					},
					"species": map[string]interface{}{
						"type":        "string",
						"description": "Species of bee hive",
					},
				},
			},
		},
		{
			"name":        "add_colmeia",
			"description": "Add a new bee hive",
			"parameters": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"colmeia_id": map[string]interface{}{
						"type":        "int",
						"description": "ID of bee hive",
					},
					"species": map[string]interface{}{
						"type":        "string",
						"description": "Species of bee hive",
					},
					"starting_date": map[string]interface{}{
						"type":        "string",
						"description": "Starting date of bee hive",
					},
					"status": map[string]interface{}{
						"type":        "string",
						"description": "Status of bee hive",
					},
				},
			},
			"required": []string{"species"},
		},
	}
}
ackage openai