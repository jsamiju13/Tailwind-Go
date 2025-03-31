package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

func conexionBD() (conexion *sql.DB) {
	Driver := "mysql"
	Usuario := "root"
	Contrasenia := ""
	Nombre := "sistema"

	conexion, err := sql.Open(Driver, Usuario+":"+Contrasenia+"@tcp(127.0.0.1)/"+Nombre)

	if err != nil {
		panic(err.Error())
	}
	return conexion
}

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

// quito la funcion  y procedo y la coloco en una funcion aparte//

func main() {
	/*fmt.Println("Probando")
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "probando hola")
	})*/
	http.HandleFunc("/", inicio)
	http.HandleFunc("/crear", crear)
	http.HandleFunc("/insertar", insertar)

	http.HandleFunc("/editar", editar)
	http.HandleFunc("/actualizar", Actualizar)
	http.HandleFunc("/borrar", borrar)

	//creando servidor //*
	fmt.Println("run Server: http://localhost:3000")
	http.ListenAndServe("localhost:3000", nil)
}

type producto struct {
	Id     int
	Nombre string
	Precio float32
	Stock  int
}

func inicio(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(rw, "probando hola")
	conexionGoty := conexionBD()
	Registros, err := conexionGoty.Query("SELECT * FROM Productos")

	if err != nil {
		panic(err.Error())
	}
	productos := producto{}
	arregloProducto := []producto{}

	for Registros.Next() {

		var Id int
		var Nombre string
		var Precio float32
		var Stock int

		err = Registros.Scan(&Id, &Nombre, &Precio, &Stock)

		if err != nil {
			panic(err.Error())
		}

		productos.Id = Id
		productos.Nombre = Nombre
		productos.Precio = Precio
		productos.Stock = Stock

		arregloProducto = append(arregloProducto, productos)
	}

	plantillas.ExecuteTemplate(w, "inicio", arregloProducto)
}

func crear(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "crear", nil)
}

func insertar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		Nombre := r.FormValue("nombre")
		Precio := r.FormValue("precio")
		Stock := r.FormValue("stock")

		conexionGoty := conexionBD()
		insertarRegistros, err := conexionGoty.Prepare("INSERT INTO productos(Nombre,Precio,Stock) VALUES(?,?,?)")

		if err != nil {
			panic(err.Error())
		}
		insertarRegistros.Exec(Nombre, Precio, Stock)

		http.Redirect(w, r, "/", 301)
	}
}

func borrar(w http.ResponseWriter, r *http.Request) {
	idProducto := r.URL.Query().Get("id")

	conexionGoty := conexionBD()
	borrarRegistros, err := conexionGoty.Prepare("DELETE FROM Productos WHERE Id=?")

	if err != nil {
		panic(err.Error())
	}
	borrarRegistros.Exec(idProducto)

	http.Redirect(w, r, "/", 301)
}

func editar(w http.ResponseWriter, r *http.Request) {
	idProducto := r.URL.Query().Get("id")
	fmt.Println(idProducto)

	conexionGoty := conexionBD()
	editarRegistro, err := conexionGoty.Query("SELECT * FROM Productos WHERE Id=?", idProducto)

	if err != nil {
		panic(err.Error())
	}

	productos := producto{}

	for editarRegistro.Next() {

		var Id int
		var Nombre string
		var Precio float32
		var Stock int

		err = editarRegistro.Scan(&Id, &Nombre, &Precio, &Stock)

		if err != nil {
			panic(err.Error())
		}

		productos.Id = Id
		productos.Nombre = Nombre
		productos.Precio = Precio
		productos.Stock = Stock

	}
	fmt.Println(productos)
	plantillas.ExecuteTemplate(w, "editar", productos)

}

func Actualizar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		Id := r.FormValue("id")
		Nombre := r.FormValue("nombre")
		Precio := r.FormValue("precio")
		Stock := r.FormValue("stock")

		conexionGoty := conexionBD()
		modificarRegistros, err := conexionGoty.Prepare("UPDATE Productos SET Nombre=?, Precio =?, Stock=? WHERE Id=?")

		if err != nil {
			panic(err.Error())
		}
		modificarRegistros.Exec(Nombre, Precio, Stock, Id)

		http.Redirect(w, r, "/", 301)
	}
}
