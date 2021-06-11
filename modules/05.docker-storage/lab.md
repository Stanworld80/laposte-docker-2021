# Docker storage

## Using volumes

1. Create a volume

   ```
   docker volume create my-vol
   ```

2. View the volume

   ```
   docker volume ls
   docker volume inspect my-vol
   ```

3. Start a new container, mount the volume with `--mount`, write inside of it

   ```
   docker run \
     -it --rm \
     --mount source=my-vol,target=/app \
     alpine:latest \
     sh -c "echo 'pong' > /app/ping"
   ```

4. Or using the `-v` flag

   ```
   docker run \
     -it --rm \
     -v my-vol:/app \
     alpine:latest \
     sh -c "echo 'pong again' > /app/ping"
   ```

5. Read the volume from another container

   ```
   docker run \
     -it --rm \
     --mount source=my-vol,target=/app \
     alpine:latest \
     cat /app/ping
   ```

6. Using an alternative image won't affect the output.

   ```
   docker run \
     -it --rm \
     --mount source=my-vol,target=/app \
     ubuntu:latest \
     cat /app/ping
   ```

## Using bind mount

1. Create a local directory

   ```
   mkdir /my-bind-mount
   ```

   

2. Mount the local directory in the container and write to it

   ```
   docker run \
     -it --rm \
     -v `pwd`/my-bind-mount:/app \
     alpine:latest \
     sh -c "echo 'pong' > /app/ping"
   ```

3. Check that a new file was created on the host machine

   ```
   cat ./my-bind-mount/ping
   ```

   