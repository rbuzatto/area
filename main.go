package area

import "math"

// Pi é a relação comprimento e raio de uma circunferência
const Pi = 3.1416

// Circ calcula area da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect calcula area de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

func _TrianguloEq(base, altura float64) float64 {
	return base * altura / 2
}
