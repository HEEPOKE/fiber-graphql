package middleware

import (
	"github.com/HEEPOKE/fiber-graphql/internals/domains/models/response"
	"github.com/HEEPOKE/fiber-graphql/pkg/configs"
	"github.com/HEEPOKE/fiber-graphql/pkg/constants"
	jwtWare "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func JwtMiddleware() fiber.Handler {
	return jwtWare.New(jwtWare.Config{
		SigningKey: jwtWare.SigningKey{Key: []byte(configs.Cfg.SECRET_KEY)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			statusMessage := response.StatusMessage{
				Code:        constants.CODE_INVALID,
				Message:     constants.MESSAGE_INVALID,
				Service:     constants.JWT_SERVICE,
				Description: constants.TOKEN_FAILED,
			}

			response := response.ResponseMessage{
				Status:  statusMessage,
				Payload: nil,
			}

			return c.Status(fiber.StatusUnauthorized).JSON(response)
		},
	})
}
