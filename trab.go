// Go ofrece soporte incluido para codificar y descifrar JSON,
// incluyendo desde tipos de datos incorporados hasta
// tipos de datos personalizados.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Alumno struct {
	Universidad                     string `json:"universidad"`
	Curso                           string `json:"curso"`
	Alumno                          string `json:"alumno"`
	Período                         string `json:"perido"`
	Lenguajedeprogramacionpreferido string `json:"lenguaje"`
	AspiraciónPostGraduación        string `json:"aspiracion"`
}
type Alumnos []Alumno

// Vamos a usar estas dos estructuras para demostrar cifrado y
// descifrado de los tipos personalizados mostrados a continuación.
func allAlumnos(w http.ResponseWriter, r *http.Request) {
	alumnos := Alumnos{
		Alumno{
			Universidad:                     "UTPL",
			Curso:                           "Procesos de Ingeniería de Software",
			Alumno:                          "Jordy Alexander Celi Mancheno",
			Período:                         "Abr/Ago 2021",
			Lenguajedeprogramacionpreferido: "JAVA",
			AspiraciónPostGraduación:        "Trabajar como desarrollador"},
	}
	fmt.Println("Endpoint Hit:Alumno Endpoint")
	json.NewEncoder(w).Encode(alumnos)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HomePage endpoint hit")
}
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/alumnos", allAlumnos)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
func main() {
	handleRequests()
}
