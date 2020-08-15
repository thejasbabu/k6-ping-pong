# K6 Ping Pong

This is a simple ping-pong http rest service used for testing the network between two services using k6 tool

## Requirement

1. Golang
1. [K6](https://k6.io/docs/getting-started/installation)

## Set-up

1. Go get the project
    ```
    go get github.com/thejasbabu/k6-ping-pong
    ```

1. Build the project
    ```
    go build
    ```

1. Run the pong service on port 8080
    ```
    ./k6-ping-pong -pong -port 8080
    ```

1. Run the ping service on port 8081
    ```
    ./k6-ping-pong -ping -port 8081 -endpoint "http://localhost:8080/"
    ```

1. Run the run.js script to generate the load
    ```
    k6 run -e ENDPOINT="localhost:8081" run.js
    ```

Docker compose set up makes it easier

1. Run `docker-compose up` which runs all these components and binds the `run.js`