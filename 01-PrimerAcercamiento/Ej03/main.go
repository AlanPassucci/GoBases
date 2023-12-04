/*
Ejercicio 3 - Declaración de variables
Un profesor de programación está corrigiendo los exámenes de sus estudiantes de la materia Programación I para poder brindarles las correspondientes devoluciones. Uno de los puntos del examen consiste en declarar distintas variables.
Necesita ayuda para:
Detectar cuáles de estas variables que declaró el alumno son correctas.
Corregir las incorrectas.
*/

package main

func main() {
	// var 1nombre string -> Incorrecto: Error de compilacion, el nombre no puede empezar con numero
	// var apellido string -> Correcto
	// var int edad -> Incorrecto: El type debe ir luego del nombre de la variable
	// 1apellido := 6 -> Incorrecto: Error de compilacion, el nombre no puede empezar con numero
	// var licencia_de_conducir = true -> Correcto, pero no utiliza la notacion camelCase
	// var estatura de la persona int -> Incorrecto: Error de compilacion, las variables no pueden llevar espacios
	// cantidadDeHijos := 2 -> Correcto
}
