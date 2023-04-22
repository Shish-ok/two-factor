// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/send": {
            "post": {
                "description": "Отправляет код двухфакторной авторизации, возвращая id запроса и сам код\nКод ошибки 400: неверный json или невалидный номер",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Отправляет код для двухфакторной авторизации",
                "parameters": [
                    {
                        "description": "Входные параметры",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/two-factor-auth_internal_api.SendCodeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/confirmations.Confirmation"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/two-factor-auth_internal_api.Error"
                        }
                    }
                }
            }
        },
        "/verify": {
            "post": {
                "description": "Верифицирует код двухфакторной авторизации, возвращая текущую дату в UnixTime\nКод ошибки 400: неверный json или невалидный код авторизации\nКод ошибки 403: исчерпан лимит попыток или код уже не действителен",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Верифицирует код двухфакторной авторизации",
                "parameters": [
                    {
                        "description": "Входные параметры",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/confirmations.Confirmation"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal_api.VerifyCodeResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal_api.Error"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/internal_api.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "confirmations.Confirmation": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "requestId": {
                    "type": "string"
                }
            }
        },
        "internal_api.Error": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "internal_api.SendCodeRequest": {
            "type": "object",
            "properties": {
                "number": {
                    "type": "string"
                }
            }
        },
        "internal_api.VerifyCodeResponse": {
            "type": "object",
            "properties": {
                "verifiedAt": {
                    "type": "integer"
                }
            }
        },
        "two-factor-auth_internal_api.Error": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "two-factor-auth_internal_api.SendCodeRequest": {
            "type": "object",
            "properties": {
                "number": {
                    "type": "string"
                }
            }
        },
        "two-factor-auth_internal_api.VerifyCodeResponse": {
            "type": "object",
            "properties": {
                "verifiedAt": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "two-factor-auth doc",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}