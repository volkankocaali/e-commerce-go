package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	constant "github.com/volkankocaali/e-commorce-go/pkg/constant/middleware"
	"net/http"
	"strings"
)

func UserAuthMiddleware(ctx *fiber.Ctx) error {
	// use fiber
	tokenString := ctx.Get("Authorization")

	if tokenString == "" {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": constant.MissingAuthToken})
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("ecommercegoapplication"), nil
	})

	if err != nil || !token.Valid {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": constant.InvalidToken})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": constant.InvalidToken})
	}

	fmt.Println("claims", claims)

	role, ok := claims["role"].(string)
	if !ok || role != "client" {
		return ctx.Status(http.StatusForbidden).JSON(fiber.Map{"error": constant.Unauthorized})
	}

	id, ok := claims["id"].(float64)
	if !ok || id == 0 {
		return ctx.Status(http.StatusForbidden).JSON(fiber.Map{"error": constant.ErrorInRetrievingId})
	}

	ctx.Locals("role", role)
	ctx.Locals("id", int(id))

	return ctx.Next()
}
