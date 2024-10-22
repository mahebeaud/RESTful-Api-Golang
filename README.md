# RESTful API Golang

## Description

A RESTful API built in Golang with basic authentication features such as login, register, and logout. This backend is designed to support additional routes and extend the logic for any application you want to create. This project uses the [Gin framework](https://gin-gonic.com/) to build the API.

## Requirements

- [Golang](https://go.dev/doc/install)
- Make command
  ```bash
  sudo apt install make
  ```

**Development environment:**
- You need [air](https://github.com/cosmtrek/air) for hot reloading during golang development. It can also be installed using the project's Makefile.

## Build

You can modify the configuration of the Air command in the **.air.toml** file located at the root of the project.

**Development:**
```bash
make all
```

**Production:**
Use Docker Compose to build the production environment. Remember to add all your environment variables in the Docker Compose file, along with your GitHub Actions workflow or any other CI/CD tool.
```bash
docker compose up -d
```

### Author

- Mahé BEAUD

[![LinkedIn](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/mahe-beaud/?locale=en_US)

