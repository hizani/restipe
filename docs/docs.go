// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/recipes": {
            "get": {
                "description": "get all recipes",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipe"
                ],
                "summary": "get all recipes",
                "operationId": "get-all-recipes",
                "parameters": [
                    {
                        "description": "recipe search request info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.GetAllRecipesReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "create recipe",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipe"
                ],
                "summary": "create recipe",
                "operationId": "create-recipe",
                "parameters": [
                    {
                        "description": "recipe info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateRecipeReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    }
                }
            }
        },
        "/api/recipes/{id}": {
            "get": {
                "description": "get recipe by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipe"
                ],
                "summary": "get recipe by id",
                "operationId": "get-recipe-by-id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "update recipe",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipe"
                ],
                "summary": "update recipe",
                "operationId": "update-recipe",
                "parameters": [
                    {
                        "description": "recipe update info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateRecipeReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "delete recipe",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipe"
                ],
                "summary": "delete recipe",
                "operationId": "delete-recipe",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    }
                }
            }
        },
        "/api/recipes/{id}/ingredients": {
            "get": {
                "description": "get all recipe ingredients",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipe ingredient"
                ],
                "summary": "get ingredients",
                "operationId": "get-ingredients",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "add ingredient to recipe",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipe ingredient"
                ],
                "summary": "add ingredient",
                "operationId": "add-ingredient",
                "parameters": [
                    {
                        "description": "ingredient info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AddIngredientReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    }
                }
            }
        },
        "/api/recipes/{id}/ingredients/{ingid}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "remove ingredient from recipe",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipe ingredient"
                ],
                "summary": "remove ingredient",
                "operationId": "remove-ingredient",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    }
                }
            }
        },
        "/api/recipes/{id}/ingredients/{stepid}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "remove step from recipe",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipe step"
                ],
                "summary": "remove step",
                "operationId": "remove-step",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    }
                }
            }
        },
        "/api/recipes/{id}/step": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "add step to recipe",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipe step"
                ],
                "summary": "add step",
                "operationId": "add-step",
                "parameters": [
                    {
                        "description": "ingredient info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AddStepReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    }
                }
            }
        },
        "/api/recipes/{id}/steps": {
            "get": {
                "description": "get all recipe steps",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipe step"
                ],
                "summary": "get steps",
                "operationId": "get-steps",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    }
                }
            }
        },
        "/auth/signin": {
            "post": {
                "description": "log into user account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "SingIn",
                "operationId": "login",
                "parameters": [
                    {
                        "description": "account info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SigninUserReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    }
                }
            }
        },
        "/auth/signup": {
            "post": {
                "description": "create account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "SingUp",
                "operationId": "create-account",
                "parameters": [
                    {
                        "description": "account info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SignupUserReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/ginhandler.ginHandlerError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ginhandler.ginHandlerError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "model.AddIngredientReq": {
            "type": "object",
            "required": [
                "ingredient_id",
                "quantity"
            ],
            "properties": {
                "ingredient_id": {
                    "type": "integer",
                    "example": 1
                },
                "quantity": {
                    "type": "integer",
                    "example": 10
                }
            }
        },
        "model.AddStepReq": {
            "type": "object",
            "required": [
                "description",
                "duration"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer",
                    "example": 7200
                }
            }
        },
        "model.CreateRecipeReq": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "ingredients": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.AddIngredientReq"
                    }
                },
                "name": {
                    "type": "string"
                },
                "steps": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.CreateStepReq"
                    }
                }
            }
        },
        "model.CreateStepReq": {
            "type": "object",
            "required": [
                "description",
                "duration"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer",
                    "example": 240
                },
                "recipe_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "model.GetAllRecipesReq": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "integer",
                    "example": 1
                },
                "duration_filter": {
                    "type": "integer",
                    "example": 7200
                },
                "duration_sort": {
                    "type": "string",
                    "example": "DESC"
                },
                "ingredient_filter": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    },
                    "example": [
                        1
                    ]
                }
            }
        },
        "model.SigninUserReq": {
            "type": "object",
            "required": [
                "Login",
                "password"
            ],
            "properties": {
                "Login": {
                    "type": "string",
                    "example": "user1"
                },
                "password": {
                    "type": "string",
                    "example": "user1"
                }
            }
        },
        "model.SignupUserReq": {
            "type": "object",
            "required": [
                "Login",
                "name",
                "password"
            ],
            "properties": {
                "Login": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.UpdateRecipeReq": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "localhost:8000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Recipe App API",
	Description:      "API Server for Recipe Application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
