# Simple Go REST API.
A simple Go REST API web application for learning how to build a web api using [Golang](https://go.dev/) and [Gin Web Framework](https://gin-gonic.com/docs/).

## What the Web API Does.
The Web API exposes cability to act on notes. Notes are entities that track actions that a user is interested in. A note can have comments by other users (similar to posts). The notes can also have reactions. A user can invite other users to either: edit, view and even delete the notes.  

### Technologies Used
- [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/) for containerizing.
- Postgres image for persistance.
- OAuth for auth - ?? maybe ??

## Getting Started (Developer)
This section covers what devs need to do to get started.

## Usual Tools:
Running the postgres image only from the docker compose:  
```shell
sudo docker-compose -f docker-compose.yml up  --build 'postgres'
```  

Executing the running Postgres container then query the postgres command line:  
```shell
# get postgres container id
docker ps 

# result of docker ps:
CONTAINER ID   IMAGE          COMMAND                  CREATED      STATUS         PORTS                                       NAMES
4112657f0f8f   src_postgres   "docker-entrypoint.s…"   2 days ago   Up 9 minutes   0.0.0.0:5432->5432/tcp, :::5432->5432/tcp   src_postgres_1

# take CONTAINER ID then exec in the container and open the shell/terminal (sh):
docker exec -it 4112657f0f8f sh 

# In postgres docker image:
psql -U postgres # connect to postgres cli as default user postgres.

#  result of psql -U postgres:
postgres=# 

# Then you can do commands like /l , /dt or etc.
```
