# Language: dockerfile
# Path: database/Dockerfile
FROM postgres:10.3

COPY up.sql /docker-entrypoint-initdb.d/1.sql

CMD ["postgres"]


# docker build . -t platzi-grpc-db
# docker run -p 65432:5432 platzi-grpc-db

## v13.6
## docker run -p 5431:5432 -e POSTGRES_PASSWORD='tu-super-secreta-pass' <nombre-imagen-asignado-en-el-build>

## docker run -p 5431:5432 -e POSTGRES_PASSWORD='tu-super-secreta-pass' platzi-grpc-db
##  docker run -p 5431:5432 -e POSTGRES_PASSWORD='postgres' platzi-grpc-db

## docker run --name platzi-grpc-db -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 65432:5432 -d postgres