package main_test

import (
	"testing"

	main "chainOfResponsability.com/example"
)

func Test_LogicaAutorizador(t *testing.T) {
	prueba := main.Autorizador{&main.ManejadorBase{}}
	datos := &main.PedidoDeLogin{}

	aniosMayorDeEdad := []int{2004, 2005, 2006}
	aniosMenorDeEdad := []int{2007, 2008, 2009}
	valoresQueDieronError := []int{}

	for _, num := range aniosMayorDeEdad {
		datos.AnioDeNacimiento = num
		if !prueba.Autorizar(datos) { // si no sale como se espera agregar a la lista
			valoresQueDieronError = append(valoresQueDieronError, num)
		}
	}

	for _, num := range aniosMenorDeEdad {
		datos.AnioDeNacimiento = num
		if prueba.Autorizar(datos) { // si no sale como se espera agregar a la lista
			valoresQueDieronError = append(valoresQueDieronError, num)
		}
	}

	if len(valoresQueDieronError) > 0 {
		t.Errorf("El test dio error con el/los resultados %v", valoresQueDieronError)
	}
}

func Test_LogicaAutenticador(t *testing.T) {
	prueba := main.Autenticador{&main.ManejadorBase{}}
	datos := &main.PedidoDeLogin{}

	if !prueba.ManejarPedido(datos) {
		t.Errorf("No hay logica de autenticacion actualmente, deberia dar verdadero todo el tiempo")
	}
}
