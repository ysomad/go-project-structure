// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "title": "Server",
    "version": "1.0.0"
  },
  "paths": {
    "/log-level": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "logging"
        ],
        "summary": "Текущий уровень логирования",
        "operationId": "getLogLevel",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/LogLevel"
            }
          }
        }
      },
      "put": {
        "security": [
          {
            "apiKey": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "tags": [
          "logging"
        ],
        "summary": "Обновить уровень логирования",
        "operationId": "updateLogLevel",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/LogLevel"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "No Content"
          }
        }
      }
    },
    "/ping": {
      "get": {
        "tags": [
          "health"
        ],
        "summary": "Пинг",
        "operationId": "ping",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/PingResponse"
            }
          }
        }
      }
    },
    "/v1/products": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "product"
        ],
        "summary": "Список товаров",
        "operationId": "list_products_v1",
        "parameters": [
          {
            "enum": [
              "men",
              "women"
            ],
            "type": "string",
            "name": "gender",
            "in": "query",
            "required": true
          },
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "name": "categories",
            "in": "query"
          },
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "name": "brands",
            "in": "query"
          },
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "name": "colors",
            "in": "query"
          },
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "name": "sizes",
            "in": "query"
          },
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "name": "labelsin",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "name": "min_price",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "name": "max_price",
            "in": "query"
          },
          {
            "enum": [
              "price_asc",
              "price_desc"
            ],
            "type": "string",
            "name": "sort",
            "in": "query"
          },
          {
            "maximum": 100,
            "type": "integer",
            "format": "int32",
            "default": 60,
            "name": "page_size",
            "in": "query"
          },
          {
            "type": "string",
            "name": "page_token",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/ListProductsResponse"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/v1/products/{product_id}": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "product"
        ],
        "summary": "Карточка товара",
        "operationId": "get_product",
        "parameters": [
          {
            "type": "string",
            "name": "product_id",
            "in": "path",
            "required": true
          },
          {
            "enum": [
              "men",
              "women"
            ],
            "type": "string",
            "name": "gender",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/ProductCard"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ApiKey": {
      "type": "string"
    },
    "Brand": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "Category": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "Color": {
      "type": "object",
      "properties": {
        "hex": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "DeliveryPeriod": {
      "type": "object",
      "properties": {
        "end_day": {
          "type": "integer",
          "format": "int32"
        },
        "start_day": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "Error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "ListProductsResponse": {
      "type": "object",
      "properties": {
        "next_page_token": {
          "type": "string"
        },
        "products": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Product"
          }
        },
        "total_pages": {
          "type": "number",
          "format": "int32"
        }
      }
    },
    "LogLevel": {
      "type": "object",
      "properties": {
        "level": {
          "type": "string",
          "enum": [
            "TRACE",
            "DEBUG",
            "INFO",
            "WARN",
            "ERROR",
            "FATAL"
          ]
        }
      }
    },
    "OfferMerchant": {
      "type": "object",
      "properties": {
        "icon_url": {
          "type": "string",
          "format": "uri"
        },
        "logo_text": {
          "type": "string"
        },
        "logo_url": {
          "type": "string",
          "format": "uri"
        },
        "name": {
          "type": "string"
        },
        "official": {
          "type": "boolean"
        },
        "service_name": {
          "type": "string"
        }
      }
    },
    "PingResponse": {
      "type": "object",
      "properties": {
        "status": {
          "type": "string"
        }
      }
    },
    "Price": {
      "type": "object",
      "properties": {
        "currency_code": {
          "type": "string"
        },
        "current": {
          "type": "integer",
          "format": "int32"
        },
        "discount_percentage": {
          "type": "number",
          "format": "double"
        },
        "old": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "PriceWithShipment": {
      "type": "object",
      "properties": {
        "currency_code": {
          "type": "string"
        },
        "current": {
          "type": "integer",
          "format": "int32"
        },
        "discount_percentage": {
          "type": "number",
          "format": "double"
        },
        "old": {
          "type": "integer",
          "format": "int32"
        },
        "shipment": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "Product": {
      "type": "object",
      "properties": {
        "brand": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "images": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "uri"
          }
        },
        "label": {
          "type": "string"
        },
        "merchant_count": {
          "type": "integer",
          "format": "int32"
        },
        "merchant_icons": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "uri"
          }
        },
        "price": {
          "$ref": "#/definitions/Price"
        },
        "sizes": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "title": {
          "type": "string"
        }
      }
    },
    "ProductCard": {
      "type": "object",
      "properties": {
        "brand": {
          "$ref": "#/definitions/Brand"
        },
        "categories": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Category"
          }
        },
        "color": {
          "$ref": "#/definitions/Color"
        },
        "description": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "images": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "uri"
          }
        },
        "label": {
          "type": "string"
        },
        "merchant_count": {
          "type": "integer",
          "format": "int32"
        },
        "merchant_icons": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "uri"
          }
        },
        "price": {
          "$ref": "#/definitions/PriceWithShipment"
        },
        "title": {
          "type": "string"
        },
        "variants": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ProductVariant"
          }
        }
      }
    },
    "ProductVariant": {
      "type": "object",
      "properties": {
        "is_default": {
          "type": "boolean"
        },
        "min_price": {
          "$ref": "#/definitions/PriceWithShipment"
        },
        "offers": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ProductVariantOffer"
          }
        },
        "scale_abbreviation": {
          "type": "string"
        },
        "scale_description": {
          "type": "string"
        },
        "size": {
          "type": "string"
        }
      }
    },
    "ProductVariantOffer": {
      "type": "object",
      "properties": {
        "delivery": {
          "$ref": "#/definitions/DeliveryPeriod"
        },
        "in_stock": {
          "type": "boolean"
        },
        "last_left": {
          "type": "boolean"
        },
        "merchant": {
          "$ref": "#/definitions/OfferMerchant"
        },
        "price": {
          "$ref": "#/definitions/PriceWithShipment"
        },
        "return_policy": {
          "type": "boolean"
        },
        "variant_id": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "apiKey": {
      "type": "apiKey",
      "name": "X-API-Key",
      "in": "header"
    }
  },
  "tags": [
    {
      "description": "Товары",
      "name": "product"
    },
    {
      "description": "Логирование",
      "name": "logging"
    },
    {
      "name": "health"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "title": "Server",
    "version": "1.0.0"
  },
  "paths": {
    "/log-level": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "logging"
        ],
        "summary": "Текущий уровень логирования",
        "operationId": "getLogLevel",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/LogLevel"
            }
          }
        }
      },
      "put": {
        "security": [
          {
            "apiKey": []
          }
        ],
        "consumes": [
          "application/json"
        ],
        "tags": [
          "logging"
        ],
        "summary": "Обновить уровень логирования",
        "operationId": "updateLogLevel",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/LogLevel"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "No Content"
          }
        }
      }
    },
    "/ping": {
      "get": {
        "tags": [
          "health"
        ],
        "summary": "Пинг",
        "operationId": "ping",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/PingResponse"
            }
          }
        }
      }
    },
    "/v1/products": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "product"
        ],
        "summary": "Список товаров",
        "operationId": "list_products_v1",
        "parameters": [
          {
            "enum": [
              "men",
              "women"
            ],
            "type": "string",
            "name": "gender",
            "in": "query",
            "required": true
          },
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "name": "categories",
            "in": "query"
          },
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "name": "brands",
            "in": "query"
          },
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "name": "colors",
            "in": "query"
          },
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "name": "sizes",
            "in": "query"
          },
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "name": "labelsin",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "name": "min_price",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "name": "max_price",
            "in": "query"
          },
          {
            "enum": [
              "price_asc",
              "price_desc"
            ],
            "type": "string",
            "name": "sort",
            "in": "query"
          },
          {
            "maximum": 100,
            "type": "integer",
            "format": "int32",
            "default": 60,
            "name": "page_size",
            "in": "query"
          },
          {
            "type": "string",
            "name": "page_token",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/ListProductsResponse"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/v1/products/{product_id}": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "product"
        ],
        "summary": "Карточка товара",
        "operationId": "get_product",
        "parameters": [
          {
            "type": "string",
            "name": "product_id",
            "in": "path",
            "required": true
          },
          {
            "enum": [
              "men",
              "women"
            ],
            "type": "string",
            "name": "gender",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/ProductCard"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ApiKey": {
      "type": "string"
    },
    "Brand": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "Category": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "Color": {
      "type": "object",
      "properties": {
        "hex": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "DeliveryPeriod": {
      "type": "object",
      "properties": {
        "end_day": {
          "type": "integer",
          "format": "int32"
        },
        "start_day": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "Error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "ListProductsResponse": {
      "type": "object",
      "properties": {
        "next_page_token": {
          "type": "string"
        },
        "products": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Product"
          }
        },
        "total_pages": {
          "type": "number",
          "format": "int32"
        }
      }
    },
    "LogLevel": {
      "type": "object",
      "properties": {
        "level": {
          "type": "string",
          "enum": [
            "TRACE",
            "DEBUG",
            "INFO",
            "WARN",
            "ERROR",
            "FATAL"
          ]
        }
      }
    },
    "OfferMerchant": {
      "type": "object",
      "properties": {
        "icon_url": {
          "type": "string",
          "format": "uri"
        },
        "logo_text": {
          "type": "string"
        },
        "logo_url": {
          "type": "string",
          "format": "uri"
        },
        "name": {
          "type": "string"
        },
        "official": {
          "type": "boolean"
        },
        "service_name": {
          "type": "string"
        }
      }
    },
    "PingResponse": {
      "type": "object",
      "properties": {
        "status": {
          "type": "string"
        }
      }
    },
    "Price": {
      "type": "object",
      "properties": {
        "currency_code": {
          "type": "string"
        },
        "current": {
          "type": "integer",
          "format": "int32"
        },
        "discount_percentage": {
          "type": "number",
          "format": "double"
        },
        "old": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "PriceWithShipment": {
      "type": "object",
      "properties": {
        "currency_code": {
          "type": "string"
        },
        "current": {
          "type": "integer",
          "format": "int32"
        },
        "discount_percentage": {
          "type": "number",
          "format": "double"
        },
        "old": {
          "type": "integer",
          "format": "int32"
        },
        "shipment": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "Product": {
      "type": "object",
      "properties": {
        "brand": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "images": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "uri"
          }
        },
        "label": {
          "type": "string"
        },
        "merchant_count": {
          "type": "integer",
          "format": "int32"
        },
        "merchant_icons": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "uri"
          }
        },
        "price": {
          "$ref": "#/definitions/Price"
        },
        "sizes": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "title": {
          "type": "string"
        }
      }
    },
    "ProductCard": {
      "type": "object",
      "properties": {
        "brand": {
          "$ref": "#/definitions/Brand"
        },
        "categories": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Category"
          }
        },
        "color": {
          "$ref": "#/definitions/Color"
        },
        "description": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "images": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "uri"
          }
        },
        "label": {
          "type": "string"
        },
        "merchant_count": {
          "type": "integer",
          "format": "int32"
        },
        "merchant_icons": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "uri"
          }
        },
        "price": {
          "$ref": "#/definitions/PriceWithShipment"
        },
        "title": {
          "type": "string"
        },
        "variants": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ProductVariant"
          }
        }
      }
    },
    "ProductVariant": {
      "type": "object",
      "properties": {
        "is_default": {
          "type": "boolean"
        },
        "min_price": {
          "$ref": "#/definitions/PriceWithShipment"
        },
        "offers": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ProductVariantOffer"
          }
        },
        "scale_abbreviation": {
          "type": "string"
        },
        "scale_description": {
          "type": "string"
        },
        "size": {
          "type": "string"
        }
      }
    },
    "ProductVariantOffer": {
      "type": "object",
      "properties": {
        "delivery": {
          "$ref": "#/definitions/DeliveryPeriod"
        },
        "in_stock": {
          "type": "boolean"
        },
        "last_left": {
          "type": "boolean"
        },
        "merchant": {
          "$ref": "#/definitions/OfferMerchant"
        },
        "price": {
          "$ref": "#/definitions/PriceWithShipment"
        },
        "return_policy": {
          "type": "boolean"
        },
        "variant_id": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "apiKey": {
      "type": "apiKey",
      "name": "X-API-Key",
      "in": "header"
    }
  },
  "tags": [
    {
      "description": "Товары",
      "name": "product"
    },
    {
      "description": "Логирование",
      "name": "logging"
    },
    {
      "name": "health"
    }
  ]
}`))
}
