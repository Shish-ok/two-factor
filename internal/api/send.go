package api

import (
	"github.com/gin-gonic/gin"
	"two-factor-auth/internal/models/confirmations"
)

type SendCodeRequest struct {
	UserNumber confirmations.Number `json:"number"`
}

// SendCode godoc
// @Summary Отправляет код для двухфакторной авторизации
// @Schemes
// @Description Отправляет код двухфакторной авторизации, возвращая id запроса и сам код
// @Description Код ошибки 400: неверный json или невалидный номер
// @Tags auth
// @Accept json
// @Produce json
// @Param data body SendCodeRequest true "Входные параметры"
// @Success 200 {object} confirmations.Confirmation
// @Failure 400 {object} Error
// @Router /send [post]
func (api *API) SendCode(ctx *gin.Context) {

}
