package clases

import (
	"fmt"

	interfaces "github.com/Billones142/ChainOfResponsability/Interfaces"
)

// Autorizador maneja la autorización
type Autorizador struct {
	*ManejadorBase
}

func (a *Autorizador) ManejarPedido(pedido *interfaces.Pedido) bool {
	if a.Autorizar(pedido) {
		fmt.Println("Autorización exitosa")
		return a.ManejadorBase.ManejarPedido(pedido)
	} else {
		fmt.Println("Autorización fallida")
		return false
	}
}

func (a *Autorizador) Autorizar(pedido *interfaces.Pedido) bool {
	// Se realiza un "type assertion" para afirmar que el valor subyacente de la interfaz es de tipo PedidoDeLogin
	pedidoDeLogin, ok := (*pedido).(*PedidoDeLogin)
	if !ok {
		// Si el type assertion falla, significa que el tipo subyacente de la interfaz no es PedidoDeLogin
		// Manejar el caso en el que el pedido no sea de tipo PedidoDeLogin (puede devolver un valor predeterminado o generar un error)
		// en este caso solo hago que devuelva falso para tomar como invalida la autenticacion
		return false
	}
	// Ahora se puede acceder a los campos de PedidoDeLogin

	// Lógica de autorización
	autorizado := false
	if (2024 - pedidoDeLogin.AnioDeNacimiento) >= 18 {
		autorizado = true
	}

	return autorizado
}
