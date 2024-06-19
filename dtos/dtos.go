package dtos

type Service struct {
	ID      uint   `params:"id"`
	Name    string `json:"name" form:"name"`
	Command string `json:"command" form:"command"`
}

type ServiceResponse struct {
	ID            uint   `json:"id" form:"id,omitempty"`
	Name          string `json:"name" form:"name"`
	Command       string `json:"command" form:"command"`
	LastCheckTime int64  `json:"lastCheckTime"`
	Status        int    `json:"status" form:"status"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
