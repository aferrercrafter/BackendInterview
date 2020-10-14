![POD](https://img.shields.io/badge/language-Go-blue.svg)
![POD](https://img.shields.io/badge/version-v1.0.0-yellow.svg)
![POD](https://img.shields.io/badge/license-MIT-lightgrey.svg)
![GoMod compatible](https://img.shields.io/badge/GoMod-compatible-4BC51D.svg?style=flat)
[![codecov](https://codecov.io/gh/aferrercrafter/operacion-fuego-quazar/branch/main/graph/badge.svg?token=FVJNCM55WP)](https://codecov.io/gh/aferrercrafter/operacion-fuego-quazar)

# Operacion Fuego de Quasar

Han Solo ha sido recientemente nombrado General de la Alianza
Rebelde y busca dar un gran golpe contra el Imperio Galáctico para
reavivar la llama de la resistencia.

El servicio de inteligencia rebelde ha detectado un llamado de auxilio de
una nave portacarga imperial a la deriva en un campo de asteroides. El
manifiesto de la nave es ultra clasificado, pero se rumorea que
transporta raciones y armamento para una legión entera.

Definición del problema aquí: [Problema](https://github.com/aferrercrafter/operacion-fuego-quazar/blob/main/docs/problema/OperacionFuegoQuasarV1.1.pdf)

## Mínimos requirimientos

Technology | Version
------- | --------
GoLang | 1.14

## Cómo ejecutar el programa

Para ejecutar , use `go run`:

    go run main.go

El servidor empezará a escuchar request en los siguientes endpoints:

**LOCAL MODE**

    [POST] localhost:8080/api/v1/topsecret
    [GET]  localhost:8080/api/v1/topsecret_split
    [POST] localhost:8080/api/v1/topsecret_split/{satellite_name}

**Google Cloud Platform**        

    [POST] http://fuego-quazar.rj.r.appspot.com/api/v1/topsecret
    [GET]  http://fuego-quazar.rj.r.appspot.com/api/v1/topsecret_split
    [POST] http://fuego-quazar.rj.r.appspot.com/api/v1/topsecret_split/{satellite_name}

**Swagger**

    http://localhost:8080/swagger/index.html

## Acerca de la solución

    // input: distancia al emisor tal cual se recibe en cada satélite
    // output: las coordenadas ‘x’ e ‘y’ del emisor del mensaje
    func GetLocation(distances ...float32) (x, y float32)

Para obtener la localicazión se uso el algoritmo de [Trileteración](https://es.wikipedia.org/wiki/Trilateraci%C3%B3n#:~:text=La%20trilateraci%C3%B3n%20es%20un%20m%C3%A9todo,forma%20an%C3%A1loga%20a%20la%20triangulaci%C3%B3n.) 2D

Se puede acceder a una implementación de dicho algoritmo en [Programming with Java: A Multimedia Approach](https://www.jblearning.com/science-technology/computing/computer-graphics/productdetails/9781449638610#:~:text=Programming%20with%20Java%3A%20A%20Multimedia%20Approach%20uses%20multimedia%2Dbased%20programs,%2C%20display%20video%2C%20and%20more.)

    // input: el mensaje tal cual es recibido en cada satélite
    // output: el mensaje tal cual lo genera el emisor del mensaje
    func GetMessage(messages ...[]string) (msg string)

Para obtener el mensaje se asumen los siguientes criterios dado las consideraciones

* Un defasaje de 0, es válido (todos los mensajes son del mismo tamaño)
* El defasaje se encuentra al principio del mensaje

La solución se basa en la búsqueda y eliminación del defasaje, y el encuentro con la primera palabra válida en la misma posición entre los tres satélites

## Code coverage

Para actualizar el report de coverage ejecutar desde la carpeta raiz del proyecto

    go test -coverprofile=coverage.txt -covermode=atomic ./...


Y para subir el reporte 

    bash <(curl -s https://codecov.io/bash) -t [iamatoken-nousar-...]




    



