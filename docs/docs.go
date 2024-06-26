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
        "/api/v1/change-password": {
            "post": {
                "description": "Final step to user recovery password",
                "tags": [
                    "User"
                ],
                "summary": "Change user password",
                "parameters": [
                    {
                        "description": "Payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.ChangePasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.UserDto"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/exceptions.DomainError"
                        }
                    }
                }
            }
        },
        "/api/v1/forgot-password": {
            "post": {
                "description": "First step to user recovery password",
                "tags": [
                    "User"
                ],
                "summary": "Forgot user password",
                "parameters": [
                    {
                        "description": "Payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.ForgotPasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.UserDto"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/exceptions.DomainError"
                        }
                    }
                }
            }
        },
        "/api/v1/health": {
            "get": {
                "tags": [
                    "Utils"
                ],
                "summary": "API health check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/v1/login": {
            "post": {
                "description": "Authentication",
                "tags": [
                    "User"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.jwt"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/exceptions.DomainError"
                        }
                    }
                }
            }
        },
        "/api/v1/users": {
            "post": {
                "description": "Create a new user",
                "tags": [
                    "User"
                ],
                "summary": "User signup",
                "parameters": [
                    {
                        "description": "Payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.SignupRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.UserDto"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/exceptions.DomainError"
                        }
                    }
                }
            }
        },
        "/api/v1/validate-code": {
            "post": {
                "description": "Second step to user recovery password",
                "tags": [
                    "User"
                ],
                "summary": "Validate code",
                "parameters": [
                    {
                        "description": "Payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.ValidateCodeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/exceptions.DomainError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.UserDto": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "exceptions.DomainError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "internalCode": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handlers.jwt": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "requests.ChangePasswordRequest": {
            "type": "object",
            "required": [
                "code",
                "email",
                "password"
            ],
            "properties": {
                "code": {
                    "type": "string",
                    "maxLength": 5,
                    "minLength": 5
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                }
            }
        },
        "requests.ForgotPasswordRequest": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "requests.LoginRequest": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 255
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "requests.SignupRequest": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 255
                },
                "name": {
                    "type": "string",
                    "maxLength": 255
                },
                "password": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 6
                }
            }
        },
        "requests.ValidateCodeRequest": {
            "type": "object",
            "required": [
                "code",
                "email"
            ],
            "properties": {
                "code": {
                    "type": "string",
                    "maxLength": 5,
                    "minLength": 5
                },
                "email": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Squad 10x Golang boilerplate",
	Description:      "Swagger iterativo para API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
