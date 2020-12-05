package main

import (
	"fmt"
	"net/rpc"
)

type Body struct {
	Alumno       string
	Materia      string
	Calificacion float64
}

func client() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var op int64
	for {
		fmt.Println("1) Agregar calificación de una materia")
		fmt.Println("2) Mostrar el promedio de un Alumno")
		fmt.Println("3) Mostrar el promedio general")
		fmt.Println("4) Mostrar el promedio de una materia ")
		fmt.Println("5) Ver info")
		fmt.Println("0) Exit")
		fmt.Scanln(&op)

		switch op {
		case 1:
			var nombre, materia string
			var calificacion float64
			fmt.Println("Nombre del alumno: ")
			fmt.Scanln(&nombre)
			fmt.Println("Materia: ")
			fmt.Scanln(&materia)
			fmt.Println("Calificación: ")
			fmt.Scanln(&calificacion)

			body := Body{nombre, materia, calificacion}

			var result string
			err = c.Call("Server.AgregarCalificacionPorMateria", body, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.AgregarCalificacionPorMateria =", result)
			}
		case 2:
			var alumno string
			fmt.Print("Alumno: ")
			fmt.Scanln(&alumno)

			var result float64
			err = c.Call("Server.ObtenerPromedioAlumno", alumno, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.ObtenerPromedioAlumno", result)
			}
		case 3:
			var result float64
			err = c.Call("Server.ObtenerPromedioAlumnos", "", &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.ObtenerPromedioAlumnos: ", result)
			}
		case 4:
			var materia string
			fmt.Println("Materia: ")
			fmt.Scanln(&materia)
			var result float64
			err = c.Call("Server.ObtenerPromedioPorMateria", materia, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.ObtenerPromedioPorMateria: ", result)
			}
		case 5:
			var result string
			err = c.Call("Server.VerInfo", "", &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.VerInfo", result)
			}
		case 0:
			return
		}
	}
}

func main() {
	client()
}
