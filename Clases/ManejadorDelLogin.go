package clases

import (
	"fmt"

	interfaces "github.com/Billones142/ChainOfResponsability/Interfaces"
)

// ManejadorDelLogin maneja el registro
type ManejadorDelLogin struct {
	*ManejadorBase
}

func (l *ManejadorDelLogin) ManejarPedido(pedido *interfaces.Pedido, mensajes ...string) (bool, []string) {
	pedidoDeLogin, ok := (*pedido).(*PedidoDeLogin) // "type assertion"
	if !ok {
		return false, mensajes
	}

	fmt.Println("Registro de solicitud " + pedidoDeLogin.Apellido + " " + pedidoDeLogin.Nombre)
	return l.ManejadorBase.ManejarPedido(pedido, mensajes...)
}
