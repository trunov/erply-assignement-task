# Erply assignement task
This is a Golang project that provides two endpoints to manage Erply customers. It connects to a local PostgreSQL database and an external Erply service for user authentication and reading/writing customer data.

## Requirements
* Golang 1.19 or higher installed on your machine.
* You need Docker client to run the tests.

## Configuration
The application requires certain environment variables or flags to run. You can set them directly in your environment, or pass them as flags when running the application.

Here are the environment variables/flags:

-   `DATABASE_DSN`: Your PostgreSQL connection string.
-   `PORT`: The port for the server to run on (default: `8080`).
-   `CLIENT_CODE`: The client code for your Erply account.
-   `USERNAME`: Your Erply account username.
-   `PASSWORD`: Your Erply account password.
-   `AUTH`: The authentication method (default: `admin`).

You can run the project by using the following command:

`go run . -d=postgres://user:password@localhost:5432/db?sslmode=disable -c=clientcode -u=username -pass=password` 

## Endpoints
The application provides the following endpoints:

-   `GET /customer/{id}`: Fetch a customer. This endpoint uses caching - it tries to find the customer in the PostgreSQL database first.
-   `POST /customer`: Add a new customer.
The body of the request should be a JSON object with the following fields:
```json
{
    "firstName": "string",
    "lastName": "string",
    "email": "string",
    "phone": "string",
    "twitterID": "string"
}
```
Where:
-   `firstName` (required): The first name of the customer.
-   `lastName` (required): The last name of the customer.
-   `email` (required): The email of the customer.
-   `phone`: The phone number of the customer.
-   `twitterID`: The Twitter ID of the customer.

## Middleware
There are two middleware functions that are used in the application:

-   `VerifyErplyUser`: This middleware function authenticates the user with the Erply service, and stores the session key in a cookie.
-   `TokenAuthorization`: This middleware function checks the bearer token in the Authorization header.

## Running tests
You can run the tests by using the following command:

`go test ./...` 

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change. Please make sure to update tests as appropriate.