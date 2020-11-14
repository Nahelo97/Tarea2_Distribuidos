Natalia        Herrera  Lora   201673501-1
Andrés Eduardo Shehadeh Gullón 201673560-7

-----------Como ejecutar-----------
Entrar a la carpeta repo2, donde se encuentra el makefile y escribir en consola

$make [nombre]

donde [nombre] corresponde al sistema que desea correr:
  -cliente
  -datanode
  -namenode
  --------------------------------------------
  -----------Cosas a considerar-----------
1.- Para mantener las IP's utilizadas en los códigos, se recomienda ejecutar los sistemas
   en las máquinas asignadas:
       datanode   --->maquina 93,94,95
       namenode   --->maquina 96
       cliente    --->cualquiera

2.-Si usted quiere utilizar su propio input de libros, remplace los archivos que se encuentran en la carpeta Books (en la maquina que usted vaya a ejecutar al cliente)

3.-Para evitar una sobrecarga de información y que se repitan los libros o sus chunks, hicimos que al inicio de la ejecución de los programas se borren sus archivos guardados temporales (no se borrara nada de la carpeta Books, pero si se borraran los libros descargados por el cliente por ejemplo)

4.-Si correrá los sistemas en otras máquinas haga los cambios necesarios en las conexiones, además de las instalaciones

5.-Todo funcionó correctamente, asique si hay algún problema no dude en contactarnos :D
