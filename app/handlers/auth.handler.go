package handlers

import (
	"base-fiber/app/utils"
	"base-fiber/app/utils/rules"
	"base-fiber/src/role"
	"base-fiber/src/user"
	"base-fiber/src/verification"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	userService user.Service
	roleService role.Service
	verificationService verification.Service
}

func NewAuthHandler(userService user.Service, roleService role.Service,verificationService verification.Service) *authHandler {
	return &authHandler{userService,roleService, verificationService}
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

	code := utils.RandCode(6,"number")

	// create verification
	verification := verification.InputVerification{
		UserId: int(newUser.ID),
		Types: "verification",
		Code: code,
	}

	_, errVerif := h.verificationService.Save(verification)
	if errVerif != nil {
		response := utils.ApiRespone(errVerif.Error(), http.StatusUnprocessableEntity, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	utils.Verification(newUser.Name, newUser.Email, code, "sigitbudianto423@gmail.com")

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

func (h *authHandler) Verification(c *fiber.Ctx) error {
	bodyUser := new(user.InputVerification)
	if err := c.BodyParser(bodyUser); err != nil{ 
		response := utils.ApiRespone("Invalid format Request", http.StatusUnprocessableEntity, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	errValidate := user.ValidateVerification(*bodyUser)
	if errValidate != nil {
		response := utils.ApiRespone("Verification failed", http.StatusBadRequest, "error", errValidate)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	newUser, errFindEmail := h.userService.FindEmail(bodyUser.Email)
	if errFindEmail != nil {
		response := utils.ApiRespone(errFindEmail.Error(), http.StatusBadRequest, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	verif, errVerif :=h.verificationService.Find(bodyUser.Code,"verification", int(newUser.ID))
	if errVerif != nil {
		utils.Log(errVerif.Error(), "err", "service")
		response := utils.ApiRespone(errVerif.Error(), http.StatusUnprocessableEntity, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if verif.Expired.Unix() <= time.Now().Unix(){
		errUpdateStatus := h.verificationService.UpdateStatus(int(verif.ID))
		if errUpdateStatus != nil {
			utils.Log(errUpdateStatus.Error(), "err", "service")
		}
		response := utils.ApiRespone("Expired Code", http.StatusUnprocessableEntity, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	errUpdateUser := h.userService.Verification(user.InputVerification{
		Email: bodyUser.Email,
		Code: verif.Code,
	})

	if errUpdateUser != nil {
		utils.Log(errUpdateUser.Error(), "err", "service")
		response := utils.ApiRespone(errUpdateUser.Error(), http.StatusUnprocessableEntity, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	errUpdateVerification := h.verificationService.UpdateStatus(int(verif.ID))
	if errUpdateVerification != nil {
		utils.Log(errUpdateVerification.Error(), "err", "service")
		response := utils.ApiRespone(errUpdateVerification.Error(), http.StatusUnprocessableEntity, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := utils.ApiRespone("Verification Successfully", http.StatusOK, "success", nil)

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *authHandler) SendOtp(c *fiber.Ctx) error {
	bodyUser := new(user.InputSendOtp)
	if err := c.BodyParser(bodyUser); err != nil{ 
		response := utils.ApiRespone("Invalid format Request", http.StatusUnprocessableEntity, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	
	errValidate := user.ValidateEmail(*bodyUser)
	if errValidate != nil {
		response := utils.ApiRespone("Login failed", http.StatusBadRequest, "error", errValidate)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	
	newUser, errFindEmail := h.userService.FindEmail(bodyUser.Email)
	if errFindEmail != nil {
		response := utils.ApiRespone(errFindEmail.Error(), http.StatusBadRequest, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	
	if bodyUser.Types == "forgot-password"{
		if newUser.Status != 1 {
			response := utils.ApiRespone("Your account has not been verified", http.StatusBadRequest, "error", nil)
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}
	}else{
		if newUser.Status == 1 {
			response := utils.ApiRespone("Your account has been verified", http.StatusBadRequest, "error", nil)
			return c.Status(fiber.StatusBadRequest).JSON(response)	
		}
	}
	var code string
	if bodyUser.Types == "forgot-password" {
		code = utils.RandCode(20,"string")
	}else{	
		code = utils.RandCode(6,"number")
	}

	// create verification
	verification := verification.InputVerification{
		UserId: int(newUser.ID),
		Types: bodyUser.Types,
		Code: code,
	}

	_, errVerif := h.verificationService.Save(verification)
	if errVerif != nil {
		response := utils.ApiRespone(errVerif.Error(), http.StatusUnprocessableEntity, "error", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	if bodyUser.Types == "forgot-password"{
		utils.ForgotPassword(code, newUser.Email, newUser.Name)
	}else{
		utils.Verification(newUser.Name, newUser.Email, code, "sigitbudianto423@gmail.com")
	}
	response := utils.ApiRespone("Send Otp/forgot password Successfully", http.StatusOK, "success", nil)

	return c.Status(fiber.StatusOK).JSON(response)
}