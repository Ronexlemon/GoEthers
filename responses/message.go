package responses


type MessageResponse struct{
	Message string `json:"message"`
	Status Int  `json:"status"`
	Data  map[string]interface{} `json:"data"`
}