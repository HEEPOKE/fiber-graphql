package examples

type FailedCommonResponse struct {
	Status  FailedStatusMessage `json:"status"`
	Payload interface{}         `json:"payload"`
}
