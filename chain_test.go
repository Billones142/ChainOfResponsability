package main_test

import (
	"testing"

	main "chainOfResponsability.com/example"
)

func LogicaAutorizador_test(t *testing.T) {
	prueba := main.Autorizador{&main.ManejadorBase{}}
	datos := &main.PedidoDeLogin{}

	valoresAProbar := []int{2004, 2005, 2006, 2007, 2008, 2009}
	valoresQueDanError := []int{}

	for _, num := range valoresAProbar {
		datos.AnioDeNacimiento = num
		if !prueba.Autorizar(datos) {
			valoresQueDanError = append(valoresQueDanError, num)
		}
	}

	if len(valoresQueDanError) > 0 {
		t.Errorf("El test dio error con el/los resultados %v", valoresQueDanError)
	}
}

func LogicaAutenticador_test(t *testing.T) {
	prueba := main.Autenticador{&main.ManejadorBase{}}
	datos := &main.PedidoDeLogin{}

	if prueba.ManejarPedido(datos) {
		t.Errorf("No hay logica de autenticacion actualmente, deberia dar verdadero todo el tiempo")
	}
}
