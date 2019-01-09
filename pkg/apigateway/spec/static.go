// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// Code generated by "makestatic"; DO NOT EDIT.

package static

var Files = map[string]string{
	"api.swagger.json": `{
  "swagger": "2.0",
  "info": {
    "title": "Notification Project",
    "version": "0.0.1",
    "contact": {
      "name": "Notification Project",
      "url": "https://github.com/openpitrix/notification"
    }
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/SayHello": {
      "post": {
        "summary": "#API 0.SayHello:gRPC testing,Sends a greeting.",
        "operationId": "SayHello",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbHelloReply"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbHelloRequest"
            }
          }
        ],
        "tags": [
          "notification"
        ]
      }
    },
    "/v1/nf/DescribeNfs": {
      "post": {
        "summary": "#API 2.DescribeNfs:describe single Notification with filter.",
        "operationId": "DescribeNfs",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbDescribeNfsResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbDescribeNfsRequest"
            }
          }
        ],
        "tags": [
          "notification"
        ]
      }
    },
    "/v1/notification/CreateNfWithAddrs": {
      "post": {
        "summary": "#API 1.CreateNfWithAddrs：create notification with addrs(email addrs, phone numbers).",
        "operationId": "CreateNfWithAddrs",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbCreateNfResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbCreateNfWithAddrsRequest"
            }
          }
        ],
        "tags": [
          "notification"
        ]
      }
    }
  },
  "definitions": {
    "pbCreateNfResponse": {
      "type": "object",
      "properties": {
        "notification_id": {
          "type": "string"
        }
      }
    },
    "pbCreateNfWithAddrsRequest": {
      "type": "object",
      "properties": {
        "content_type": {
          "type": "string"
        },
        "sent_type": {
          "type": "string"
        },
        "addrs_str": {
          "type": "string"
        },
        "title": {
          "type": "string"
        },
        "content": {
          "type": "string"
        },
        "short_content": {
          "type": "string"
        },
        "expired_days": {
          "type": "string"
        },
        "owner": {
          "type": "string"
        },
        "status": {
          "type": "string"
        }
      }
    },
    "pbDescribeNfsRequest": {
      "type": "object",
      "properties": {
        "content_type": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "sent_type": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "title": {
          "type": "string"
        },
        "content": {
          "type": "string"
        },
        "owner": {
          "type": "string"
        },
        "userids_str": {
          "type": "string"
        },
        "status": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "limit": {
          "type": "integer",
          "format": "int64"
        },
        "offset": {
          "type": "integer",
          "format": "int64"
        },
        "sort_key": {
          "type": "string"
        }
      }
    },
    "pbDescribeNfsResponse": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "pbHelloReply": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      },
      "title": "The response message containing the greetings"
    },
    "pbHelloRequest": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "BearerAuth": {
      "type": "apiKey",
      "description": "The Authorization header must be set to Bearer followed by a space and a token. For example, 'Bearer vHUabiBEIKi8n1RdvWOjGFulGSM6zunb'.",
      "name": "Authorization",
      "in": "header"
    }
  },
  "security": [
    {
      "BearerAuth": []
    }
  ]
}
`,
}
