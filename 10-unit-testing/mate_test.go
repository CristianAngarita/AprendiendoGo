package main

import "testing"

func TestSumar(t *testing.T) {
	/* resultado := Sumar(3, 5)
	esperado := 8

	if resultado != esperado {
		t.Errorf("Error al sumar: se esperaba %d y se obtuvo %d", esperado, resultado)
	} */

	//Conjuntos de casos a probar
	casos := []struct {
		a, b, e int
	}{
		{3, 2, 5},
		{4, 5, 9},
		{456, 123, 579},
		{0, 0, 0},
		{-1, 1, 0},
	}

	for _, caso := range casos {
		r := Sumar(caso.a, caso.b)
		if r != caso.e {
			t.Errorf("Error al sumar: se esperaba %d y se obtuvo %d", caso.e, r)
		}
	}

}

func TestMayor(t *testing.T) {
	casos := []struct {
		a, b, e int
	}{
		{3, 2, 3},
		{3, 10, 10},
		{70, 21, 70},
		{30, 22, 30},
	}
	//recorre caso y se prueba
	for _, caso := range casos {
		r := Mayor(caso.a, caso.b)

		//Se compara el caso con el resultado esperado
		if r != caso.e {
			t.Errorf("Error al comparar: se esperaba %d y se obtuvo %d", caso.e, r)
		}
	}
}

func TestFibonaccio(t *testing.T) {
	casos := []struct {
		n, e int
	}{
		{0, 0},
		{1, 1},
		{7, 13},
		{40, 102334155},
		{50, 12586269025},
	}

	for _, caso := range casos {
		r := Fibonacci(caso.n)

		if r != caso.e {
			t.Errorf("Error al calcular fibonaccio: se esperaba %d y se obtuvo %d", caso.e, r)
		}
	}

}
