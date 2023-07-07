package ginhandler

import (
	"net/http"
	"restipe/internal/service"

	"github.com/gin-gonic/gin"
)

type GinHandler struct {
	http.Handler
	service *service.Service
}

func New(service *service.Service) *GinHandler {
	router := gin.New()
	h := &GinHandler{router, service}

	auth := router.Group("/auth")
	{
		auth.POST("/signup", h.signUp)
		auth.POST("/signin", h.signIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/recipes")
		{
			lists.POST("/", h.createRecipe)
			lists.GET("/", h.getAllRecipe)
			lists.GET("/:id", h.getRecipeById)
			lists.PUT("/:id", h.updateRecipe)
			lists.DELETE("/:id", h.deleteRecipe)

			ingredients := lists.Group(":id/ingredients")
			{
				ingredients.POST("/", h.createIngredient)
				ingredients.GET("/", h.getAllIngredients)
				ingredients.DELETE("/:id", h.deleteIngredient)
			}

			steps := lists.Group(":id/steps")
			{
				steps.POST("/", h.addStep)
				steps.GET("/", h.getAllSteps)
				steps.DELETE("/:id", h.deleteStep)
			}
		}
	}
	return h
}
