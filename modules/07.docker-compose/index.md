
# Docker Compose

- is a tool for defining and running multi-container Docker applications
- Uses a compose (configuration) YAML file   
- A way to document and configure all of the application’s service dependencies (databases, queues, caches, web service APIs, etc)
- Only one single command to create and start containers - `docker-compose up`

## Use cases

1. **Development environments**   
  When you’re developing software, the ability to run an application in an isolated environment and interact with it is crucial. The Compose command-line tool can be used to create the environment and interact with it.

2. **Automated testing environments**   
  An important part of any Continuous Deployment or Continuous Integration process is the automated test suite. Automated end-to-end testing requires an environment in which to run tests. Docker Compose provides a convenient way to create and destroy isolated testing environments for your test suite.

3. **Single host deployments**   
  To deploy to a remote Docker Engine.

## Using Docker Compose

A three-step process:

1. Define your app’s environment with a `Dockerfile` so it can be reproduced anywhere.
2. Define the services that make up your app in `docker-compose.yml` so they can be run together in an isolated environment.
3. Run `docker-compose up` and Compose starts and runs your entire app.

## Structure of `docker-compose.yml`

- The `version` of the compose file
- The `services` which will be built
- All used `volumes`
- The `networks` which connect the different services

[Read more](https://docs.docker.com/compose/compose-file/compose-file-v3/)

## Example: WordPress website

The `docker-compose.yml` file contains:

```yaml
version: '3.3'

services:
   db:
     image: mysql:5.7
     volumes:
       - db_data:/var/lib/mysql
     networks:
       - backend
     environment:
       MYSQL_ROOT_PASSWORD: somewordpress
       MYSQL_DATABASE: wordpress
       MYSQL_USER: wordpress
       MYSQL_PASSWORD: wordpress

   wordpress:
     depends_on:
       - db
     image: wordpress:latest
     ports:
       - "8000:80" # host:container
     networks:
       - backend
     environment:
       WORDPRESS_DB_HOST: db:3306
       WORDPRESS_DB_USER: wordpress
       WORDPRESS_DB_PASSWORD: wordpress
       WORDPRESS_DB_NAME: wordpress
volumes:
    db_data: {}
networks:
  backend:
    driver: bridge
```

## Services configuration in Docker Compose

TODO...

Top-level keys:
- `build`
- `deploy`
- `depends_on`
- `networks`

[Read more](https://docs.docker.com/compose/compose-file/compose-file-v3/#service-configuration-reference)

## Docker Compose commands

- `docker-compose up` - Create and start containers
- `docker-compose down` - Stop and remove containers, networks, images, and volumes
- `docker-compose start` - Start services
- `docker-compose stop` - Stop services
- `docker-compose exec` - Execute a command in a running container
- `docker-compose rm` - Remove stopped containers
- `docker-compose scale` - Set number of containers for a service
- ...
