package examples

type SuccessStatusMessage struct {
	Code        string `json:"code" example:"0000"`
	Message     string `json:"message" example:"Success"`
	Service     string `json:"service"`
	Description string `json:"description"`
}

type FailedStatusMessage struct {
	Code        string `json:"code" example:"0001"`
	Message     string `json:"message" example:"Failed"`
	Service     string `json:"service"`
	Description string `json:"description"`
	Error       string `json:"error" example:"error message"`
}

type NotFoundStatusMessage struct {
	Code        string `json:"code" example:"0002"`
	Message     string `json:"message" example:"Not Found"`
	Service     string `json:"service"`
	Description string `json:"description"`
}

type DuplicateStatusMessage struct {
	Code        string `json:"code" example:"0003"`
	Message     string `json:"message" example:"Duplicate"`
	Service     string `json:"service"`
	Description string `json:"description"`
}

type InvalidStatusMessage struct {
	Code        string `json:"code" example:"0004"`
	Message     string `json:"message" example:"Invalid"`
	Service     string `json:"service"`
	Description string `json:"description"`
}
