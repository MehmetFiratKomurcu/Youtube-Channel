// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/orders": {
            "post": {
                "description": "Creating Order with given request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Creating Order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "code of Order",
                        "name": "x-correlationid",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Request of Creating Order Object",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/orders/code/{orderCode}": {
            "get": {
                "description": "Getting Order by Code in detail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Getting Order by Code",
                "parameters": [
                    {
                        "type": "string",
                        "description": "code of Order",
                        "name": "x-correlationid",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "code of Order",
                        "name": "orderCode",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.CreateOrderRequest": {
            "description": "Request about creating Order",
            "type": "object",
            "required": [
                "age",
                "countryCode",
                "shipmentNumber"
            ],
            "properties": {
                "age": {
                    "description": "age to make sure you are young",
                    "type": "integer"
                },
                "countryCode": {
                    "description": "country code like: tr, us",
                    "type": "string"
                },
                "shipmentNumber": {
                    "description": "shipment no of Order",
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
	Title:            "Order Api",
	Description:      "This is an Order Api just for young people",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}