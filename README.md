# Anime & Manga Tracker — Backend

Este repositorio contiene el backend del proyecto de Series Tracker, enfocado en registrar animes vistos y mangas leídos.
La idea principal es tener una API REST que maneje toda la lógica y los datos, mientras que el frontend solo se encarga de mostrarlos.

El backend está hecho en Go usando `net/http` y una base de datos SQLite.

---

# Cómo está pensado

A diferencia del laboratorio anterior, aquí el servidor **no genera HTML**.
En lugar de eso, responde únicamente con JSON.

Esto permite que cualquier cliente (web, app móvil, etc.) pueda consumir la API sin depender de cómo se ve la interfaz.

---

# Tecnologías

* Go (net/http)
* SQLite
* Swagger (OpenAPI)

---

# Modelo de datos

Cada serie (anime o manga) tiene esta estructura:

```json
{
  "id": 1,
  "title": "Frieren",
  "type": "anime",
  "total": 24,
  "progress": 12,
  "image": "https://..."
}
```

* `type`: "anime" o "manga"
* `total`: número total de episodios o capítulos
* `progress`: cuánto llevas visto o leído

---

# Endpoints

## GET /series

Devuelve todas las series

También se puede buscar por título:

```
/series?q=frieren
```

---

## GET /series/{id}

Devuelve una serie específica

---

## POST /series

Crea una serie nueva

Ejemplo:

```json
{
  "title": "One Piece",
  "type": "anime",
  "total": 1000,
  "progress": 10,
  "image": "https://..."
}
```

---

## PUT /series/{id}

Actualiza toda la serie

---

## PUT /series/progress

Actualiza solo el progreso (episodios o capítulos)

```json
{
  "id": 1,
  "progress": 25
}
```

---

## DELETE /series/{id}

Elimina una serie

---

# Swagger

Se usa Swagger para documentar la API.

Sirve básicamente para:

* ver todos los endpoints
* probarlos directamente desde el navegador
* tener claro qué datos se envían y se reciben

Se puede acceder en:

```
/swagger/
```

---

# CORS

El frontend y el backend corren en distintos puertos, entonces el navegador bloquea las peticiones por seguridad.

Para evitar eso, se configuró CORS permitiendo todos los orígenes durante desarrollo con estos headers:

```
Access-Control-Allow-Origin: *
Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS
Access-Control-Allow-Headers: Content-Type
```

---

# Cómo correrlo

1. Instalar dependencias

```
go mod tidy
```

2. Ejecutar el servidor

```
go run .
```

Después queda corriendo en:

```
http://localhost:8080
```

---

# Deploy

El backend está desplegado en Render:

```

```

Swagger también está disponible ahí:

```
```

# Estructura

```
backend/
│── main.go
│── handlers.go
│── models.go
│── db.go
│── openapi.yaml
│── go.mod
```

---

# Frontend

El frontend está en otro repositorio (requerido por el lab):

https://github.com/Tiffany24630/Proyecto-1-Full-Stack.-Web-front 
---

# Comentario final

La parte importante de este proyecto no es solo que funcione, sino cómo está separado:

* el backend solo maneja datos
* el frontend solo los consume

Eso hace que la aplicación sea más flexible y fácil de escalar a futuro.

---
