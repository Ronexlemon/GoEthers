package responses


type MessageResponse struct{
	Message string `json:"message"`
	Status int  `json:"status"`
	Data  map[string]interface{} `json:"data"`
}

