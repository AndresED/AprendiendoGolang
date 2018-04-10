package main

//ES NECESARIO TENER INSTALADO LOS PAQUETES DE MYSQL Y DE GORM
// MYSQL => go get -u github.com/go-sql-driver/mysql
//GORM =>  go get -u github.com/jinzhu/gorm

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Producto estructura de la tabla producto
type Producto struct {
	gorm.Model   //ID,CreatedAt,DeleteAt
	CodigoBarras string
	Precio       uint
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@/edcurso?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Error en la conexión a la bd")
	}
	defer db.Close()
	fmt.Println("Se conecto a la BD")
	//CREANDO UNA TABLA
	/*
		db.CreateTable(&Producto{})
		fmt.Println("Se creo la tabla producto")
	*/
	//INSERTANDO UN VALOR A LA BD
	/*fmt.Println("INSERTANDO .....")
	db.Create(&Producto{CodigoBarras: "1215465656542", Precio: 2000})
	fmt.Println("SE INSERTO")*/
	//SELECT
	var p Producto

	db.First(&p, 2)
	fmt.Println("Precio del producto consultado", p.Precio)

	//UPDATE
	p.Precio = 4500
	db.Save(&p)
	fmt.Println("El nuevo Precio del producto consultado es", p.Precio)
}
