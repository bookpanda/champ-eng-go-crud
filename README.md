# ChAMP SWE Track Backend Assignment

## Installation
- git clone https://github.com/bookpanda/champ-eng-go-crud.git
- Create .env from .env.template, no need to change values of variables
- go install
- go install github.com/cosmtrek/air@latest
- docker compose up
- go run .\migrate\migrate.go
- Connect database management tool to database docker
    - Host: localhost
    - Port: 5432
    - User: test
    - Password: 1234
    - Database: db
- air main.go (please run this command in Linux/WSL)

## API Documentation
You can view the API Docs at localhost:3000/swagger/index.html
