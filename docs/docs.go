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
        "/lists": {
            "get": {
                "description": "Get all Lists",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "List"
                ],
                "summary": "Get all Lists",
                "responses": {}
            },
            "post": {
                "description": "Create a new List from JSON. \"Order\" field will be 0 if omitted.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "List"
                ],
                "summary": "Create a new List",
                "parameters": [
                    {
                        "description": "CreateListDto",
                        "name": "ListDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.CreateListDto"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/lists/{id}": {
            "get": {
                "description": "Get a List with ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "List"
                ],
                "summary": "Get a List",
                "parameters": [
                    {
                        "type": "string",
                        "description": "List ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "put": {
                "description": "Update a List of ID with JSON. Only changes the fields that are in the JSON.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "List"
                ],
                "summary": "Update a List",
                "parameters": [
                    {
                        "type": "string",
                        "description": "List ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "UpdateListDto",
                        "name": "ListDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.UpdateListDto"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Delete a List with ID, and its Tasks",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "List"
                ],
                "summary": "Delete a List",
                "parameters": [
                    {
                        "type": "string",
                        "description": "List ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/tasks": {
            "get": {
                "description": "Get all Tasks",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Get all Tasks",
                "responses": {}
            },
            "post": {
                "description": "Create a new Task from JSON. \"Order\" field will be 0 if omitted.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Create a new Task",
                "parameters": [
                    {
                        "description": "CreateTaskDto",
                        "name": "TaskDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.CreateTaskDto"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/tasks/{id}": {
            "get": {
                "description": "Get a Task with ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Get a Task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "put": {
                "description": "Update a Task of ID with JSON. Only changes the fields that are in the JSON.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Update a Task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "UpdateTaskDto",
                        "name": "TaskDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.UpdateTaskDto"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Delete a Task with ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Delete a Task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "services.CreateListDto": {
            "type": "object",
            "properties": {
                "order": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "services.CreateTaskDto": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "dueDate": {
                    "type": "string"
                },
                "listID": {
                    "type": "integer"
                },
                "order": {
                    "type": "integer"
                }
            }
        },
        "services.UpdateListDto": {
            "type": "object",
            "properties": {
                "order": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "services.UpdateTaskDto": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "dueDate": {
                    "type": "string"
                },
                "listID": {
                    "type": "integer"
                },
                "order": {
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
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
