Golang RestFul API to do CRUD to mySQL database using MVC pattern.

List of API :
Header : X-Public -> true || false

GET     /users/:user_id
GET     /internal/users/search?status=active
POST    /users
PUT     /users/:user_id
DELETE  /users/:user_id


note : the Header X-Public represents access token to determine response payload
