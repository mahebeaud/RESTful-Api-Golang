# RESTful API in Golang

## Description

A RESTful API built in Golang with basic authentication features such as login, register, and logout. This backend is designed to support additional routes and extend the logic for any application you want to create. This project uses the [Gin framework](https://gin-gonic.com/) to build the API.

## Table of Contents

- [Description](#description)
- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Development](#development)
- [Build](#build)
- [Author](#author)
- [Acknowledgements](#acknowledgements)

## Features

- User authentication (login, register, logout)
- JWT token-based authentication
- Basic CRUD operations
- Extendable architecture

## Requirements

- [Golang](https://go.dev/doc/install)
- Make command
  ```bash
  sudo apt install make
  ```
- [air](https://github.com/cosmtrek/air) for hot reloading during development

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/mahebeaud/RESTful-Api-Golang.git
    cd RESTful-Api-Golang
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

    3. Install `air` for hot reloading or directly use the Makefile for installation:
    ```bash
    curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
    ```

## Usage

1. Start the development server:
    ```bash
    make all
    ```

2. Access the API at `http://localhost:8080`

## Development

You can modify the configuration of the Air command in the **.air.toml** file located at the root of the project.

## Build

**Development:**
```bash
make all
```

**Production:**
Use Docker Compose to build the production environment. Remember to add all your environment variables in the Docker Compose file, along with your GitHub Actions workflow or any other CI/CD tool.
```bash
docker compose up -d
```

## Author

- Mahé BEAUD

[![LinkedIn](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/mahe-beaud/?locale=en_US)

## Acknowledgements

Special thanks to the [Gin framework](https://gin-gonic.com/) team for their excellent work on the framework used in this project.