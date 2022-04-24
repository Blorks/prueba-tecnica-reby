# Prueba Técnica para Reby
En este readme se explican y detallan las decisiones tomadas en la realización de la prueba técnica para Reby

## Despliegue
La aplicación está desplegada en un servicio de Cloud Run. La URL para acceder es https://prueba-tecnica-reby-a5eb547yfa-ew.a.run.app

A continuación añado llamadas de ejemplo para ambos endpoint:

### Iniciar un "Ride"
URL: https://prueba-tecnica-reby-a5eb547yfa-ew.a.run.app/rides
JSON Body:
<code>
{
    "idUser": 1,
    "idVehicle": 1
}
</code>

### Finalizar un "Ride"
URL: https://prueba-tecnica-reby-a5eb547yfa-ew.a.run.app/rides/{id}/finish

Ejemplo de uso: https://prueba-tecnica-reby-a5eb547yfa-ew.a.run.app/rides/1/finish

---

En cuanto al despliegue en local, el archivo main.go está en la ruta: app/cmd/main.go

Simplemente hay que usar el comando <code>go run main.go</code> para lanzarlo

## Base de Datos
Por el hecho de siempre haber usado MySQL he elegido MySQL como gestor de la BBDD en lugar de PostgreSQL.

La aplicación se conecta a una instancia MySQL que corre en Google Cloud. Sentios libres de crear una conexión a la BBDD para consultar su contenido:
Datos de conexión:
* Usuario: reby
* Contraseña: reby
* IP: 34.140.134.87
* Puerto: 3306
* Nombre de la BD: prueba-tecnica-reby

La siguiente imagen muestra el diseño de base de datos usado:
![Image text](https://github.com/Blorks/prueba-tecnica-reby/blob/develop/cmd/server/bd_design_v3.png)

**_NOTE:_** Aunque la aplicación se lance en local, apunta a la base de datos creada en CloudSQL

### Aclaraciones respecto al diseño de la BBDD
He intentando mantener el diseño lo más básico posible para no complicar demasiado la prueba:

* La tabla "Rides" tiene como columnas 2 fechas, que se usarán para calcular el coste del trayecto, y 2 FK's, una a la tabla "users" y otra a la tabla "vehicles"
* He añadido un atributo "balance" a la tabla "Users" para simular que el usuario tiene un saldo que pueda gastar. De esta forma, un usuario no podrá empezar un "ride" si no tiene suficiente saldo
* He añadido un atributo "state" a la tabla "Vehicles" porque entiendo que no se puede elegir un vehiculo para un trayecto que está siendo usado para otro trayecto. Esta columna sólo puede recibir los valores "free" e "in_use" para simular esa condición

## Microservicio
El microservicio ha sido desarrollado en Go. He elegido una estructura por capas para mantener cada "funcionalidad" totalmente aislada. Las capas creadas son las siguientes:
* Capa models: Aquí se encuentran todas las entidades del microservicio (una por cada tabla de la BBDD)
* Capa services: Aquí se encuentra la lógica de negocio del microservicio
* Capa controllers: Aquí se encuentran los puntos de entrada del microservicio (endpoints)
* Capa repositories: Aquí se encuentra el ORM usado para controlar las consultas a la BBDD (He usado GORN)
* Capa dtos: Aquí se encuentran las entidades usadas para aplicar el patrón DTO para desvincular la información requerida por el cliente y las entidades de la base de datos

## Testing
Para el testeo de la aplicación, he querido dividir los tests en 2 partes: por un lado tests de integración y por otro lado tests unitarios.

Para los tests de integración, he creado una serie de archivos SQL que se ejecutan al principio de cada tests (según sean necesarios) para crear y poblar la base de datos que va a ser usada por el test. Además, he creado un archivo llamado sqlUtils.go en el que hay una función que ejecuta estos archivos SQL

Por otro lado, he creado un archivo serverConfig.go que se encarga de levantar un servidor en el que se puedan ejecutar los tests y ejecuta los archivos SQL que sean necesarios

Para poder ejecutar estos tests, he creado una base de datos en un contenedor de docker. Añado a continuación el contenido del archivo docker-compose para crear este contenedor con la base de datos y los datos de conexión de esta.

La conexión a esta base de datos es la siguiente:
* Usuario: root
* Contraseña: toor
* IP: localhost
* Puerto: 3333

**_NOTE:_** Para que los tests se ejecuten correctamente, hay que crear manualmente una base de datos llamada "prueba-tecnica-reby"

docker-compose:
<code>
version: '3.1'

services:

  db:
    image: mysql
    command: --default-authentication-plugin=mysql_native_password
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: toor

  adminer:
    image: adminer
    restart: always
    ports:
      - 3333:3306
</code>

En cuanto a los tests unitarios, no he conseguido hacerlos debido a mi falta de experiencia con el lenguaje (en la sección de Problemas y soluciones lo explico un poco más), aunque comento de todas formas el concepto de mi idea inicial:

Mi idea era usar la separación por capas que he montado para testear cada capa por separado, mockeando la capa inmediatamente inferior. De esta forma, por ejemplo, para testear la capa service, haría mocks de los repositorios que se usaran y comprobaría el número de veces que se llama a los repositorios y con que parámetros. Siguiendo con la misma idea, para testear la capa controller, haría mocks de la capa service.

De esta forma, creo una especie de caja negra con cada capa en la que testeo los parámetros de entrada y salida, y las llamadas a otras capas que se realizan

## Problemas y soluciones
Me he encontrado con un problema a la hora de gestionar las fechas. He visto que Go tiene una función IsZero del tipo time.Time que da un valor por defecto a una fecha (0000-00-00 00:00:00). He pensado en usar esta función para comprobar si un "Ride" había terminado o no, pero MySQL toma este valor como un dato inválido y no permite insertarlo.

Para solucionarlo he inicializado el valor de Created y Finished de la entidad "Ride" al mismo valor, y he tomado como que si esas 2 fechas son iguales, el "Ride" aun no ha terminado. Se que no es la mejor solución pero no quería complicar demasiado la prueba para algo que no es un requisito directo

---

Por simplicidad, he dado por sentado que los precios indicados en la prueba no van a cambiar nunca, así que los he añadido como constantes en el servicio de "Ride". Quizás una solución más escalable hubiera sido tener esos precios en la base de datos o incluso relacionados con la entidad "Ride" para que cada "Ride" pudiera tener un precio diferente (por ejemplo, para distintos tipos de usuario: free, premium...)

---

A la hora de hacer los tests, en concreto, los tests unitarios, he tenido problemas para hacer mocks, debido a que no termino de entender muy bien como se hacen. He usado GoMock para generar mocks de los repositorios de la aplicación, pero luego no he sabido usarlos en el test como tal. Entiendo que me ha faltado el uso de interfaces para poder usarlos, pero no se muy bien como enfrentarme a este problema

He dejado una rama subida llamada "unit-testing" en la que tengo instalado GoMock, hay un mock generado para el userRepository.go y empiezo a crear un test unitario para el rideService, por si queréis echarle un vistazo, pero no he conseguido avanzar más de ahí

PD: En los tests que he hecho, no testeo la función que calcula el coste de un Ride, precisamente porque considero que no tiene sentido testearla con un test de integración. Debería ser testeada con tests unitarios precisamente para poder usar mocks para simular el caso de error de la función y controlar de una forma sencilla el resultado de la función