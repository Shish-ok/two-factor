package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"two-factor-auth/internal/models/confirmations"
)

type SendCodeRequest struct {
	UserNumber confirmations.Number `json:"number"`
}

// SendCode godoc
// @Summary Отправляет код для двухфакторной авторизации
// @Schemes
// @Description Отправляет код двухфакторной авторизации, возвращая id запроса и сам код
// @Description Код ошибки 400: неверный json
// @Tags auth
// @Accept json
// @Produce json
// @Param data body SendCodeRequest true "Входные параметры"
// @Success 200 {object} confirmations.Confirmation
// @Failure 400 {object} Error
// @Failure 500
// @Router /send [post]
func (api *API) SendCode(ctx *gin.Context) {
	var request SendCodeRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, Error{Error: err.Error()})
		return
	}

	confirmation, err := api.auth.SendCode(ctx, request.UserNumber)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, confirmation)
}
