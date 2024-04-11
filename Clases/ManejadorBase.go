package clases

import (
	interfaces "github.com/Billones142/ChainOfResponsability/Interfaces"
)

// ManejadorBase es una implementación básica de la interfaz Handler
type ManejadorBase struct {
	successor interfaces.Handler
}

func (b *ManejadorBase) ManejarPedido(request *interfaces.Pedido, mensajes ...string) (bool, []string) {
	if b.successor != nil {
		return b.successor.ManejarPedido(request, mensajes...) // si hay un sucesor pasarle el pedido para que lo procece
	}
	return true, mensajes // si no hay otro sucesor entonces la cadena se completo con exito
}

func (b *ManejadorBase) SetSucesor(successor interfaces.Handler) {
	b.successor = successor
}
