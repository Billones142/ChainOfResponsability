package clases

import (
	interfaces "github.com/Billones142/ChainOfResponsability/Interfaces"
)

// Autenticador maneja la autenticación
type Autenticador struct {
	*ManejadorBase
}

func (a *Autenticador) ManejarPedido(pedido *interfaces.Pedido, mensajes ...string) (bool, []string) {
	if a.Autenticar(pedido) {
		return a.ManejadorBase.ManejarPedido(pedido, append(mensajes, "Autenticación exitosa")...)
	} else {
		return false, append(mensajes, "Autenticación fallida")
	}
}

func (a *Autenticador) Autenticar(pedido *interfaces.Pedido) bool {
	autenticado := false

	pedidoDeLogin, ok := (*pedido).(*PedidoDeLogin) // "type assertion"
	if ok {
		// Lógica de autenticación
		nombreValido := pedidoDeLogin.Nombre != ""
		apellidoValido := pedidoDeLogin.Apellido != ""
		anioValido := (pedidoDeLogin.AnioDeNacimiento <= 2024) && (pedidoDeLogin.AnioDeNacimiento != 0)

		valoresAutenticos := nombreValido && apellidoValido && anioValido
		if valoresAutenticos {
			autenticado = true
		}

	}

	return autenticado
}
