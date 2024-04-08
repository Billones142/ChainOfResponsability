package main

import (
	"fmt"

	clases "github.com/Billones142/ChainOfResponsability/Clases"
	interfaces "github.com/Billones142/ChainOfResponsability/Interfaces"
)

func main() {
	// Configuracion para la cadena de responsabilidad

	// se crean los verificadores de la cadena
	autenticador := &clases.Autenticador{ManejadorBase: &clases.ManejadorBase{}}
	autorizador := &clases.Autorizador{ManejadorBase: &clases.ManejadorBase{}}
	manejadorDeLogin := &clases.ManejadorDelLogin{ManejadorBase: &clases.ManejadorBase{}}

	// se establece el orden en el que se ejecutan, en este caso dandole a cada uno su sucesor
	autenticador.SetSucesor(autorizador)
	autorizador.SetSucesor(manejadorDeLogin)

	// Simulación de Pedido
	pedido := clases.PedidoDeLogin{
		Nombre:           "Stefano Nahuel",
		Apellido:         "Merino De Rui",
		AnioDeNacimiento: 2003,
		Sexo:             true,
	}

	var pedidoInterfaz interfaces.Pedido = &pedido

	if autenticador.ManejarPedido(&pedidoInterfaz) { // se le da al primero de la cadena el pedido
		fmt.Println("Login exitoso")
	} else {
		fmt.Println("Login fallido")
	}
}
