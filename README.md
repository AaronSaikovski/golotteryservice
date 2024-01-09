<div align="center">

# golotteryservice

A WebApi using the Go Fiber webframework to generate lottery numbers for use in various lottery games - written in 100% Go.

[![Build Status](https://github.com/AaronSaikovski/golotteryservice](https://github.com/AaronSaikovski/golotteryservice)/workflows/build/badge.svg)](https://github.com/AaronSaikovski/golotteryservice/actions)
[![Coverage Status](https://coveralls.io/repos/github/AaronSaikovski/golotteryservice/badge.svg?branch=main)](https://coveralls.io/github/AaronSaikovski/golotteryservice?branch=main)
[![Licence](https://img.shields.io/github/license/AaronSaikovski/golotteryservice)](LICENSE)

</div>

## This project contains:

- Fiber Web framework - https://github.com/gofiber/fiber
- Zerolog - https://github.com/rs/zerolog
- Swaggo - https://github.com/swaggo/swag
- CI (Github Actions)
- Unit tests (coming soon)
- Container support with [docker](Dockerfile) and [docker-compose](docker-compose.yml)
- API key validation - https://docs.gofiber.io/api/middleware/keyauth/

## Installation

The toolchain is mainly driven by the Makefile.

```bash
help          - Display help about make targets for this Makefile
localrelease  - Builds the project in preparation for (local)release
docs          - updates the swagger docs
debug         - Builds the project in preparation for debug
buildandrun   - builds and runs the program on the target platform
run           - runs main.go for testing
clean         - Remove the old builds and any debug information
unittest      - executes unit tests
dep           - fetches any external dependencies
vet           - Vet examines Go source code and reports suspicious constructs
staticcheck   - Runs static code analyzer staticcheck - currently broken
seccheck      - Code vulnerability check
lint          - format code and tidy modules
depupdate     - Update dependencies
```

## Setup/Configuration

This API is secured via an API key which is set in the .env file and is passed in the request header to the API as `XApiKey`. Make sure to not commit the key to your public repository.
Refer to the file `.env.example` for help. You will need to set this prior to running the API, otherwise you will get this error. `missing or malformed API Key`.

## Usage

From the command line the usage is pretty simple:
All values are integers unless specified.

```bash
make run
curl http://localhost:8080/lottery?maxgames='<Max games value>'&randnum='<Random number value>;&maxnumspergame='<Max numbers per game value>'

```

Results will be returned as JSON

```json
[
  "Game: 1 - [87 90 56 7 37 4 71]",
  "Game: 2 - [90 88 47 84 45 23 32]",
  "Game: 3 - [76 26 38 43 18 1 30]",
  "Game: 4 - [28 90 64 81 60 44 76]",
  "Game: 5 - [71 58 30 13 19 20 40]",
  "Game: 6 - [83 51 87 3 7 19 88]",
  "Game: 7 - [76 26 32 39 79 53 66]",
  "Game: 8 - [85 61 32 17 34 21 48]",
  "Game: 9 - [55 70 9 59 86 45 54]",
  "Game: 10 - [36 32 7 2 16 73 58]"
]
```

### Docker Support

This project has been fully tested with Docker and includes a `Dockerfile` and a `docker-compose.yml` file and uses a multi-stage docker file to reduce the docker image to approx. 30MB.
The docker build process also optimises the Go executable down to a bare minimum.
To build the docker image type:

```
docker build -t golotteryservice:1.0.0 .
```

and to run the comtainer type:

```
docker run -p 8080:8080 golotteryservice:1.0.0
```

or you can streamline the container build and run process by typing:

```
docker compose up
```

and to tear it down type:

```
docker compose down
```

## Issues

Please feel free to lodge an [issue or pull request on GitHub](https://github.com/AaronSaikovski/golotteryservice/issues).

## Known Issues/Bugs

- None at this stage
