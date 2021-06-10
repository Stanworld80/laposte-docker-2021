
# Docker essentials

## Objectives 

1. Explore Docker CLI
2. Run and manage containers
3. Run a command inside a container

## Step 1: Explore Docker CLI

1. List available commands, either run `docker` with no parameters or execute `docker help`.

2. Find a command to display Docker version information.   
  > Tip. You can use piping to search the required command faster, for example, `docker help | grep 'information'`

3. Get the description of a specific command using `docker <COMMAND> --help` and learn available options. Find, for example, how to show all containers (not only running) using the command `docker ps`.

## Step 2 - Run and manage containers

1. Run the Redis database container using `docker run -d redis`

  What does the `-d` option? Try to run the same container without this option (press `Ctrl + C` to terminate the blocking process).

2. Find the following commands and explore the Docker host by reading and understanding the output:

  - list Docker images
  - list only running containers
  - list all containers

3. Manage containers:

  - Stop a running container
  - Start a stopped container
  - Restart a container
  - Kill a container (learn the difference from "stop" container)
  - Remove a stopped/killed container
  - Remove a running container (using the "force" mode)

4. Fetch the logs of a container

## Step 3 - Run a command inside a container

1. Run the Alpine Linux container using `docker run alpine` 

  `alpine` is a smallest Docker image based on Alpine Linux with only 5 MB in size!

2. List running containers.

  - Why `alpine` is not on this list?

3. Run the `alpine` container using the following command: `docker run -it alpine /bin/sh`.

  - What do the `-i` and `-t` options? 
  - What does the `/bin/sh` command mean?
  - Open another terminal window and make sure the `alpine` container is currently running.
  > Tip. To terminate shell inside a container print `exit`.

4. Run a Redis container with `docker run -d redis` and execute a command inside it using `docker exec -it <CONTAINER> bash`

  - What is the difference between `docker run -it <IMAGE> <COMMAND>` and `docker exec -it <CONTAINER> <COMMAND>`?
  - Explore the Linux filesystem inside the running containers with commands such as `pwd`, `ls`, `cd`.
  - Try to communicate with Redis: run Redis client with `redis-cli` and print `PING`. It should answer with the `PONG` message indicating that the Redis server is running.

## You have learned

1. The basic Docker CLI commands 
2. How to manage containers
3. How to run commands inside a container

