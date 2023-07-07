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

	api := router.Group("/api", h.Authorize)
	{
		recipes := api.Group("/recipes")
		{
			recipes.POST("/", h.createRecipe)
			recipes.GET("/", h.getAllRecipe)
			recipes.GET("/:id", h.getRecipeById)
			recipes.PUT("/:id", h.updateRecipe)
			recipes.DELETE("/:id", h.deleteRecipe)

			ingredients := recipes.Group(":id/ingredients")
			{
				ingredients.POST("/", h.createIngredient)
				ingredients.GET("/", h.getAllIngredients)
				ingredients.DELETE("/:id", h.deleteIngredient)
			}

			steps := recipes.Group(":id/steps")
			{
				steps.POST("/", h.addStep)
				steps.GET("/", h.getAllSteps)
				steps.DELETE("/:id", h.deleteStep)
			}
		}
	}
	return h
}