# Proyect-Go-CRUD-Gorm

**Luca Pierini**

---

## Descripción

**Proyect-Go-CRUD-Gorm** es una API RESTful desarrollada en Go que proporciona operaciones CRUD (Crear, Leer, Actualizar, Eliminar) para gestionar publicaciones. Utiliza Gorm como ORM para la interacción con la base de datos y PostgreSQL como sistema de gestión de bases de datos. El proyecto está diseñado para ejecutarse tanto en entornos locales como con Docker Compose.

## Funcionalidades Principales

- **Operaciones CRUD**: Gestión completa de publicaciones.
- **Interacción con Base de Datos**: Uso de Gorm para simplificar las operaciones con PostgreSQL.
- **Configuración Flexible**: Configuración de variables de entorno mediante archivos `.env`.
- **Fácil Despliegue**: Preparado para ejecutarse con Docker Compose.

## Endpoints

### Publicaciones

- **GET /posts**: Recupera todas las publicaciones.
- **GET /posts/:id**: Recupera una publicación específica por ID.
- **POST /post**: Crea una nueva publicación.
  ```json
  {
    "title": "Título de la Publicación",
    "body": "Contenido de la Publicación"
  }
  ```
- **PUT /posts/:id**: Actualiza una publicación por ID.
  ```json
  {
    "title": "Título Actualizado",
    "body": "Contenido Actualizado"
  }
  ```
- **DELETE /posts/:id**: Elimina una publicación por ID.

### Ping
- **GET /ping**: Verifica la conectividad del servicio.

## Tecnologías Utilizadas

- **Backend**: Go (Gin Framework)
- **Base de Datos**: PostgreSQL con Gorm
- **Contenedores**: Docker y Docker Compose

## Estructura del Proyecto
```
proyect-go-crud-gorm/
├── .env                 # Variables de entorno para la conexión a la base de datos
├── .gitignore           # Archivos ignorados por Git
├── api/                 # Código fuente de la API
│   ├── controllers/     # Controladores para manejar rutas
│   ├── initializers/    # Configuración inicial y conexión a la base de datos
│   ├── models/          # Modelos de base de datos
│   ├── main.go          # Punto de entrada de la aplicación
│   ├── go.mod           # Dependencias del módulo Go
│   ├── go.sum           # Checksum de dependencias
│   └── Dockerfile       # Dockerfile para la API
├── docker-compose.yml   # Configuración de Docker Compose
└── README.md            # Documentación (este archivo)
```

## Requisitos Técnicos

1. **Docker**: Asegúrate de tener instalado Docker y Docker Compose.
2. **Variables de Entorno**:
   - **DB_URL**: URL de conexión a PostgreSQL.
   - **PORT**: Puerto donde correrá la API.

## Instrucciones de Instalación

### Ejecución Local

1. Clona este repositorio:
   ```bash
   git clone https://github.com/lucapierini/proyect-go-crud-gorm.git
   cd proyect-go-crud-gorm
   ```

2. Instala las dependencias:
   ```bash
   cd api
   go mod download
   ```

3. Configura las variables de entorno creando un archivo `.env`.
Asegurate de comentar DB_URL="host=postgres y descomentar DB_URL="host=localhost.
   ```env
   DB_URL="host=localhost user=userTest password=gorm123456 dbname=gorm-crud port=5432 sslmode=disable"
   PORT=3000
   ```

4. Asegúrate de que la base de datos PostgreSQL esté corriendo localmente.
Puedes iniciar un contenedor de Docker con la imagen de Postgres con el siguiente comando:
   ```bash
   docker run -d --name postgres-db-crud -e POSTGRES_USER=userTest -e POSTGRES_PASSWORD=gorm123456 -e POSTGRES_DB=gorm-crud -p 5432:5432 postgres:15
   ```

5. Una vez que la base de datos esté activa ejecuta la aplicación:
   ```bash
   go run main.go
   ```

### Ejecución con Docker Compose

1. Clona este repositorio:
   ```bash
   git clone https://github.com/lucapierini/proyect-go-crud-gorm.git
   cd proyect-go-crud-gorm
   ```
2. Configura las variables de entorno creando un archivo `.env`. Asegurate de comentar DB_URL="host=localhost y descomentar DB_URL="host=postgres.
   ```env
   DB_URL="host=postgres user=userTest password=gorm123456 dbname=gorm-crud port=5432 sslmode=disable"
   PORT=3000
   ```
3. Levanta los servicios:
   ```bash
   docker-compose up --build
   ```

4. Accede a la API en `http://localhost:3000`.

## Notas

- El esquema de la base de datos se sincroniza automáticamente al iniciar la aplicación utilizando la funcionalidad AutoMigrate de Gorm.
- Credenciales por defecto para PostgreSQL en Docker Compose:
  - **Usuario**: `userTest`
  - **Contraseña**: `gorm123456`
  - **Base de Datos**: `gorm-crud`

## Autor
**Luca Pierini**

