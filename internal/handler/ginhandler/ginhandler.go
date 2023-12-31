package ginhandler

import (
	"net/http"
	"restipe/internal/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docs "restipe/docs"
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

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
				images := steps.Group("/:number/images")
				{
					images.GET("/", h.downloadStepImg)

				}
			}

			images := recipes.Group(":id/images")
			{
				images.GET("/", h.downloadRecipeImg)
			}

			auth := recipes.Group("/", h.Authorize)
			{
				auth.POST("/", h.createRecipe)
				auth.PUT("/:id", h.updateRecipe)
				auth.DELETE("/:id", h.deleteRecipe)

				ingredients := auth.Group(":id/ingredients")
				{
					ingredients.POST("/", h.addIngredientToRecipe)
					ingredients.DELETE("/:ingid", h.removeIngredientFromRecipe)
				}

				steps := auth.Group(":id/steps")
				{
					steps.POST("/", h.addStepToRecipe)
					steps.DELETE("/:number", h.removeStepFromRecipe)
					images := steps.Group("/:number/images")
					{
						images.POST("/", h.uploadStepImg)

					}
				}
				rates := auth.Group(":id/rates")
				{
					rates.POST("/", h.RateRecipe)
					rates.PUT("/", h.RerateRecipe)
				}

				images := auth.Group(":id/images")
				{
					images.POST("/", h.uploadRecipeImg)
				}

			}

		}

	}
	return h
}
