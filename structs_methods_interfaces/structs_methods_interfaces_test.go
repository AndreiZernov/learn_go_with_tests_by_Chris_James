package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{shape: Rectangle{Width: 10, Height: 10}, hasArea: 40.0},
		{shape: Circle{Radius: 10}, hasArea: 62.83185307179586},
		{shape: Triangle{Base: 8, Height: 4}, hasArea: 20.94427190999916},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
