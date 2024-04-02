package interfaces

// Handler define la interfaz para manejar una solicitud
type Handler interface {
	ManejarPedido(request *Pedido) bool
	SetSucesor(successor Handler)
}
