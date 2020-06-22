# Task REST API

### Specifications
* Docker 19.03
* Go 1.13 (docker image)
* Postgres 10 (docker image)
* Postman 7.6.0 

### Run

``` docker-compose up ```

### Golang container doesn't want to start?
Change line endings of the `build.sh` file from CRLF to LF


### API Routes:

GET `http://localhost:8080/`

POST `http://localhost:8080/v1/login` 

POST `http://localhost:8080/v1/register?username=johndoe&password=123`

POST `http://localhost:8080/v1/login?username=johndoe&password=123`

POST `http://localhost:8080/v1/todo/create` Authorization: Bearer token 

Body: `{
“title”:”Title task 1",
“description”:”Description task 1"
}`

GET `http://localhost:8080/v1/todo/all` 
Authorization: Bearer token

GET `http://localhost:8080/v1/todo/get/<id>`
Authorization: Bearer token

PUT `http://localhost:8080/v1/todo/update/<id>`
Authorization: Bearer token

Body: `{
“title”:”Task title updated!”,
“description”:”Task description updated!”
}`

DELETE `http://localhost:8080/v1/todo/delete/<id>`
Authorization: Bearer token
