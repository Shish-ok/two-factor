package api

import (
	"github.com/gin-gonic/gin"
	"two-factor-auth/internal/models/confirmations"
)

type VerifyCodeResponse struct {
	UnixDate confirmations.CurrentUnixDate `json:"verifiedAt"`
}

// VerifyCode godoc
// @Summary Верифицирует код двухфакторной авторизации
// @Schemes
// @Description Верифицирует код двухфакторной авторизации, возвращая текущую дату в UnixTime
// @Description Код ошибки 400: неверный json или невалидный код авторизации
// @Description Код ошибки 403: исчерпан лимит попыток или код уже не действителен
// @Tags auth
// @Accept json
// @Produce json
// @Param data body confirmations.Confirmation true "Входные параметры"
// @Success 200 {object} VerifyCodeResponse
// @Failure 400 {object} Error
// @Failure 403 {object} Error
// @Router /verify [post]
func (api *API) VerifyCode(ctx *gin.Context) {

}
