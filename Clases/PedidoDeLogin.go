package clases

import interfaces "github.com/Billones142/ChainOfResponsability/Interfaces"

// PedidoDeLogin representa una solicitud
type PedidoDeLogin struct {
	interfaces.Pedido
	Nombre           string
	Apellido         string
	AnioDeNacimiento int
	Sexo             bool //true:hombre false:mujer
}

func (r *PedidoDeLogin) GetNumeroDePedido() int {
	return 0
}
