package api

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"two-factor-auth/docs"
	"two-factor-auth/internal/config"
	"two-factor-auth/internal/services/auth"
)

func NewAPI(
	router *gin.Engine,
	cfg config.ServiceConfiguration,
	auth *auth.Service,
) *API {
	api := &API{
		router:  router,
		host:    cfg.APIConfig.GetAddr(),
		useCORS: cfg.APIConfig.UseCORS,
		auth:    auth,
	}

	api.base = api.router.Group(BasePath)
	return api
}

type API struct {
	router  *gin.Engine
	useCORS bool
	host    string
	base    *gin.RouterGroup
	auth    *auth.Service
}

// @BasePath /api/v1/
const BasePath = "/api/v1"

func (api *API) Run() {
	docs.SwaggerInfo.BasePath = BasePath
	api.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))

	if api.useCORS {
		api.router.Use(CORS())
	}
	
	api.registerRoutes()

	api.router.Run(api.host)
}

type Error struct {
	Error string `json:"error"`
}

func NewError(err string) Error {
	return Error{Error: err}
}
