package interfaces

// Handler define la interfaz para manejar una solicitud
type Handler interface {
	ManejarPedido(request *Pedido, mensajes ...string) (bool, []string)
	SetSucesor(successor Handler)
}
