package handlers

import (
	"base-fiber/app/utils"
	"base-fiber/app/utils/rules"
	"base-fiber/src/role"
	"base-fiber/src/user"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	userService user.Service
	roleService role.Service
}

func NewAuthHandler(userService user.Service, roleService role.Service) *authHandler {
	return &authHandler{userService,roleService}
}

func (h *authHandler) RegisterUser(c *fiber.Ctx) error {
	bodyUser := new(user.InputRegister)
	if err := c.BodyParser(bodyUser); err != nil{ 
		response := utils.ApiRespone("Invalid format Request", http.StatusUnprocessableEntity, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	
	errValidate := user.ValidateRegister(*bodyUser)
	if errValidate != nil {
		response := utils.ApiRespone("Register failed", http.StatusBadRequest, "error", errValidate)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if !rules.Password(bodyUser.Password) {
		response := utils.ApiRespone("The password must contain a combination of numbers, uppercase letters and symbols .", http.StatusBadRequest, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	_, errRole := h.roleService.FindById(int(bodyUser.RoleId))
	if errRole != nil {
		response := utils.ApiRespone(errRole.Error(), http.StatusUnprocessableEntity, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	newUser ,err := h.userService.RegisterUser(*bodyUser)	
	if err != nil {
		response := utils.ApiRespone(errRole.Error(), http.StatusUnprocessableEntity, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	formatter := user.FormatUser(newUser)
	response := utils.ApiRespone("Account has been regitered", http.StatusOK, "success", formatter)

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *authHandler) Login(c *fiber.Ctx) error {
	bodyUser := new(user.InputLogin)
	if err := c.BodyParser(bodyUser); err != nil{ 
		response := utils.ApiRespone("Invalid format Request", http.StatusUnprocessableEntity, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	
	errValidate := user.ValidateLogin(*bodyUser)
	if errValidate != nil {
		response := utils.ApiRespone("Login failed", http.StatusBadRequest, "error", errValidate)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	token, errService := h.userService.Login(*bodyUser)
	if errService != nil {
		utils.Log(errService.Error(), "err", "service")
		response := utils.ApiRespone("credential is wrong", http.StatusBadRequest, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	formatter := user.FormaterLogin(user.LoginFormatter{
		Token: token,
	})
	response := utils.ApiRespone("Login Successfully", http.StatusOK, "success", formatter)

	return c.Status(fiber.StatusOK).JSON(response)
}