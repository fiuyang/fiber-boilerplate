# Description
A Boilerplate Api Golang

# Features
This Boilerplate Api Golang Project with Go, Fiber, Gorm, JWT, Postgresql, Swagger
 
# Tech Used
 ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![JWT](https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens) ![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white) ![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
      
# Getting Start:
Before you running the program, make sure you've run this command:
```bash
 go get -u all
 development.env with config db
 docker compose up -d
```

### Run the program
```bash
 go run main.go
```

### Re-Init Docs Swagger
```bash
 swag init
```

### Create Table Sql
```bash
 migrate create -ext sql -dir migrations create_table_{name_table}
```

### Check Docs Swagger
```bash
 http://localhost:3000/docs/index.html#/
```

The program will run on http://localhost:3000 
<!-- </> with 💛 by readMD (https://readmd.itsvg.in) -->