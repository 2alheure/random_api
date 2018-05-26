define({ "api": [
  {
    "type": "get",
    "url": "/ping",
    "title": "Request to ping the server and check if it's working",
    "version": "1.0.0",
    "name": "Ping",
    "group": "Ping",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "pong",
            "description": "<p>The user pong</p>"
          }
        ]
      }
    },
    "filename": "./main.go",
    "groupTitle": "Ping",
    "sampleRequest": [
      {
        "url": "https://localhost:8000/ping"
      }
    ]
  },
  {
    "type": "get",
    "url": "/pong",
    "title": "Request to check if apidoc put a new doc file",
    "version": "1.0.0",
    "name": "Ping",
    "group": "Pong",
    "filename": "./main.go",
    "groupTitle": "Pong",
    "sampleRequest": [
      {
        "url": "https://localhost:8000/pong"
      }
    ]
  },
  {
    "type": "get",
    "url": "/ressource/:max",
    "title": "Request to get ressources",
    "version": "1.0.0",
    "name": "GetRessources",
    "group": "Ressources",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "max",
            "description": "<p>Set how many ressources you want to get</p>"
          }
        ]
      }
    },
    "filename": "./main.go",
    "groupTitle": "Ressources",
    "sampleRequest": [
      {
        "url": "https://localhost:8000/ressource/:max"
      }
    ]
  }
] });
