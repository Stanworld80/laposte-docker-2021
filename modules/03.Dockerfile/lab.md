
# Creating a Dockerfile

## Objectives

* Build an image using a Dockerfile.
* Start the NGINX web server in a container
* Enter a running container
* A subset of the most common Dockerfile instructions

## Step 1 - Write Dockerfile

1. Create `Dockerfile`

   To create and edit a file, you can use Vim (`vi Dockerfile`) or the following syntax:
   ```
   cat > Dockerfile << EOF
   # Replace me with valid Dockerfile intructions
   EOF
   ```

   Replace the `Dockerfile` content with
   ```bash
   FROM ubuntu
   RUN apt-get update && apt-get install nginx -y && apt-get clean
   CMD ["bash"]
   ```

   [Learn Dockerfile instructions](https://docs.docker.com/engine/reference/builder/#format)

2. Verify the file with `cat Dockerfile`

   Above Dockerfile content describes a Debian image with Apache web server installed.

   - `FROM` indicates the base image for our build
   - Each `RUN` line will be executed by Docker during the build
   - Our `RUN` commands must be non-interactive. (You can‘t provide input to Docker during the build.)
   - In many cases, we will add the `-y` flag to `apt-get`.

## Step 2 - Build an image

1. We can build a Docker image by executing `docker build -t webserver .`

  The docker build command builds an image from the `Dockerfile` and a context. The context is a directory on your local filesystem (currently is `.`).

  The Docker daemon runs the instructions in the `Dockerfile` one-by-one, committing the result of each instruction to a new image if necessary, before finally outputting the ID of your new image.

2. View the image on your host:

  - `docker images`
  - `docker history webserver`

3. Optionally, scan the image for vulnerabilities (you must be logged-in to Docker Hub). 

   `docker scan --file Dockerfile ubuntu`

## Step 3 - Run a bash session

The container is not yet ready to serve HTTP request, the web server Nginx is installed but not started. The `CMD` default to bash.

1. Enter the container

   `docker run -i -t --rm --name my_webserver webserver`

2. `-i` and `-t` keep STDIN open and allocate a pseudo TTY

   combined together, they permit to enter the container with a terminal

   Use `-it` as a shortcut

3. `--rm` destroy the container on stop, when we exit bash

4. The container is attached to the `bash` process, it `pid` is `1`:

   `ps aux`

5. `pid 1` is protected, you can't kill it with `kill -9 1` 

6. Start Nginx in attached mode `nginx -g 'daemon off;'`

## Step 4 - Enter a running container

Now that NGINX is started in attached mode, we can't issue new commands from the current terminal.

1. Open a new terminal session

2. User `docker exec` to enter inside the container:

   `docker exec -it my_webserver bash`

3. Test that NGINX is listening on port 80

   ```
   apt-get install -y curl
   
   curl http://localhost
   ```

## Step 5 - Build the final image and container

1. Create an image with NGINX started in the container

   ```bash
   FROM ubuntu
   RUN apt-get update && apt-get install nginx -y && apt-get clean
   ENTRYPOINT ["nginx", "-g", "daemon off;"]
   ```

   ```bash
   docker build -t webserver .
   ```

2. Now we have an image created from `Dockerfile`, we'll run it:

   `docker run -d -p 8080:80 --name my_webserver webserver`

3. Verify the output with curl:

   `curl http://localhost:8080`

## Step 6 - Build a web application image

1. Create a home page of a web application. We need to create the `index.html` file in the `./html` directory with the following content:

  ```html
  <!DOCTYPE html>
  <html>
    <head>
      <title>This is a title</title>
    </head>
    <body>
      <p>Hello world!</p>
    </body>
  </html>
  ```

2. Enhance our existing `Dockerfile` with the following content:

  ```bash
  FROM ubuntu
  RUN apt-get update && apt-get install nginx -y && apt-get clean
  WORKDIR /var/www/html
  COPY index.html .
  ENTRYPOINT ["nginx", "-g", "daemon off;"]
  ```

  - The `WORKDIR` instruction sets the working directory for any `RUN`, `CMD`, `ENTRYPOINT`, `COPY` and `ADD` instructions that follow it in the Dockerfile.
  - The `COPY` instruction copies files or directories from the host to the filesystem of the container at the specific path.

3. Verify:

  - the `Dockerfile`: `cat Dockerfile`
  - the HTML file: `cat html/index.html`

## Step 7 - Build and run it again

1. Rebuild the image:

  `docker build -t webserver .`

2. Stop and removing the current container, or run a new container using different port and container name:

  `docker run -d -p 8081:80 --name my_webserver_1 webserver`

3. Verify the updated output with `curl`:

  `curl localhost:8081`

## You have learned

* Build a new container
* Enter a running container
* Add assets from the host machine to the container
* Expose publicly a port opened inside the container
