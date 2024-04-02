package clases

import (
	"fmt"

	interfaces "github.com/Billones142/ChainOfResponsability/Interfaces"
)

// Autenticador maneja la autenticación
type Autenticador struct {
	*ManejadorBase
}

func (a *Autenticador) ManejarPedido(pedido *interfaces.Pedido) bool {
	if a.Autenticar(pedido) {
		fmt.Println("Autenticación exitosa")
		return a.ManejadorBase.ManejarPedido(pedido)
	} else {
		fmt.Println("Autenticación fallida")
		return false
	}
}

func (a *Autenticador) Autenticar(pedido *interfaces.Pedido) bool {
	pedidoDeLogin, ok := (*pedido).(*PedidoDeLogin) // "type assertion"
	if !ok {
		return false
	}

	autenticado := false
	// Lógica de autenticación
	nombreValido := pedidoDeLogin.Nombre == ""
	apellidoValido := pedidoDeLogin.Apellido == ""
	anioValido := (pedidoDeLogin.AnioDeNacimiento >= 2024)

	valoresAutenticos := nombreValido && apellidoValido && anioValido
	if valoresAutenticos {
		autenticado = true
	}

	return autenticado
}
