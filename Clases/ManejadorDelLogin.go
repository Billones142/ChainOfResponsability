package clases

import (
	"fmt"

	interfaces "github.com/Billones142/ChainOfResponsability/Interfaces"
)

// ManejadorDelLogin maneja el registro
type ManejadorDelLogin struct {
	*ManejadorBase
}

func (l *ManejadorDelLogin) ManejarPedido(pedido *interfaces.Pedido) bool {
	pedidoDeLogin, ok := (*pedido).(*PedidoDeLogin) // "type assertion"
	if !ok {
		return false
	}

	fmt.Println("Registro de solicitud " + pedidoDeLogin.Apellido + " " + pedidoDeLogin.Nombre)
	return l.ManejadorBase.ManejarPedido(pedido)
}
