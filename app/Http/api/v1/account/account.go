package account

import (
	"test/app/repositories"
	"test/app/util"

	"github.com/gofiber/fiber/v2"
)

// ShowAccount godoc
// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} Account
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /api/v1/accounts/{id} [get]
func ShowAccount(c *fiber.Ctx) error {
	return c.JSON(util.Success(Account{
		Id: c.Params("id"),
	}))

}

func Login(c *fiber.Ctx) error {
	payload := struct {
		Account  string `json:"account"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	tokenString, err := repositories.Login(payload.Account, payload.Password)

	if err != nil {
		return c.JSON(util.Fail())
	}
	return c.JSON(util.Success(tokenString))
}

type Account struct {
	Id string
}

type HTTPError struct {
	status  string
	message string
}
