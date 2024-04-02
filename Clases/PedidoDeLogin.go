package clases

import interfaces "chainOfResponsability.com/example/Interfaces"

// PedidoDeLogin representa una solicitud
type PedidoDeLogin struct {
	interfaces.Pedido
	Nombre           string
	Apellido         string
	AnioDeNacimiento int
	Sexo             bool //true:hombre false:mujer
}
