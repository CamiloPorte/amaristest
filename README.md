# Saludos
Este repositorio tiene la respuesta a 3 enunciados

## Pregunta 1
Realizar una funcion que reciba una cadena de texto, que contenga una lista de nombres separados por comas ejemplo "Luis,Camilo,Andres,Laura" y devuelva dos parametros: una array de strings con los nombres ordenados alfabeticamente y un entero indicando la cantidad de nombres

## Pregunta 2
Realizar una funcion que reciba un numero entero "id" de un pokemon y devuelva su nombre. Para esto la funcion debe consultar una api de pokemon enviarle el id y obtener el campo "nombre" el cual va a retornar. URL = https://pokeapi.co/api/v2/pokemon-form/{{id}}

## Pregunta 3
Se dice que dos cadenas X y Y son amigas si existen dos cadenas no vacías u y v tales que X = uv e Y = vu. Por ejemplo, "tokyo" y "kyoto" son amigas, siendo u = “to” y v = “kyo”.
Escriba una funcion que reciba como entrada dos cadenas X e Y, e imprima si X e Y son amigas.

## Instrucciones de uso
Para correr las soluciones y probarlas se debe correr con los siguientes comandos, cada uno solicitará las entradas necesarias que se describen en los enunciados:

```sh
make runq1
make runq2
make runq3
```

Para correr los test de las piexas se debe correr el siguiente comando:

```sh
make test
```
