![Palette Town API Banner][banner-short]

# Palette Town API
A RESTful API for generating appealing 5 colour, colour palettes.

## Setup
### Installation
```bash
# Download this project
git clone https://github.com/MattiasHenders/palette-town-api.git
```
### Run
```bash
# Build and Run on MacOS/Linux
make run-server

# Build and Run on Windows
make run-server-windows

# API Endpoint : http://localhost:8080
```
### Environment Variables
```bash
# Create .env in root, see .env.example for needed variables
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

## Example Palettes
* `Endpoint: /api/colour/random`
* ![image](https://user-images.githubusercontent.com/59512495/198418359-ebaf2288-92ae-44c7-ae1f-30a49ad548f8.png)

* `Endpoint: /api/colour/colours`
* `Input: colours=#fe34a1`
* ![image](https://user-images.githubusercontent.com/59512495/198418623-d09d1436-f358-410f-aa45-4878e531e0cf.png)

* `Endpoint: /api/colour/words`
* `Input: words=flame`
* ![image](https://user-images.githubusercontent.com/59512495/198418750-0877dc63-e994-4fde-8a8e-1901e095c62d.png)

* `Endpoint: /api/colour/words`
* `Input: words=mango icecream`
* ![image](https://user-images.githubusercontent.com/59512495/198418894-fd3f6784-87f4-4693-9b35-8e264b5e0f32.png)


## Todo

- [x] Support basic REST APIs.
- [ ] Support Authentication with user for securing the APIs.
- [ ] Write the tests for all APIs.
- [x] Organize the code with packages
- [ ] Make docs with GoDoc
- [ ] Building a deployment process 
- [ ] Connect to a database for users
- [ ] Host the Colourminds API on own server
- [ ] Allow for rgb inputs

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

![Palette Town API Footer][footer]


[banner-short]: https://github.com/MattiasHenders/palette-town-api/blob/main/assets/banner-short.png?raw=true "Palette Town API Banner"
[banner]: https://github.com/MattiasHenders/palette-town-api/blob/main/assets/banner.png?raw=true "Palette Town API Full Banner"
[footer]: https://github.com/MattiasHenders/palette-town-api/blob/main/assets/footer.png?raw=true "Palette Town API Footer"
