/*
Ejercicio 1
Crear un programa que cumpla los siguiente puntos:
Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.
Tener un slice global de Product llamado Products instanciado con valores.
2 métodos asociados a la estructura Product: Save(), GetAll(). El método Save() deberá tomar el slice de Products
y añadir el producto desde el cual se llama al método. El método GetAll() deberá imprimir todos los productos guardados en el slice Products.
Una función getById() al cual se le deberá pasar un INT como parámetro y retorna el producto correspondiente al parámetro pasado.
Ejecutar al menos una vez cada método y función definido desde main().
*/

package main

import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

func (p Product) Save() {
	Products = append(Products, p)
}

func (p Product) GetAll() {
	for _, product := range Products {
		fmt.Printf("ID: %v\nNombre: %v\nPrecio: %v\nDescripcion: %v\nCategoria: %v\n\n",
			product.ID, product.Name, product.Price, product.Description, product.Category)
	}
}

func getById(id int) Product {
	for _, product := range Products {
		if product.ID == id {
			return product
		}
	}
	return Product{}
}

var Products = []Product{
	{
		ID: 1, Name: "Coca Cola", Price: 1000, Description: "Bebida con gas", Category: "Bebidas",
	},
	{
		ID: 2, Name: "Galletitas Oreo", Price: 650, Description: "Galletitas de chocolate con crema", Category: "Galletitas",
	},
}

func main() {
	newProduct := Product{
		ID: 3, Name: "Coffler Block", Price: 1700, Description: "Chocolate con mani", Category: "Chocolates",
	}

	newProduct.Save()
	newProduct.GetAll()
	product := getById(1)
	if product != (Product{}) {
		fmt.Printf("Producto encontrado:\nID: %v\nNombre: %v\nPrecio: %v\nDescripcion: %v\nCategoria: %v\n\n",
			product.ID, product.Name, product.Price, product.Description, product.Category)
	}
}
