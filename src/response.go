package src

type BaseResponse struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
}

type BusinessResponse struct {
	Result  bool     `json:"result"`
	Message string   `json:"message"`
	Data    business `json:"data"`
}

type KeksResponse struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
	Data    []kek  `json:"data"`
}
