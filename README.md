![alt text][banner-short]


# Palette Town API
A RESTful API for generating appealing 5 colour, colour palettes.

## Installation & Run
```bash
# Download this project
go get github.com/MattiasHenders/palette-town-api
```

```bash
# Build and Run on MacOS/Linux
make run-server

# Build and Run on Windows
make run-server-windows

# API Endpoint : http://localhost:8080
```

## API

### Example Response JSON Structure
```bash
{
    "message": "Successfully got colour palette from given word",
    "givenInput": "forest",
    "code": 200,
    "data": {
        "colours": [
            "#747f41",
            "#538443",
            "#4f632d",
            "#8f995d",
            "#d4cc7d"
        ]
    },
    "coolorsLink": "https://coolors.co/747f41-538443-4f632d-8f995d-d4cc7d"
}
```

#### /api/colour/random
* `GET` : Get a random colour palette

#### /api/colour/colours
* `GET` : Get a colour palette based on an input of colours

#### /api/colour/words
* `GET` : Get a colour palette based on an input of words

## Todo

- [x] Support basic REST APIs.
- [ ] Support Authentication with user for securing the APIs.
- [ ] Write the tests for all APIs.
- [x] Organize the code with packages
- [ ] Make docs with GoDoc
- [ ] Building a deployment process 
- [ ] Connect to a database for users
- [ ] Host the Colourminds API on own server

## Structure
```
├── src
│   ├── main.go
│   ├── handlers                    
│   │   ├── colours.go               // Common response functions for colour endpoints
│   │   └── users.go                 // Common response functions for user endpoints
│   ├── internal                    
│   │   ├── errors                  
|   │   │   └── errors.go            // Internal error models and functions
│   │   └── server_helpers          
|   │       └── server_helpers.go    // Internal error models and functions
│   ├── model
│   |   ├── colours.go               // Models for colours
│   |   ├── users.go                 // Models for users
│   |   └── server.go                // Models for the server
│   └── pkgs
│       └── colours.go               // Functions for generating colour palettes
├── config
│   └── config.go                    // Configuration
├── tests
│   └── colours                    
│       └── colours_test.go          // Tests for colour functions
└── main.go
```


[banner-short]: https://github.com/MattiasHenders/palette-town-api/blob/main/assets/banner-short.png?raw=true "Palette Town API Banner"
[banner]: https://github.com/MattiasHenders/palette-town-api/blob/main/assets/banner.png?raw=true "Palette Town API Full Banner"
