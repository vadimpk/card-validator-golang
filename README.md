# Card Validation Service

Simple REST API application in Golang for validation of credit cards

The application's entry point is `main.go` file. Configuration options can be found in `config/config.go` file. 


### Run the server

Locally:
```
make run
```

Docker:
```
docker-compose up
```

The `.env` file can be used to run the server.


### Examples
```
curl --location 'localhost:8080/api/validate' \
--header 'Content-Type: application/json' \
--data '{
    "number": "4242424242424242",
    "expMonth": "02",
    "expYear": "2025"
}'
```


### Endpoints

- `api/validate` – validates real-life cards.
- `test/validate` – validates test cards.


### Architecture

All the logic of the application is stored inside `intenal/` directory. `pkg/` directory stores additional data that does not relate to the project directly, like `logger` package. Inside the `internal/` directory there are the following packages:
- `domain` – stores the domain logic of the application, responsible for the actual card validation with different rules.
- `service` – presents an interface with business logic of the whole application.
- `handlers` – is actually a directory that is intented to store different packages that can handle incoming requests and bound them to service methods.
  - `http` – package for handling http requests.

In the domain logic there is an interface `CardValidator` that has a single method that validates a card. The intent of this interface is to allow as many implementations of card validation as possible (Strategy pattern). This helps code extensibility. In the application there are many different concrete card validators, each validating a specific criteria of the card. Some validators encapsulate others to broad their own functionality. Also, there is a factory that helps managing different implementations of card validator.

The diagram represents a conceptual overview of the project architecture. There are handlers on the top that communicate with service and domain logic below.

<img width="610" alt="Screenshot 2024-02-27 at 15 50 28" src="https://github.com/vadimpk/card-validator-golang/assets/65962115/44c3b719-f98d-4782-87ae-8cd5338b0c6d">
