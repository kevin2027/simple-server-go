


# User API.
The purpose of this service is to provide an application
that is using plain go code to define an API
  

## Informations

### Version

0.0.1

## Content negotiation

### URI Schemes
  * http

### Consumes
  * application/json

### Produces
  * application/json

## All endpoints

###  users

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /users/id | [get single user](#get-single-user) | get a user by userID |
  


## Paths

### <span id="get-single-user"></span> get a user by userID (*getSingleUser*)

```
GET /users/id
```

This will show a user info

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-single-user-200) | OK |  |  | [schema](#get-single-user-200-schema) |

#### Responses


##### <span id="get-single-user-200"></span> 200
Status: OK

###### <span id="get-single-user-200-schema"></span> Schema

## Models
