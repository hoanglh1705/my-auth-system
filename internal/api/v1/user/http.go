package user

import (
	"my-auth-service/internal/usecase/authentusecase"

	"github.com/labstack/echo/v4"
)

// HTTP represents auth http service
type HTTP struct {
	svc authentusecase.AuthentUsecase
}

// NewHTTP creates new auth http service
func NewHTTP(svc authentusecase.AuthentUsecase, eg *echo.Group) {
	h := HTTP{svc}

	// swagger:operation POST /v1/customer/login auth-customer authCustomerLogin
	// ---
	// summary: Login customer by username and password
	// parameters:
	// - name: request
	//   in: body
	//   description: Request body
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/CredentialUserData"
	// responses:
	//   "200":
	//     description: Access token
	//     schema:
	//       "$ref": "#/definitions/AuthToken"
	//   "401":
	//     "$ref": "#/responses/errDetails"
	//   "500":
	//     "$ref": "#/responses/errDetails"
	eg.POST("/login", h.login)
}

func (h *HTTP) login(e echo.Context) error {
	
	return h.svc.Login(e.Request().Context(), )
}
