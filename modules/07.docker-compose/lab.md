
# Using Docker Compose

## Objectives

* Use and connect 2 containers
* Migrate their Docker definitions inside Docker Compose
* Interact with Docker Compose
* Use some best-practices

## Step 1 - a basic Go web server

* If go is not yet present, install it:
  `command -v go || sudo apt install -y golang-go`
  
* Create a new directory `step_1`.

* Import the file `step_1/app.go` inside it.

* From the `step_1` directory

  * Compile the application

    `go build app.go`

  * Run the application

    `./app`

* From another terminal, test the application and validate its output:

  ```
  curl http://localhost:8080/pong/step_1
  Message received: pong/step_1
  ```

* Kill the running web server

  `kill $(ps aux | grep '[a]pp' | awk '{print $2}')`

## Step 2 - wrap the application inside a container

* Create a new directory `step_2` and import the same `app.go` code, also present in `step_2/app.go`.

* Declare a new `Dockerfile` using a multi-stage build.

  * The first part install the `go` environment and compile the application.

    ```dockerfile
    FROM golang:1.16 AS builder
    WORKDIR /go/src/github.com/alexellis/href-counter/
    RUN go get -d -v golang.org/x/net/html  
    COPY app.go    .
    RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app app.go
    ```

  * The second part import the generated Go application and execute it.

    ```dockerfile
    FROM alpine:latest  
    RUN apk --no-cache add ca-certificates
    WORKDIR /root/
    COPY --from=builder /go/src/github.com/alexellis/href-counter/app .
    CMD ["./app"] 
    ```

* Build the image, it is named `ping` with version `step_2`:

  ```bash
  docker build -t ping:step_2 .
  ```

* Start a container from this image:

  ```bash
  docker run \
  	-p 8080:8080 \
    -d \
  	--name ping_app \
  	ping:step_2
  ```

  * The port `8080` is exposed
  * It is detached from the terminal
  * It is named `ping_app`

* Test the application:

  ```bash
  curl http://localhost:8080/pong/step_2
  Message received: pong/step_2
  ```

* Stop and remove the container

  ```bash
  docker rm -f ping_app
  ```

* Make some change, for example, update the message to `Got new message:`

  ```bash
  docker build -t ping:step_2 .
  docker run \
  	-p 8080:8080 \
    -d \
  	--name ping_app \
  	ping:step_2
  curl http://localhost:8080/pong/step_2
  docker rm -f ping_app
  ```

## Step 3 - connect the application to another container

* Create a new directory `step_3`.

* Start the MariaDB database inside a container

  ```bash
  docker run \
    -p 127.0.0.1:3306:3306 \
    -e MYSQL_ROOT_PASSWORD=my-secret-pw \
    -d \
    --name ping-mariadb \
    mariadb:latest
  ```

* Export it IP address and validate the database, we also declare the user password:

  ```bash
  export MYSQL_HOST=`docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' ping-mariadb`
  export MYSQL_PASSWORD='my-secret-pw'
  
  echo $MYSQL_HOST
  echo $MYSQL_PASSWORD
  
  docker run \
    -it \
    --rm \
    mariadb \
    mysql \
      -h$MYSQL_HOST \
      -p3306 \
      -uroot \
      -p$MYSQL_PASSWORD \
      -e "SHOW DATABASES;"
  ```

  

* Import the following files:

  * `step_3/main.go`

    The application is getting more complex, it initializes the database and insert messages into it.

  * `step_3/go.mod`

    Contains the list of dependency module versions.

  * `step_3/go.sum`

    Contains the checksums of the content of specific module versions.

* Looking at the `main.go` file, the database connection properties are obtained from environment variables and use default when appropriate:

  ```bash
  cat app.go | grep getEnv
  ```

* Import `step_3/Dockerfile`, it is almost identical with the addition of 2 `COPY` instructions

* Build and run the container:

  ```bash
  docker build -t ping_app:step_3 .
  docker run \
    -p 8080:8080 \
    -d \
    -e MYSQL_HOST=$MYSQL_HOST \
    -e MYSQL_PASSWORD \
    --name ping_app \
    ping_app:step_3
  ```

* Test the application:

  ```bash
  curl http://localhost:8080/pong/step_3/message_1
  curl http://localhost:8080/pong/step_3/message_2
  docker run \
    -it \
    --rm \
    mariadb \
    mysql \
      -h$MYSQL_HOST \
      -p3306 \
      -uroot \
      -p$MYSQL_PASSWORD \
      -e "SELECT * FROM ping.history;"
  ```

* Stop and remove the container and remove the environment variables, we keep the database running:

  ```bash
  docker rm -f ping_app
  unset MYSQL_HOST
  unset MYSQL_PASSWORD
  ```

## Step 4 - storing environmental variables inside `.env`

* Create an `env.sample` file, it will be committed and serve as an example:

  ```bash
  # Example configuration file,
  # use it as a source of inspiration
  # by moving and modifying its content
  # to the `.env` location.
  
  # IP or domain of the MariaDB server
  # MUST BE REPLACED by the MariaDB container IP
  # for example `docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' ping-mariadb`
  MYSQL_HOST=127.0.0.1
  
  # Password of the MariaDB user
  MYSQL_PASSWORD=my-secret-pw
  ```

* Create a local `.env` file, ignored from GIT:

  ```bash
  echo -n '' > .env
  echo "MYSQL_HOST=`docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' ping-mariadb`" >> .env
  echo "MYSQL_PASSWORD=my-secret-pw" >> .env
  cat .env
  ```

* Import the files `step_4/app.go`, `step_4/go.mod`, `step_4/go.sum` and `step_4/Dockerfile`, they remains identical to `step_3`.

* Build and run the container, now refering to the `.env` file:

  ```bash
  docker build -t ping_app:step_4 .
  docker run \
    -p 8080:8080 \
    -d \
    --env-file .env \
    --name ping_app \
    ping_app:step_4
  ```

* Test the application:

  ```bash
  curl http://localhost:8080/pong/step_4/message_1
  curl http://localhost:8080/pong/step_4/message_2
  docker run \
    -it \
    --rm \
    mariadb \
    mysql \
      -h$MYSQL_HOST \
      -p3306 \
      -uroot \
      -p$MYSQL_PASSWORD \
      -e "SELECT * FROM ping.history;"
  ```

* Stop and remove the containers, including the database container:

  ```bash
  docker rm -f ping_app ping-mariadb
  ```

## Step 5 - Use Docker Compose

* Import the files `app.go`, `go.mod`, `go.sum` and `Dockerfile`. They remain identical to `step_4`.

* The `.env` file is different:

  ```bash
  cat >.env <<ENV
  MYSQL_HOST=step_5_db_1
  MYSQL_PASSWORD=my-secret-pw
  MYSQL_ROOT_PASSWORD=my-secret-pw
  ENV
  ```

  * The MariaDB server is now accessible by hostname: `step_5_db_1`
  * The MariaDB image expect the `MYSQL_ROOT_PASSWORD` which could be different from the user password if we connect with another user than `root` (recommanded).

* Create the `docker-compose.yml` file:

  ```yml
  services:
    web:
      build: .
      ports:
        - "8080:8080"
      env_file:
        - ".env"
    db:
      image: 'mariadb'
      env_file:
        - ".env"
  ```

* Install Docker Compose if not present:

  ```bash
  command -v docker-compose || \
  	sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose && \
  	sudo chmod +x /usr/local/bin/docker-compose
  ```

* Start the containers

  ```bash
  docker-compose up
  ```

* Chances are that the `web` container crashes because the database is not yet listening on port `3306`

  * The `ps` command print an exit code `2`:

    ```
    docker-compose ps
        Name                 Command             State     Ports  
    --------------------------------------------------------------
    step_5_db_1    docker-entrypoint.sh mysqld   Up       3306/tcp
    step_5_web_1   ./app                         Exit 2
    
  * The `logs` command display more details:

    ```bash
    docker-compose logs web
    Attaching to step_5_web_1
    web_1  | panic: dial tcp 172.18.0.2:3306: connect: connection refused
    ...
    ```

  * Start the `web` container:

    ```bash
    docker-compose start web
    docker-compose ps
        Name                 Command             State                    Ports                  
    ---------------------------------------------------------------------------------------------
    step_5_db_1    docker-entrypoint.sh mysqld   Up      3306/tcp                                
    step_5_web_1   ./app                         Up      0.0.0.0:8080->8080/tcp,:::8080->8080/tcp
    ```



