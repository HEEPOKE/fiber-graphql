package middleware

import (
	"crypto/sha256"
	"crypto/subtle"
	"errors"

	"github.com/HEEPOKE/fiber-graphql/internals/domains/models/response"
	"github.com/HEEPOKE/fiber-graphql/pkg/configs"
	"github.com/HEEPOKE/fiber-graphql/pkg/constants"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
)

var ErrMissingOrMalformedAPIKey = errors.New("missing or malformed API Key")

var ConfigDefault = keyauth.Config{
	SuccessHandler: func(c *fiber.Ctx) error {
		return c.Next()
	},
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		statusMessage := response.StatusMessage{
			Code:        constants.CODE_INVALID,
			Message:     constants.MESSAGE_INVALID,
			Service:     constants.KEY_AUTH_SERVICE,
			Description: constants.AUTH_API_KEY_FAILED,
		}

		response := response.ResponseMessage{
			Status:  statusMessage,
			Payload: nil,
		}

		return c.Status(fiber.StatusUnauthorized).JSON(response)
	},
	KeyLookup:  "header:api_key",
	AuthScheme: "",
	Validator: func(c *fiber.Ctx, key string) (bool, error) {
		hashedAPIKey := sha256.Sum256([]byte(configs.Cfg.API_KEY))
		hashedKey := sha256.Sum256([]byte(key))

		if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) == 1 {
			return true, nil
		}
		return false, ErrMissingOrMalformedAPIKey
	},
}
