package middleware

import (
	"base-fiber/app/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Authenticate(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	if token == "" {
		blankToken := utils.ApiRespone("unauthenticated", fiber.StatusUnauthorized, "error", nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(blankToken)
	}

	if !strings.Contains(token, "Bearer"){
		inValidToken := utils.ApiRespone("Invalid format token", fiber.StatusUnauthorized, "error", nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(inValidToken)
	}
	
	tokenString := ""
	arrayToken := strings.Split(token, " ")
	if len(arrayToken) == 2{
		tokenString = arrayToken[1]
	}

	claims, err := utils.DecodeToken(tokenString)
	if err != nil {
		utils.Log(err.Error(),"err","utils")
		inValidDecodeToken := utils.ApiRespone(err.Error(), fiber.StatusUnauthorized, "error", nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(inValidDecodeToken)
	}

	ctx.Locals("email", claims["email"].(string))
	return ctx.Next()
}