* Test the application:

  ```bash
  curl http://localhost:8080/pong/step_5/message_1
  curl http://localhost:8080/pong/step_5/message_2
  docker run \
    -it \
    --rm \
    mariadb \
    mysql \
      -hstep_5_db_1 \
      -p3306 \
      -uroot \
      -pmy-secret-pw \
      -e "SELECT * FROM ping.history;"
  ```
  
* Stop and remove the containers:

  ```bash
  docker-compose stop
  docker-compose rm
  ```

## Step 6 - control the startup order and persist data

* Import the files `app.go`, `go.mod`, `go.sum` and `Dockerfile`. They remain identical to `step_4`.

* Modify the `docker-compose.yml` declaration:

  ```yml
  services:
    web:
      build: .
      depends_on:
        - db
      ports:
        - "8080:8080"
      env_file:
        - ".env"
    db:
      image: 'mariadb'
      env_file:
        - ".env"
      volumes:
        - db_data:/var/lib/mysql
  
  volumes:
      db_data: {}
  ```

  * A volume `db_data` persist the MariaDB database.
  * Container `web` depends on container `db`.
  * Container `db` is reaching with the hostname `db`.

* Modify the `.env` file to reflect the new hostname of container `db`.

  ```bash
  cat >.env <<ENV
  MYSQL_HOST=db
  MYSQL_PASSWORD=my-secret-pw
  MYSQL_ROOT_PASSWORD=my-secret-pw
  ENV
  ```

* Build and start the containers:

  ```bash
  docker-compose build
  docker-compose create
  docker-compose start
  ```

* Check the volume creation:

  ```bash
  docker volume ls | grep step_6
  ```

* Insert some data:

  ```bash
  curl http://localhost:8080/pong/step_6/message_1
  curl http://localhost:8080/pong/step_6/message_2
  ```

* Destroy the container and restart them:

  ```
  docker-compose stop
  docker-compose rm
  docker-compose ps
  docker-compose create
  docker-compose start
  docker-compose ps
  ```

* Insert new data and validate that previous data where persisted:

  ```bash
  curl http://localhost:8080/pong/step_6/message_3
  curl http://localhost:8080/pong/step_6/message_4
  docker-compose exec \
    db \
    mysql \
      -hlocalhost \
      -p3306 \
      -uroot \
      -pmy-secret-pw \
      -e "SELECT * FROM ping.history;"
  ```

* Stop and remove all resources:

  ```bash
  docker-compose stop
  docker-compose rm
  docker volume rm step_6_db_data
  ```

  

## Step 7 - further improvment

* Reducing service duplication in `docker-compose.yml` with aliases and anchors, eg by sharing the `.env` declaration.
* Introducing a health check in the web service.
* Create a `docker-compose.override.yml` and intialize the database in the starter script or based on the presence of an environment variable whose value differ between the production and development environments.
* ...

