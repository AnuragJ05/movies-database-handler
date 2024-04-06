
# Movies-Database-Handler

This project is a simple API for managing a movies database, built using Go and the Gorilla Mux library. It provides two endpoints for creating and retrieving movies.


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
git clone https://github.com/your-username/movies-database-handler.git
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
curl -X POST -d '{"isbn":"ISBN 0-061-96436-0", "title":"marvel movie", "director": "anurag jain"}' -H "Content-Type: application/json" http://10.239.40.166:5000/movies
```

### Retrieve movies

To retrieve all movies, send a GET request to the /movies endpoint:

```
curl -X GET http://10.239.40.166:5000/movies
```

