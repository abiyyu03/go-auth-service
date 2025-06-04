package middleware

import (
	"log"
	"strings"

	"github.com/abiyyu03/auth-service/constant/message"
	"github.com/abiyyu03/auth-service/entity/dto"
	"github.com/abiyyu03/auth-service/service"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v5"
)

func HandleJWTMiddleware(roleService *service.RoleService, allowedRoles []int) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningMethod: "HS256",
		SigningKey:    []byte("SecretBangetNih"),
		ContextKey:    "user",
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			log.Print("JWT Error: ", err.Error())
			return dto.ResponseFailedStruct(ctx, message.ErrUnauthorized.Error(), fiber.ErrUnauthorized.Code)
		},
		SuccessHandler: func(ctx *fiber.Ctx) error {
			// Ambil token string dari Authorization header
			authHeader := ctx.Get("Authorization")
			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

			// Parse token ulang pakai jwt/v5
			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fiber.ErrUnauthorized
				}
				return []byte("SecretBangetNih"), nil
			})
			if err != nil || !token.Valid {
				return fiber.ErrUnauthorized
			}

			_, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return fiber.ErrUnauthorized
			}
			castedRoleID := token.Claims.(jwt.MapClaims)["role_id"].(float64)

			// Fetch the role by ID using the service
			allowedRoleId, err := roleService.FindById(int(castedRoleID))

			if err != nil {
				return err
			}

			if !HasRequiredRole(allowedRoleId.ID, allowedRoles) {
				return dto.ResponseSuccessStruct(ctx, message.ErrInvalidID.Error(), fiber.ErrForbidden.Code, nil)
			}

			return ctx.Next()
		},
	})
}

func HasRequiredRole(userRoles int, allowedRoles []int) bool {
	for _, allowed := range allowedRoles {
		if userRoles == allowed {
			return true
		}
	}
	return false
}
