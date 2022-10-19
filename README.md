# PLANTILLA BÁSICA PARA UN PROYECTO DE GO
**Esta plantilla está basada en mi experiencia trabajando con proyectos de tipo REST con Go, así como  en aplicaciones en .Net, puede tomarlo como base para trabajar sus proyectos, si desea puede extender su funcionalidad para agregarlas a esta base de código**

📟 Esta es una plantilla básica para proyectos de Go, a continuación encontrará la descripción del scaffolding (estructura de archivos) del proyecto.

## Estructura
Esta es la estructura base de nuestra plantilla, con la cual pretendemos tener la menor cantidad de código acoplado.

``` bash
.
└── Api-Template/
    ├── cmd/
    │   ├── server
    │   └── console
    ├── src/
    │   ├── app/
    │   │   ├── controller/
    │   │   │   └── category_controller
    │   │   └── middleware
    │   ├── domain/
    │   │   ├── model
    │   │   └── service/
    │   │       └── category_service
    │   └── infrastructure/
    │       ├── database
    │       ├── repository/
    │       │   └── category_repo
    │       ├── router/
    │       │   ├── category_router
    │       │   └── main_router
    │       └── utils
    └── main
```
1. Iniciamos con el archivo principal de arranque
   **main.go**: este archivo debe mantenerse lo más limpio posible cuidando únicamente de construir nuestras dependencias y realizar el arranque de nuestra aplicación.

2. Nos encontramos con dos carpetas adicionales:

   **cmd**: esta carpeta contiene las diferentes implementaciones de las cuales dispondremos en nuestra aplicación, puede ser un microservicio, un servicio REST, una aplicación de consola etc., la intención es que esta sea la capa de comunicación.

   **src (source)**:  esta carpeta contiene básicamente todo nuestro código fuente, el código que es parte de las implementaciones que necesitamos.

3. Pasamos a la carpeta src y nos encontraremos con diferentes implementaciones por lo que encontraremos 3 capas bien identificadas

   3.1 **app**: contiene los puntos de arranque de nuestra aplicación, para este caso trabajamos con una aplicación REST  por lo que encontraremos controllers (handlers)  y middlewares, así como una implementación base de manejador de respuestas http.

   

   3.2 **domain**: contiene la funcionalidad propia de nuestra aplicación, esta deberá contener dos carpetas para, nuestras entidades de negocio, **model** en este caso usamos modelos para identificarlos aunque es mejor no exponer estas entidades al cliente. También dispondremos de una carpeta **services** donde alojaremos nuestros servicios o implementación de lógica de negocio. En nuestra capa de dominio debemos cuidar de hacer el menor uso posible de paquetes de terceros.

   3.3 **infrastructure**: contiene la implementación de paquetes de terceros y/o de base de datos utilizando ORM, sqlRaw etc. Para este caso usaremos el patrón repositorio y crearemos interfaces para crear abstracciones y poder inyectar funcionalidad de forma más eficiente.

4. Para trabajar en modo desarrollador puede utilizar la variable APP_ENVIRONMENT y definir comportamientos para sus espacios de trabajo.

5. Para trabajar con sus logs puede utilizar el paquete que ya se proporciona en el proyecto, es básicamente una implementación  sobre logrus con una inicialización más abstracta para que pueda extender la interfaz y adaptarlo a sus necesidades.

Necesitara agregar sus archivos .env para trabajar en desarrollo, considero que para iniciar es una buena opción, si desea trabara de forma más segura con sus secretos, considere utilizar una implementación con [Vault Hashicorp](https://developer.hashicorp.com/vault)

```diff
+ Esta implementación es solamente una base, todos los recursos son construidos al iniciar su programa por lo que para una funcionalidad multitenant quizá deba adaptar varias cosas de aquí.
```

## Recursos adicionales

- [A Tour of Go](https://tour.golang.org/welcome/1)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [50 Shades of Go: Traps, Gotchas, and Common Mistakes for New Golang Devs](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/)
- [Learn X in Y Minutes](https://learnxinyminutes.com/docs/go/)

## Herramientas

Hot reload: [fresh](https://github.com/pilu/fresh)

