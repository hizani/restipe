package ginhandler

import (
	"net/http"
	"restipe/internal/service"

	"github.com/gin-gonic/gin"
)

const (
	userCtx = "userId"
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
		recipes := api.Group("/recipes")
		{
			recipes.GET("/", h.getAllRecipes)
			recipes.GET("/:id", h.getRecipeById)

			ingredients := recipes.Group(":id/ingredients")
			{
				ingredients.GET("/", h.getAllIngredientsFromRecipe)
			}

			steps := recipes.Group(":id/steps")
			{
				steps.GET("/", h.getAllStepsFromRecipe)
			}

			auth := recipes.Group("/", h.Authorize)
			{
				auth.POST("/", h.createRecipe)
				auth.PUT("/:id", h.updateRecipe)
				recipes.DELETE("/:id", h.deleteRecipe)

				ingredients := auth.Group(":id/ingredients")
				{
					ingredients.POST("/", h.addIngredientToRecipe)
					ingredients.DELETE("/:ingid", h.removeIngredientFromRecipe)
				}

				steps := auth.Group(":id/steps")
				{
					steps.POST("/", h.addStepToRecipe)
					steps.DELETE("/:stepid", h.removeStepFromRecipe)
				}
			}

		}

	}
	return h
}
