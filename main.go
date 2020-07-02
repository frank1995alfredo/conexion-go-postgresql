package main

import (
	_ "database/sql"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Usuario ...
type Usuario struct {
	ID     int `gorm:"primary_key"`
	Nombre string
	Email  string
}

//Persona ...
type Persona struct {
	ID       int    `gorm:"primary_key"`
	Nombre   string `gorm:"size:50"`
	Apellido string `gorm:"size:50"`
	Genero   string `gorm:"size:50"`
	Cedula   string `gorm:"size:50"`
}

//Migracion del modelo usuario
func modeloUsuario() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=golang password=1234 sslmode=disable")
	defer db.Close()

	if err != nil {
		log.Fatal(err.Error())
	}
	db.Debug().DropTableIfExists(&Usuario{})
	db.Debug().AutoMigrate(&Usuario{})
}

//Migracion del modelo persona
func modeloPersona() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=golang password=1234 sslmode=disable")
	defer db.Close()

	if err != nil {
		log.Fatal(err.Error())
	}

	//db.Debug().DropTableIfExists(&Persona{})
	db.Debug().AutoMigrate(&Persona{})

}

//funcion para insertar un usuario
func insertarUsuario() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=golang password=1234 sslmode=disable")
	defer db.Close()

	if err != nil {
		log.Fatal(err.Error())
	}

	usuario := &Usuario{Nombre: "Franklin", Email: "frank1995alfredo@gmail.com"}
	db.Create(usuario)
	fmt.Println("Usuario insertado")
}

//funcion personas de forma masiva
func insertarPersona() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=golang password=1234 sslmode=disable")
	defer db.Close()

	if err != nil {
		log.Fatal(err.Error())
	}

	personas := []Persona{
		{Nombre: "Franklin2", Apellido: "Canadas2", Genero: "Masculino"},
		{Nombre: "Maria2", Apellido: "Olmedo2", Genero: "Femenino"},
		{Nombre: "Justin2", Apellido: "Kent2", Genero: "Masculino"},
	}

	for _, persona := range personas {
		db.Create(&persona)
	}
	fmt.Println("Datos insertados correctamente")
}

func insertarPersona2(nombre string, apellido string, genero string) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=golang password=1234 sslmode=disable")
	defer db.Close()

	if err != nil {
		log.Fatal(err.Error())
	}

	if nombre == "" || apellido == "" || genero == "" {
		fmt.Println("Falta ingresar un dato.")
	} else {
		persona := &Persona{Nombre: nombre, Apellido: apellido, Genero: genero}
		db.Create(persona)
		fmt.Println("Persona registrada satisfactoriamente.")
	}
}

//funcion para buscar una persona
func presentarPersona() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=golang password=1234 sslmode=disable")
	defer db.Close()

	if err != nil {
		log.Fatal(err.Error())
	}
	persona := []Persona{}
	db.Debug().Find(&persona).Rows()
	fmt.Println(&persona)
}

//funcion para actualizar a una persona
func actualizarPersona(id int, nombre string, apellido string, genero string) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=golang password=1234 sslmode=disable")
	defer db.Close()

	if err != nil {
		log.Fatal(err.Error())
	}

	db.Table("personas").Where("id = ?", id).Update("nombre", nombre, "apellido", apellido, "genero", genero).Order("id asc")
	fmt.Println("Persona actualizada")
}

//funcion para eliminar a una persona
func eliminarPersona(id int) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=golang password=1234 sslmode=disable")
	defer db.Close()

	if err != nil {
		log.Fatal(err.Error())
	}

	db.Table("personas").Where("id = ?", id).Delete(&Persona{})
	fmt.Println("Persona eliminada")
}

func conexionBD() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=golang password=1234 sslmode=disable")
	defer db.Close()

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Conexion exitosa")

}

func main() {
	conexionBD()
	//modeloUsuario()
	//modeloPersona()
	//insertarUsuario()
	//insertarPersona()
	presentarPersona()
	//eliminarPersona()
	//actualizarPersona(3, "Juanita2", "Belen2", "Femenino2")
	//insertarPersona2("", "Olmedo", "Masculino")
}
