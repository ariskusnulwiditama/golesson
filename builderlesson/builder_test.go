package builderlesson

import (
	"fmt"
	"testing"
)

// Product
type Car struct {
	brand  string
	color  string
	engine string
}

// builder interface
type CarBuilder interface {
	SetBrand(brand string) CarBuilder
	SetColor(color string) CarBuilder
	SetEngine(color string) CarBuilder
	Build() Car
}

// concrete builder
type SportCarBuilder struct {
	car Car
}

func (b *SportCarBuilder) SetBrand(brand string) CarBuilder {
	b.car.brand = brand
	return b
}

func (b *SportCarBuilder) SetColor(color string) CarBuilder {
	b.car.color = color
	return b
}

func (b *SportCarBuilder) SetEngine(engine string) CarBuilder {
	b.car.engine = engine
	return b
}

func (b *SportCarBuilder) Build() Car {
	return b.car
}

// director
type CarDirector struct {
	builder CarBuilder
}

func (d *CarDirector) SetBuilder(builder CarBuilder) {
	d.builder = builder
}

func (d *CarDirector) BuildSportsCar() Car {
	return d.builder.SetBrand("Ferrari").SetColor("red").SetEngine("V8").Build()
}

func TestBuilder(t *testing.T) {
	director := CarDirector{}
	builder := &SportCarBuilder{}
	director.SetBuilder(builder)

	sportsCar := director.BuildSportsCar()
	fmt.Println("Sports Car:")
	fmt.Println("Brand:", sportsCar.brand)
	fmt.Println("Color:", sportsCar.color)
	fmt.Println("Engine:", sportsCar.engine)
}
