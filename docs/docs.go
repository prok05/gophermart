// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/user/balance": {
            "get": {
                "security": [
                    {
                        "AuthToken": []
                    }
                ],
                "description": "Retrieve user balance",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "User balance",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.UserBalance"
                        }
                    },
                    "401": {
                        "description": "User is not authenticated",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Server encountered a problem",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/user/balance/withdraw": {
            "post": {
                "security": [
                    {
                        "AuthToken": []
                    }
                ],
                "description": "Request for withdrawing user balance",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Withdraw balance",
                "parameters": [
                    {
                        "description": "Order number and amount",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.WithdrawBalance"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "401": {
                        "description": "User is not authenticated",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "402": {
                        "description": "Balance is not enough",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "422": {
                        "description": "Wrong order number format",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Server encountered a problem",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Authenticates user with given login and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.LoginUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "401": {
                        "description": "Wrong login/password",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Server encountered a problem",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/user/orders": {
            "get": {
                "security": [
                    {
                        "AuthToken": []
                    }
                ],
                "description": "Get user orders",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Get user orders",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.UserOrder"
                            }
                        }
                    },
                    "204": {
                        "description": "Data is empty",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "401": {
                        "description": "User is not authenticated",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Server encountered a problem",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "AuthToken": []
                    }
                ],
                "description": "Load order for future processing",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Load order",
                "parameters": [
                    {
                        "description": "Order number",
                        "name": "orderNumber",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Order was already loaded",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "202": {
                        "description": "Accepted"
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "401": {
                        "description": "User is not authenticated",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "409": {
                        "description": "Order was loaded by another user",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "422": {
                        "description": "Wrong order number format",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Server encountered a problem",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "Creates a new user with given login and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register new user",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RegisterUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "409": {
                        "description": "Resource already exists",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Server encountered a problem",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/user/withdrawals": {
            "get": {
                "security": [
                    {
                        "AuthToken": []
                    }
                ],
                "description": "Gets user withdrawals",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "User withdrawals",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.UserWithdrawal"
                            }
                        }
                    },
                    "204": {
                        "description": "Data is empty",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "401": {
                        "description": "User is not authenticated",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Server encountered a problem",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.OrderStatus": {
            "type": "string",
            "enum": [
                "NEW",
                "PROCESSING",
                "INVALID",
                "PROCESSED"
            ],
            "x-enum-varnames": [
                "OrderNewStatus",
                "OrderProcessingStatus",
                "OrderInvalidStatus",
                "OrderProcessedStatus"
            ]
        },
        "entity.UserBalance": {
            "type": "object",
            "properties": {
                "current": {
                    "type": "number",
                    "example": 500.5
                },
                "user_id": {
                    "type": "string",
                    "example": "UUID"
                },
                "withdrawn": {
                    "type": "number",
                    "example": 42
                }
            }
        },
        "entity.UserOrder": {
            "type": "object",
            "properties": {
                "accrual": {
                    "type": "number"
                },
                "id": {
                    "type": "string"
                },
                "number": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/entity.OrderStatus"
                },
                "uploaded_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "entity.UserWithdrawal": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "example": "UUID"
                },
                "order": {
                    "type": "string",
                    "example": "2377225624"
                },
                "processed_at": {
                    "type": "string",
                    "example": "2020-12-09T16:09:57+03:00"
                },
                "sum": {
                    "type": "number",
                    "example": 500
                },
                "user_id": {
                    "type": "string",
                    "example": "UUID"
                }
            }
        },
        "request.LoginUser": {
            "type": "object",
            "required": [
                "login",
                "password"
            ],
            "properties": {
                "login": {
                    "type": "string",
                    "maxLength": 100,
                    "example": "login"
                },
                "password": {
                    "type": "string",
                    "maxLength": 72,
                    "minLength": 8,
                    "example": "password"
                }
            }
        },
        "request.RegisterUser": {
            "type": "object",
            "required": [
                "login",
                "password"
            ],
            "properties": {
                "login": {
                    "type": "string",
                    "maxLength": 100,
                    "example": "login"
                },
                "password": {
                    "type": "string",
                    "maxLength": 72,
                    "minLength": 8,
                    "example": "password"
                }
            }
        },
        "request.WithdrawBalance": {
            "type": "object",
            "required": [
                "order",
                "sum"
            ],
            "properties": {
                "order": {
                    "type": "string",
                    "maxLength": 100,
                    "example": "2377225624"
                },
                "sum": {
                    "type": "number",
                    "example": 751
                }
            }
        },
        "response.Error": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "response.LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string",
                    "example": "token"
                }
            }
        }
    },
    "securityDefinitions": {
        "AuthToken": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:80",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Gophermart",
	Description:      "Service for orders",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
