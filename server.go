package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
)

var materias = make(map[string]map[string]float64)
var alumnos = make(map[string]map[string]float64)

type Body struct {
	Alumno       string
	Materia      string
	Calificacion float64
}

type Server struct{}

func (this *Server) AgregarCalificacionPorMateria(body Body, reply *string) error {
	*reply = fmt.Sprintf("AgregarCalificacionPorMateria: %s, %s, %f", body.Alumno, body.Materia, body.Calificacion)

	alumno := alumnos[body.Alumno]
	if alumno == nil {
		alumno = make(map[string]float64)
	}
	if alumno[body.Materia] != 0 {
		return errors.New("La calificación ya fue asignada")
	}
	alumno[body.Materia] = body.Calificacion
	alumnos[body.Alumno] = alumno

	materia := materias[body.Materia]
	if materia == nil {
		materia = make(map[string]float64)
	}
	materia[body.Alumno] = body.Calificacion
	materias[body.Materia] = materia
	return nil
}

func (this *Server) ObtenerPromedioAlumno(alumno string, reply *float64) error {
	*reply = calcularPromedioDeAlumno(alumno)
	return nil
}
func (this *Server) ObtenerPromedioAlumnos(param string, reply *float64) error {
	var total float64
	for alumno, _ := range alumnos {
		total += calcularPromedioDeAlumno(alumno)
	}
	promedio := total / float64(len(alumnos))
	fmt.Println("total: ", total)
	fmt.Println("alumnos: ", float64(len(alumnos)))
	fmt.Println("promedio: ", promedio)
	fmt.Println("******************************")
	*reply = float64(promedio)
	return nil
}
func (this *Server) ObtenerPromedioPorMateria(materia string, reply *float64) error {
	var total float64
	for _, calificacion := range materias[materia] {
		total += calificacion
	}
	promedio := total / float64(len(materias[materia]))
	fmt.Println("Materia: ", materia)
	fmt.Println("total: ", total)
	fmt.Println("calificaciones: ", float64(len(materias[materia])))
	fmt.Println("promedio: ", promedio)
	fmt.Println("-----------------------------")
	*reply = promedio
	return nil
}

func (this *Server) VerInfo(param string, reply *string) error {
	fmt.Printf("Materias: %+v \n", materias)
	fmt.Printf("Alumnos: %+v \n", alumnos)
	fmt.Println("-----------------------------")
	*reply = "ok"
	return nil
}

func calcularPromedioDeAlumno(alumno string) float64 {
	var total float64 = 0.0
	for _, calificacion := range alumnos[alumno] {
		total = total + calificacion
	}

	promedio := total / float64(len(alumnos[alumno]))
	fmt.Println("Alumno: ", alumno)
	fmt.Println("total: ", total)
	fmt.Println("materias: ", float64(len(alumnos[alumno])))
	fmt.Println("promedio: ", promedio)
	fmt.Println("-----------------------------")
	return float64(promedio)
}

func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	fmt.Println("Servidor iniciado en puerto: 9999")
	if err != nil {
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {
	go server()

	var input string
	fmt.Scanln(&input)
}
