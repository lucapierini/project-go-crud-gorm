# project-go-crud-gorm

Comando para ejecutar un contenedor de docker con imagen de Postgres
docker run -d --name postgres-db-crud -e POSTGRES_USER=userTest -e POSTGRES_PASSWORD=gorm123456 -e POSTGRES_DB=gorm-crud -p 5432:5432 postgres:15


Comando para ejecutar la api local
*desde la carpeta api 
go run main.go