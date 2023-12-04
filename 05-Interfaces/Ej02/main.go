/*
Ejercicio 2 - Productos
Algunas tiendas ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total.
La empresa tiene 3 tipos de productos: Small, Medium y Large. (Se espera que sean muchos más)

Y los costos adicionales son:
Small: solo tiene el costo del producto
Medium: el precio del producto + un 3% de mantenerlo en la tienda
Large: el precio del producto + un 6% de mantenerlo en la tienda, y adicional a eso $2500 de costo de envío.

El porcentaje de mantenerlo en la tienda es en base al precio del producto.
El costo de mantener el producto en stock en la tienda es un porcentaje del precio del producto.
Se requiere una función factory que reciba el tipo de producto y el precio y retorne una interfaz Product que tenga el método Price.
Se debe poder ejecutar el método Price y que el método me devuelva el precio total en base al costo del producto y los adicionales
en caso que los tenga
*/

package main

import "fmt"

const (
	Small  = "Small"
	Medium = "Medium"
	Large  = "Large"
)

type Product interface {
	Price() float64
}

type EcommerceProduct struct {
	TypeOfProduct string
	ProductPrice  float64
}

func (ep EcommerceProduct) Price() float64 {
	switch ep.TypeOfProduct {
	case Medium:
		return ep.ProductPrice + ep.ProductPrice*.03
	case Large:
		return ep.ProductPrice + ep.ProductPrice*.06 + 2500
	default:
		return ep.ProductPrice
	}
}

func factory(typeOfProduct string, price float64) Product {
	return EcommerceProduct{
		TypeOfProduct: typeOfProduct,
		ProductPrice:  price,
	}
}

func main() {
	products := []Product{
		factory(Small, 1000),
		factory(Medium, 1000),
		factory(Large, 1000),
	}

	for _, product := range products {
		ep, ok := product.(EcommerceProduct)

		if !ok {
			panic("Error de type assertion")
		}

		fmt.Printf("\nTipo: %s\nPrecio: $ %.2f\nPrecio total: $ %.2f\n", ep.TypeOfProduct, ep.ProductPrice, ep.Price())
	}
}
