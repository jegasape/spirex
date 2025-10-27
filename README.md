# Spirex

This tool allows you to easily monitor the S&P 500 and the Dominican Republic Stock Exchange.

--- 
>[!IMPORTANT]
> To start testing or working on the project, simply run `make up`. This will get the containers running so you can begin exploring or testing the application. Other commands can be found in the `Makefile`. Let me know if you need further adjustments!
---

## Requirements

- Go (version go1.24.5 linux/amd64)
- Docker (for `up` and `down` commands)

## Available Commands

### `make run`
Runs the application using `go run`.

### `make build`
Builds the application for Linux (amd64).

### `make up`
Starts the Docker containers using `docker-compose`. This is the recommended first step to get the environment up and running.

### `make down`
Stops and removes the Docker containers and cleans up unused resources.
