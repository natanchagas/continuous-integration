package main

func TesteSoma(t *testing.T){
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado da soma inválido!\nResultado: %d\n Esperado: %d", total, 30)
	}
}