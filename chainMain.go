package main

import "fmt"

// PedidoDeLogin representa una solicitud
type PedidoDeLogin struct {
	Nombre           string
	Apellido         string
	AnioDeNacimiento int
	Sexo             bool //true:hombre false:mujer
}

// Handler define la interfaz para manejar una solicitud
type Handler interface {
	ManejarPedido(request *PedidoDeLogin) bool
	SetSucesor(successor Handler)
}

// ManejadorDelLogin maneja el registro
type ManejadorDelLogin struct {
	*ManejadorBase
}

func (l *ManejadorDelLogin) ManejarPedido(request *PedidoDeLogin) bool {
	fmt.Println("Registro de solicitud " + request.Apellido + " " + request.Nombre)
	return l.ManejadorBase.ManejarPedido(request)
}

// ManejadorBase es una implementación básica de la interfaz Handler
type ManejadorBase struct {
	successor Handler
}

func (b *ManejadorBase) ManejarPedido(request *PedidoDeLogin) bool {
	if b.successor != nil {
		return b.successor.ManejarPedido(request) // si hay un sucesor pasarle el pedido para que lo procece
	}
	return true // si no hay otro sucesor entonces la cadena se completo con exito
}

func (b *ManejadorBase) SetSucesor(successor Handler) {
	b.successor = successor
}

// Autenticador maneja la autenticación
type Autenticador struct {
	*ManejadorBase
}

func (a *Autenticador) ManejarPedido(request *PedidoDeLogin) bool {
	if a.Autenticar(request) {
		fmt.Println("Autenticación exitosa")
		return a.ManejadorBase.ManejarPedido(request)
	} else {
		fmt.Println("Autenticación fallida")
		return false
	}
}

func (a *Autenticador) Autenticar(request *PedidoDeLogin) bool {
	autenticado := false
	// Lógica de autenticación
	autenticado = true
	return autenticado
}

// Autorizador maneja la autorización
type Autorizador struct {
	*ManejadorBase
}

func (a *Autorizador) ManejarPedido(request *PedidoDeLogin) bool {
	if a.Autorizar(request) {
		fmt.Println("Autorización exitosa")
		return a.ManejadorBase.ManejarPedido(request)
	} else {
		fmt.Println("Autorización fallida")
		return false
	}
}

func (a *Autorizador) Autorizar(pedido *PedidoDeLogin) bool { //chequea que el usuario sea mayor de edad
	// Lógica de autorización
	autorizado := false
	if (2024 - pedido.AnioDeNacimiento) >= 18 {
		autorizado = true
	}

	return autorizado
}

func main() {
	// Configuracion para la cadena de responsabilidad

	// se crean los verificadores de la cadena
	autenticador := &Autenticador{&ManejadorBase{}}
	autorizador := &Autorizador{&ManejadorBase{}}
	manejadorDeLogin := &ManejadorDelLogin{&ManejadorBase{}}

	// se establece el orden en el que se ejecutan, en este caso dandole a cada uno su sucesor
	autenticador.SetSucesor(autorizador)
	autorizador.SetSucesor(manejadorDeLogin)

	// Simulación de Pedido
	Pedido := &PedidoDeLogin{
		Nombre:           "Stefano Nahuel",
		Apellido:         "Merino De Rui",
		AnioDeNacimiento: 2003,
		Sexo:             true,
	}
	if autenticador.ManejarPedido(Pedido) { // se le da al primero de la cadena el pedido
		fmt.Println("Login exitoso")
	} else {
		fmt.Println("Login fallido")
	}
}
