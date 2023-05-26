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
        "/getBooks": {
            "get": {
                "description": "Get all books from the database",
                "produces": [
                    "application/json"
                ],
                "summary": "Get all books",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/cmd.Books"
                            }
                        }
                    }
                }
            }
        },
        "/getBooksByAuthor": {
            "get": {
                "description": "Get books by author from the database",
                "produces": [
                    "application/json"
                ],
                "summary": "Get books by author",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Author name",
                        "name": "author",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/cmd.Books"
                            }
                        }
                    }
                }
            }
        },
        "/getBooksById": {
            "get": {
                "description": "Get a book by ID from the database",
                "produces": [
                    "application/json"
                ],
                "summary": "Get book by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Book ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/cmd.Books"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "cmd.Books": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isbn": {
                    "type": "string"
                },
                "publication_year": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Book Catalogue APIs",
	Description:      "A book catalogue service using gin-framework",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
