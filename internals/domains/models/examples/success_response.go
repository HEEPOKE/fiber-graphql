package examples

import (
	"github.com/HEEPOKE/fiber-graphql/internals/domains/models"
	"github.com/HEEPOKE/fiber-graphql/internals/domains/models/response"
)

type SuccessLoginResponse struct {
	Status  SuccessStatusMessage   `json:"status"`
	PayLoad response.LoginResponse `json:"payload"`
}

type SuccessRegisterAccountResponse struct {
	Status  SuccessStatusMessage `json:"status"`
	PayLoad string               `json:"payload"`
}

type SuccessAccountsGetAllResponse struct {
	Status  SuccessStatusMessage  `json:"status"`
	PayLoad []models.AccountModel `json:"payload"`
}

type SuccessAccountsProfileResponse struct {
	Status  SuccessStatusMessage          `json:"status"`
	PayLoad response.AccountResponseModel `json:"payload"`
}
