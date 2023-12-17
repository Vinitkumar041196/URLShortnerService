# URL Shortener

- This service is a REST based API server that exposes APIs to shorten any URL with guarantee that same url will always generate same shortened URL.
- It also redirects the shortened URL to orginal URL with its in memory URL store.
- You can also get the top N shortened domains using its metrics API.

## To use service:

1. For local environment with docker installed, you can run the service using docker compose.
   ```
   docker-compose up -d
   ```

2. For non docker environment without docker, you would need to install go and run the below command 
    ```
    go run main.go
    ```

3. By default the server will start at port **8080**. To change this you can update the **SERVER_ADDR**
    - for docker env, update the docker-compose file 
        ```
          - SERVER_ADDR=localhost:8090
        ```
    - for non docker env, export the env variable directly
        ```
        export SERVER_ADDR=localhost:8090
        ```

4. The Swagger documentation of the service will be available at **{HOST}/docs/index.html** e.g. *http://localhost:8080/docs/index.html*
   
5. To run unit test, run the below command from root folder of the project
    ```
    go test ./... -v
    ```
6. To enable TLS you would need to set up few variables as mentioned below
    ```
      - ENABLE_TLS=true
      - SSL_CRT_PATH="path to ssl certificate file"
      - SSL_KEY_PATH="path to ssl key file"
    ```

## Docker image build

The dockerfile is available in *deployment* folder.

To build the docker image run the below command from root folder of the project
    ```
    docker build . -f ./deployment/dockerfile -t url_shortener
    ```

An already built image of this project is available at:
**https://hub.docker.com/repository/docker/vinitondocker/url_shortener/general**