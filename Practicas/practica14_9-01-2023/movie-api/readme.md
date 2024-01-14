## API-MOVIES
primero instalar las paquetes necesarios 

github.com/joho/godotenv
go get github.com/gorilla/mux
go get github.com/lib/pq
go get github.com/rs/cors

# Administración de Películas con PostgreSQL y Go

## **Objetivo**:

 Desarrollar una aplicación en Go que se conecte a una base de datos PostgreSQL para administrar un conjunto de películas.

## **Descripción**:

1. **Creación de la Base de Datos**:
    - Diseñar e implementar una base de datos en PostgreSQL para almacenar información sobre películas. La estructura debe seguir el formato establecido en el ejercicio 2 de esta [práctica](https://www.notion.so/Pr-cticas-SGBD-II-fe7041dce12b410f9b156899ceceb3a7?pvs=21)
    - Scripts de SQL para la creación y configuración de la base de datos en una carpeta llamada SQL dentro de nuestro proyecto.
2. **Desarrollo del Backend en Go**:
    - Implementar un backend en Go que proporcione las siguientes funcionalidades (ABML):
        - **Alta (A) - Agregar Nueva Película**
            - **POST** `/api/peliculas`
            
            Este endpoint recibe la información de una nueva película  en el cuerpo de la solicitud y la agrega a la base de datos.
            
        - **Baja (B) - Eliminar Película**
            - **DELETE** `/api/peliculas/{id}`
            
            Este endpoint elimina una película específica de la base de datos, identificada por su ID. El ID de la película se pasa como parte de la URL.
            
        - **Modificación (M) - Actualizar Película**
            - **PUT** `/api/peliculas/{id}`
            
            Este endpoint actualiza la información de una película existente. El ID de la película a actualizar se especifica en la URL, y los nuevos datos se envían en el cuerpo de la solicitud.
            
        - **Listado (L) - Obtener Películas**
            - **GET** `/api/peliculas`
            
            Este endpoint recupera un listado de todas las películas en la base de datos. Ideal para consultas generales.
            
        - **Obtener Detalles de una Película Específica**
            - **GET** `/api/peliculas/{id}`
                
                Este endpoint proporciona los detalles completos de una película específica, identificada por su ID.
                
        - **Listado con Filtro por Nombre**
            - **GET** `/api/peliculas?nombre={nombre}`
            
            Este endpoint permite a los usuarios buscar películas por nombre utilizando búsqueda por contenido exacto o parcial (usando **`ILIKE`** o expresiones regulares con **`~`**). Si el parámetro `nombre` está presente, se filtran las películas; si no, se devuelven todas.
            
        - **Listado Avanzado con Filtros Adicionales**
            - **GET** (Opcional)
            
            Este endpoint podría ofrecer búsquedas más complejas, como filtrar por director, año, género, etc. Los filtros específicos se pasarían como parámetros de consulta.
            

## **Consideraciones**:

- Asegúrate de manejar adecuadamente las conexiones a la base de datos y gestionar los posibles errores.
- Si no se te ocurre como distribuir el código, te damos esta sugerencia: considera organizar el proyecto en los siguientes módulos:
    - main.go → Donde gerenciamos nuestros endpoints, conectamos el servidor y ejecutamos la conexión a la base de datos
    - controllers → donde vamos a guardar nuestros handlers
    - services → donde vamos a poner toda nuestra lógica de negocio, por ejemplo los request a nuestra base de datos para crear, editar eliminar o listar las películas así como la conexión a la base de datos en sí.
    - sql → carpeta donde guardamos nuestros scripts para la creación de las tablas
    - models → donde guardaremos todas nuestras structs que representan un único elemento de la tabla, en nuestro caso una película
- Para administrar endpoints y nuestra conexión a la red usaremos el [paquete mux](https://github.com/gorilla/mux) junto con el paquete net/http
    
    <aside>
    💡 Para entender mejor como funciona el paquete mux consulta esta página: 
    ‣
    
    </aside>
    
- Para manejar nuestra conexión con Postgres con el paquete database/sql de Go utilizaremos el paquete [github.com/lib/pq](http://github.com/lib/pq)
- Para manejar los permisos CORS, puedes utilizar el [paquete Cors](https://github.com/rs/cors)