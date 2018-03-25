# City Travel Api

The City Travel Api

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

* Docker >= v17.05
* go >= 1.10

### Installing

```
git clone https://gitlab.com/andreluizmachado/go-challenge-ac001.git
```

```
cd go-challenge-ac001
```

```
./development.sh
```

### Running the binary

```
./binary
```

The service listen the 3000 port

## API Resources

### POST /city
Resource to create a city

Request:
```shell
POST http://localhost:3000/city
Header Content-Type: application/json

Content-Type: application/json
{
    "name": "Jandira",
    "borders": [2,3]
}
```

Response:
```shell
Header Location /city/1
201 Created

{
    "id": 1,
    "name": "Jandira",
    "borders": [
        2,
        3
    ]
}
```

### GET /city/:id
Resource to take a city

Request:
```shell
GET /city/1
Header Content-Type: application/json
```

Response
```shell
200 OK
{
    "id": 1,
    "name": "Jandira",
    "borders": [
        3
    ]
}
```

Response
```shell
404 Not Found
```

### PUT /city/:id
Resource to update a city

Request:
```shell
PUT /city/3
Header Content-Type: application/json
```

Response
```shell
200 OK
{
    "id": 3,
    "name": "Jandira",
    "borders": [
        4
    ]
}
```

Response
```shell
404 Not Found
```

### DELETE /city/:id
Resource to delete a city

Request:
```shell
DELETE /city/3
Header Content-Type: application/json
```

Response
```shell
200 OK
```

Response
```shell
404 Not Found
```

### DELETE /cities
Resource to delete all cities

Request:
```shell
DELETE /cities
Header Content-Type: application/json
```

Response
```shell
200 OK
```

### GET /cities
Resource to get all cities

Request:
```shell
GET /cities
Header Content-Type: application/json
```

Response
```shell
200 OK
{
    "cities": [
        {
            "id": 1,
            "name": "Jandira",
            "borders": [
                3
            ]
        }
    ]
}
```



### GET /city/:oirigin/travel/:destinate
Resource to get a valid path between origin and destinate

Request:
```shell
GET /city/2/travel/1
Header Content-Type: application/json
```

Response
```shell
200 OK
{
    "path": [
        2,
        4,
        3,
        1
    ]
}
```

Response
```shell
404 Not Found
```

## Authors

[André Luiz Machado](https://github.com/andreluizmachado)

