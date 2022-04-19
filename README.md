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

## Base de Datos
Por el hecho de siempre haber usado MySQL he elegido MySQL como gestor de la BBDD en lugar de PostgreSQL.

La aplicación se conecta a una instancia MySQL en Google Cloud. Sentios libres de crear una conexión a la BBDD para consultar su contenido:
Datos de conexión:
* Usuario: reby
* Contraseña: reby
* IP: 34.140.134.87
* Puerto: 3306
* Nombre de la BD: prueba-tecnica-reby

La siguiente imagen muestra el diseño de base de datos usado:
![Image text](https://github.com/Blorks/prueba-tecnica-reby/blob/develop/cmd/server/bd_design_v3.png)

### Aclaraciones respecto al diseño de la BBDD
He intentando mantener el diseño lo más básico posible para no complicar demasiado la prueba:

* La tabla "Rides" tiene como columnas 2 fechas, que se usarán para calcular el coste del trayecto, y 2 FK's, una a la tabla "users" y otra a la tabla "vehicles"
* He añadido un atributo "balance" a la tabla "Users" para simular que el usuario tiene un saldo que pueda gastar. De esta forma, un usuario no podrá empezar un "ride" si no tiene suficiente saldo
* He añadido un atributo "state" a la tabla "Vehicles" porque entiendo que no se puede elegir un vehiculo para un trayecto que está siendo usado para otro trayecto. Esta columna sólo puede recibir los valores "FREE" e "IN_USE" para simular esa condición

## Microservicio
El microservicio ha sido desarrollado en Go. He elegido una estructura por capas para mantener cada "función" del totalmente aislada. Las capas creadas son las siguientes:
* Capa models: Aquí se encuentran todas las entidades del microservicio (una por cada tabla de la BBDD)
* Capa services: Aquí se encuentra la lógica de negocio del microservicio
* Capa controllers: Aquí se encuentran los puntos de entrada del microservicio (endpoints)
* Capa repositories: Aquí se encuentra el ORM usado para controlar las consultas a la BBDD (He usado GORN)
* Capa dtos: Aquí se encuentran las entidades usadas para aplicar el patrón DTO para desvincular la información requerida por el cliente y las entidades de la base de datos

## Testing

## Problemas y soluciones
Me he encontrado con un problema a la hora de gestionar las fechas. He visto que Go tiene una función IsZero del tipo time.Time que da un valor por defecto a una fecha (0000-00-00 00:00:00). He pensado en usar esta función para comprobar si un "Ride" había terminado o no, pero MySQL toma este valor como un dato inválido y no permite insertarlo.

Para solucionarlo he inicializado el valor de Created y Finished de la entidad "Ride" al mismo valor, y he tomado como que si esas 2 fechas son iguales, el "Ride" aun no ha terminado. Se que no es la mejor solución pero no quería complicar demasiado la prueba para algo que no es un requisito directo

---

Por simplicidad, he dado por sentado que los precios indicados en la prueba no van a cambiar nunca, así que los he añadido como constantes en el servicio de "Ride". Quizás una solución más escalable hubiera sido tener esos precios en la base de datos o incluso relacionados con la entidad "Ride" para que cada "Ride" pudiera tener un precio diferente (por ejemplo, para distintos tipos de usuario: free, premium...)
