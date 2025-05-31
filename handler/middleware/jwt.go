package middleware

import (
	"log"

	"github.com/abiyyu03/auth-service/constant/message"
	"github.com/abiyyu03/auth-service/entity/dto"
	"github.com/abiyyu03/auth-service/service"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v5"
)

func HandleJWTMiddleware(roleService *service.RoleService, allowedRoles []int) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningMethod: "RS256",
		SigningKey:    "SecretBangetNih",
		ContextKey:    "user",
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return dto.ResponseFailedStruct(ctx, message.ErrUnauthorized.Error(), fiber.ErrUnauthorized.Code)
		},
		SuccessHandler: func(ctx *fiber.Ctx) error {
			// Extract user claims from the token
			log.Print("Allowed Role ID: ", 1)
			user := ctx.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)

			// Extract the role_id from the JWT claims
			roleID := claims["role_id"].(float64)
			castedRoleID := int(roleID)

			// Fetch the role by ID using the service
			allowedRoleId, err := roleService.FindById(castedRoleID)

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
