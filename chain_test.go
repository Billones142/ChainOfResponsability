package main_test

import (
	"testing"

	clases "github.com/Billones142/ChainOfResponsability/Clases"
	interfaces "github.com/Billones142/ChainOfResponsability/Interfaces"
)

func Test_LogicaAutorizador(t *testing.T) {
	prueba := clases.Autorizador{ManejadorBase: &clases.ManejadorBase{}}
	datos := &clases.PedidoDeLogin{}

	aniosMayorDeEdad := []int{2004, 2005, 2006}
	aniosMenorDeEdad := []int{2007, 2008, 2009}
	valoresQueDieronError := []int{}

	var datosI interfaces.Pedido = datos

	for _, num := range aniosMayorDeEdad {
		datos.AnioDeNacimiento = num
		if !prueba.Autorizar(&datosI) { // si no sale como se espera agregar a la lista
			valoresQueDieronError = append(valoresQueDieronError, num)
		}
	}

	for _, num := range aniosMenorDeEdad {
		datos.AnioDeNacimiento = num
		if prueba.Autorizar(&datosI) { // si no sale como se espera agregar a la lista
			valoresQueDieronError = append(valoresQueDieronError, num)
		}
	}

	if len(valoresQueDieronError) > 0 {
		t.Errorf("El test dio error con el/los resultados %v", valoresQueDieronError)
	}
}

func Test_LogicaAutenticador(t *testing.T) {
	prueba := clases.Autenticador{ManejadorBase: &clases.ManejadorBase{}}
	datos := &clases.PedidoDeLogin{}

	var datosI interfaces.Pedido = datos

	booleano, _ := prueba.ManejarPedido(&datosI)

	if booleano {
		t.Errorf("No se lleno ningun campo e igualmente dio correcto")
		t.FailNow()
	}

	datos.Nombre = "s"
	booleano, _ = prueba.ManejarPedido(&datosI)

	if booleano {
		t.Errorf("Con solo llenar el nombre ya fue valido")
		t.FailNow()
	}

	datos.Apellido = "m"
	booleano, _ = prueba.ManejarPedido(&datosI)

	if booleano {
		t.Error("agregando el apellido ya fue valido")
		t.FailNow()
	}

	datos.AnioDeNacimiento = 2024
	booleano, _ = prueba.ManejarPedido(&datosI)

	if !booleano {
		t.Errorf("agregando el anio 2024 no dio correcto")
		t.FailNow()
	}

	datos.AnioDeNacimiento = 2025
	booleano, _ = prueba.ManejarPedido(&datosI)

	if booleano {
		t.Errorf("agregando el anio 2025 dio correcto")
		t.FailNow()
	}
}
