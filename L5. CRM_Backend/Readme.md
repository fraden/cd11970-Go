# CRM Backend
This is the final project for the Udacity GoLang course.

## Installation
Run
```go mod download```
to install all dependencies, defined in `go.mod`file.

## Launch
To execute the application, run
```go run main.go```

## Usage
There are multiple APIs, that can be used:

* GET "/customers" to get all customers 
* GET "/customers/{id}" to get a specific customer, defined by id
* PUT "/customers/{id}" to update a specific customer, defined by id
* DELETE "/customers/{id}" to delete a specific customer, defined by id
* POST "/customers" to create a new customer. You need to provide the customer data as json body, that looks like the following:
```
{
    "name": "Max Mustermann",
    "role": "solution architect",
    "email": "max.mustermann@coolmail.com",
    "phone": "0171/22446688",
    "contacted": false
}
```