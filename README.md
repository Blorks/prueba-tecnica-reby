# Prueba Técnica para Reby
En este readme se explican y detallan las decisiones tomadas en la realización de la prueba técnica para Reby

## Base de Datos
Por el hecho de siempre haber usado MySQL he elegido MySQL como gestor de la BBDD en lugar de PostgreSQL.

La aplicación se conecta a una instancia MySQL en Google Cloud. Sentios libres de crear una conexión a la BBDD para consultar su contenido:
Datos de conexión:
* Usuario: reby
* Contraseña: reby
* IP: 34.140.134.87
* Puerto: 3306

La siguiente imagen muestra el diseño de base de datos usado:
![Image text](https://www.dropbox.com/s/e7hwsb9mebtwoih/1.png?dl=0)

### Aclaraciones respecto al diseño de la BBDD
He intentando mantener el diseño lo más básico posible para no complicar demasiado la prueba:

* La tabla "Rides" tiene como columnas 2 fechas, que se usarán para calcular el coste del trayecto, y 2 FK's, una a la tabla "users" y otra a la tabla "vehicles"
* He añadido un atributo "balance" a la tabla "Users" para simular que el usuario tiene un saldo que pueda gastar. De esta forma, un usuario no podrá empezar un "ride" si no tiene suficiente saldo
* He añadido un atributo "state" a la tabla "Vehicles" porque entiendo que no se puede elegir un vehiculo para un trayecto que está siendo usado para otro trayecto. Esta columna sólo puede recibir los valores "FREE" e "IN_USE" para simular esa condición
* He creado una tabla para guardar las configuraciones de los "Rides", de forma que puedan ser modificadas sin tener que cambiar el microservicio. Por simplicidad, esta cuarta tabla no la he relacionado con "Rides", dando por sentado que todos los "Rides" tienen la misma configuración

## Servidor
El servidor ha sido desarrollado en Go. He elegido una estructura por capas para mantener cada "función" del servidor totalmente aislada. Las capas creadas son las siguientes:
* Capa models: Aquí se encuentran todas las entidades del microservicio (una por cada tabla de la BBDD)
* Capa services: Aquí se encuentra la lógica de negocio del microservicio
* Capa controllers: Aquí se encuentran los puntos de entrada del microservicio (endpoints)
* Capa repositories: Aquí se encuentra el ORM usado para controlar las consultas a la BBDD (He usado GORN)
