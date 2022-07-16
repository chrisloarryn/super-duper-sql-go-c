# Language: dockerfile
# Path: database/Dockerfile
FROM mysql:8.0.29-oracle
##postgres:10.3

COPY up.sql /docker-entrypoint-initdb.d/1.sql

CMD ["mysql"]


# docker build . -t edteam-course-db
# docker run -p 65432:5432 edteam-course-db

## v13.6
## docker run -p 5431:5432 -e POSTGRES_PASSWORD='tu-super-secreta-pass' <nombre-imagen-asignado-en-el-build>

## docker run -p 5431:5432 -e POSTGRES_PASSWORD='tu-super-secreta-pass' edteam-course-db
##  docker run -p 5431:5432 -e POSTGRES_PASSWORD='postgres' edteam-course-db

## docker run --name edteam-course-db -e MYSQL_ROOT_USER=mysql -e MYSQL_ROOT_PASSWORD=mysql -p 33060:3306 -d mysql/mysql-server:8.0.29-oracle