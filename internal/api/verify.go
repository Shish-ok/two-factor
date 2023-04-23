package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"two-factor-auth/internal/models/confirmations"
	"two-factor-auth/internal/services/auth"
)

type VerifyCodeResponse struct {
	UnixDate confirmations.CurrentUnixDate `json:"verifiedAt"`
}

// VerifyCode godoc
// @Summary Верифицирует код двухфакторной авторизации
// @Schemes
// @Description Верифицирует код двухфакторной авторизации, возвращая текущую дату в UnixTime
// @Description Код ошибки 400: неверный json
// @Description Код ошибки 403: исчерпан лимит попыток или неверный код
// @Tags auth
// @Accept json
// @Produce json
// @Param data body confirmations.Confirmation true "Входные параметры"
// @Success 200 {object} VerifyCodeResponse
// @Failure 400 {object} Error
// @Failure 403 {object} Error
// @Failure 500
// @Router /verify [post]
func (api *API) VerifyCode(ctx *gin.Context) {
	var requestConfirmation confirmations.Confirmation
	if err := ctx.BindJSON(&requestConfirmation); err != nil {
		ctx.JSON(http.StatusBadRequest, Error{Error: err.Error()})
		return
	}

	unixDate, err := api.auth.Verify(ctx, requestConfirmation)
	if err == auth.ErrToManyAttempts {
		ctx.JSON(http.StatusForbidden, Error{Error: err.Error()})
		return
	}
	if err == auth.ErrWrongCode {
		ctx.JSON(http.StatusForbidden, Error{Error: err.Error()})
	}
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, VerifyCodeResponse{UnixDate: unixDate})
}
