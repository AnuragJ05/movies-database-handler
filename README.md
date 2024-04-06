
# Movies-Database-Handler

This project is a simple API for managing a movies database, built using Go and the Gorilla Mux library. It provides two endpoints for creating and retrieving movies.

![Golang SDE Assignment-2](https://github.com/AnuragJ05/movies-database-handler/assets/46484628/d31d446c-3d0b-4258-8e84-872bbafa52db)


## Features

```
RESTful API with two endpoints:
    - POST /movies: Create a new movie.
    - GET /movies: Retrieve all movies.
Uses PostgreSQL as the database.
Utilizes Gorilla Mux for routing.
```

## Prerequisites

Go (1.21+)

Docker and Docker Compose

PostgreSQL (installed via Docker Compose)

## Setup

1. Clone the repository:

```
git clone https://github.com/AnuragJ05/movies-database-handler.git
```

2. Navigate to the project directory:

```
cd movies-database-handler
```

3. Build and start the PostgreSQL container and go app using Docker Compose:
```
docker-compose up --build
```


## Usage

### Create a movie

To create a new movie, send a POST request to the /movies endpoint with a JSON payload containing the movie details:

```
curl -X POST -d '{"isbn":"ISBN 0-061-96436-0", "title":"RRR", "director": "S. S. Rajamouli"}' -H "Content-Type: application/json" http://localhost:5000/movies

curl -X POST -d '{"isbn":"ISBN 0-011-12436-1", "title":"SOTY", "director": "karan johar"}' -H "Content-Type: application/json" http://localhost:5000/movies

curl -X POST -d '{"isbn":"ISBN 1-011-22431-2", "title":"PK", "director": "Rajkumar Hirani"}' -H "Content-Type: application/json" http://localhost:5000/movies
```

### Retrieve movies

To retrieve all movies, send a GET request to the /movies endpoint:

```
curl -X GET http://localhost:5000/movies
```

## Check the Data updated in the database

```
Exec inside contanier
    docker exec -it <container id> sh

Interact with PostgreSQL databases via the command line
    psql -h localhost -p 5432 -U postgres --dbname=postgres

List the Database
    \l

List the Table
    \dt

List the data from table
    select * from movies;
```
