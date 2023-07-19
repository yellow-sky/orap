// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/ap/connections": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Get list of AP connection settings with short info: id, name, description and connection state.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ap_info"
                ],
                "summary": "List of AP connections",
                "responses": {
                    "200": {
                        "description": "Common JSON response with list of AP connections",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/api.CommonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/nm_connection_manager.NmWifiConnectionSettings"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized error",
                        "schema": {
                            "$ref": "#/definitions/api.CommonResponse"
                        }
                    },
                    "500": {
                        "description": "Unhandled server error",
                        "schema": {
                            "$ref": "#/definitions/api.CommonResponse"
                        }
                    }
                }
            }
        },
        "/auth/token": {
            "post": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Get JWT token for sent username and password.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "User Login method",
                "responses": {
                    "200": {
                        "description": "Common JSON response with JWT as data",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/api.CommonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized error",
                        "schema": {
                            "$ref": "#/definitions/api.CommonResponse"
                        }
                    },
                    "500": {
                        "description": "Unhandled server error",
                        "schema": {
                            "$ref": "#/definitions/api.CommonResponse"
                        }
                    }
                }
            }
        },
        "/devices": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Get list of devices with short info: id, name, description and connection state.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "devices_info"
                ],
                "summary": "List of network devices",
                "responses": {
                    "200": {
                        "description": "Common JSON response with list of devices",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/api.CommonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/nm_device_manager.DeviceShortInfo"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized error",
                        "schema": {
                            "$ref": "#/definitions/api.CommonResponse"
                        }
                    },
                    "500": {
                        "description": "Unhandled server error",
                        "schema": {
                            "$ref": "#/definitions/api.CommonResponse"
                        }
                    }
                }
            }
        },
        "/devices/{device_id}": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Get info about device: common info, state, etc.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "devices_info"
                ],
                "summary": "Device detailed information",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Device id from list of devices",
                        "name": "device_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Common JSON response with modem detailed info as data",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/api.CommonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/nm_device_manager.DeviceDetailedInfo"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized error",
                        "schema": {
                            "$ref": "#/definitions/api.CommonResponse"
                        }
                    },
                    "404": {
                        "description": "Device not found error",
                        "schema": {
                            "$ref": "#/definitions/api.CommonResponse"
                        }
                    },
                    "500": {
                        "description": "Unhandled server error",
                        "schema": {
                            "$ref": "#/definitions/api.CommonResponse"
                        }
                    }
                }
            }
        },
        "/modems": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Get list of modems with short info: id, name, description and connection state.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "modems_info"
                ],
                "summary": "List of modems",
                "responses": {
                    "200": {
                        "description": "Common JSON response with list of modems",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/api.CommonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/api.ModemShortInfo"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized error",
                        "schema": {
                            "$ref": "#/definitions/api.CommonResponse"
                        }
                    },
                    "500": {
                        "description": "Unhandled server error",
                        "schema": {
                            "$ref": "#/definitions/api.CommonResponse"
                        }
                    }
                }
            }
        },
        "/modems/{modem_id}": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Get info about modem: common info, state, etc.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "modems_info"
                ],
                "summary": "Modem detailed information",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Modem id from list of modems",
                        "name": "modem_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Common JSON response with modem detailed info as data",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/api.CommonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized error",
                        "schema": {
                            "$ref": "#/definitions/api.CommonResponse"
                        }
                    },
                    "404": {
                        "description": "Modem not found error",
                        "schema": {
                            "$ref": "#/definitions/api.CommonResponse"
                        }
                    },
                    "500": {
                        "description": "Unhandled server error",
                        "schema": {
                            "$ref": "#/definitions/api.CommonResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.CommonResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "error": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "api.ModemShortInfo": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "nm_connection_manager.NmCommonSettings": {
            "type": "object",
            "properties": {
                "autoconnect": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "permissions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "timestamp": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "nm_connection_manager.NmWifiConnectionSettings": {
            "type": "object",
            "properties": {
                "802-11-wireless": {
                    "$ref": "#/definitions/nm_connection_manager.NmWirelessSettings"
                },
                "802-11-wireless-security": {
                    "$ref": "#/definitions/nm_connection_manager.NmWirelessSecuritySettings"
                },
                "connection": {
                    "$ref": "#/definitions/nm_connection_manager.NmCommonSettings"
                }
            }
        },
        "nm_connection_manager.NmWirelessSecuritySettings": {
            "type": "object",
            "properties": {
                "key-mgmt": {
                    "type": "string"
                }
            }
        },
        "nm_connection_manager.NmWirelessSettings": {
            "type": "object",
            "properties": {
                "mac-address": {
                    "type": "string"
                },
                "mac-address-blacklist": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "mode": {
                    "type": "string"
                },
                "security": {
                    "type": "string"
                },
                "seen-bssids": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "ssid": {
                    "description": "SSID              []uint8  ` + "`" + `mapstructure:\"ssid\"` + "`" + `",
                    "type": "string"
                }
            }
        },
        "nm_device_manager.DeviceDetailedInfo": {
            "type": "object",
            "properties": {
                "auto_connect": {
                    "type": "boolean"
                },
                "driver": {
                    "type": "string"
                },
                "driver_version": {
                    "type": "string"
                },
                "firmware_missing": {
                    "type": "boolean"
                },
                "fw_version": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "interface": {
                    "type": "string"
                },
                "managed": {
                    "type": "boolean"
                },
                "mtu": {
                    "type": "integer"
                },
                "nm_plugin_missing": {
                    "type": "boolean"
                },
                "physical_port_id": {
                    "type": "string"
                },
                "real": {
                    "type": "boolean"
                },
                "state": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "nm_device_manager.DeviceShortInfo": {
            "type": "object",
            "properties": {
                "driver": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "interface": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "tags": [
        {
            "description": "Auth operations",
            "name": "auth"
        },
        {
            "description": "Common operations",
            "name": "common"
        },
        {
            "description": "Devices info operations",
            "name": "devices_info"
        },
        {
            "description": "Modems info operations",
            "name": "modems_info"
        }
    ]
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.1",
	Host:        "",
	BasePath:    "/api",
	Schemes:     []string{},
	Title:       "OrAP Service API",
	Description: "API for network management - AP, ethernet, modems.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
