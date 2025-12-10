package middleware

import (
	"rapnews/config"
	"rapnews/internal/adapter/handler/response"
	"rapnews/lib/auth"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Middleware interface {
	CheckToken() fiber.Handler
}

type Options struct {
	authJwt auth.Jwt
}

// CheckToken implements Middleware.
func (o *Options) CheckToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var errorResponse response.ErrorResponseDefault

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			errorResponse.Meta.Status = false
			errorResponse.Meta.Message = "Missing Authorization header"
			return c.Status(fiber.StatusUnauthorized).JSON(errorResponse)
		}

		// Harus menggunakan format: "Bearer <token>"
		parts := strings.Fields(authHeader)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			errorResponse.Meta.Status = false
			errorResponse.Meta.Message = "Invalid Authorization header format"
			return c.Status(fiber.StatusUnauthorized).JSON(errorResponse)
		}

		tokenString := parts[1]
		claims, err := o.authJwt.VerifyAccessToken(tokenString)
		if err != nil {
			errorResponse.Meta.Status = false
			errorResponse.Meta.Message = "Invalid or expired token"
			return c.Status(fiber.StatusUnauthorized).JSON(errorResponse)
		}

		// Simpan claims ke context
		c.Locals("user", claims)
		return c.Next()
	}
}

func NewMiddleware(cfg *config.Config) Middleware {
	return &Options{
		authJwt: auth.NewJwt(cfg),
	}
}
