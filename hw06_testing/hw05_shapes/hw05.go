package hw05

import (
	"errors"
	"fmt"
	"log"
	"math"
)

type Shape interface {
	Area() float64
}

type circle struct {
	radius int
}

func (c *circle) Area() float64 {
	r := float64(c.radius)
	return math.Pi * math.Pow(r, 2.0)
}

type triangle struct {
	base, height int
}

func (t *triangle) Area() float64 {
	return 0.5 * float64(t.base) * float64(t.height)
}

type rectangle struct {
	width, height int
}

func (r *rectangle) Area() float64 {
	return float64(r.width) * float64(r.height)
}

func CalculateArea(s any) (float64, error) {
	sh, ok := s.(Shape)
	if !ok {
		textError := errors.New("this type of shape isn't available")
		return 0.0, textError
	}

	return sh.Area(), nil
}

func PrintCalculateAreaSomeShapes() {
	var (
		area float64
		err  error
	)

	myCircle := &circle{5}
	area, err = CalculateArea(myCircle)
	if err != nil {
		log.Println("error:", err)
	}
	fmt.Println("Circle: radius", myCircle.radius)
	fmt.Println("The area of the circle:", area)

	myRectangle := &rectangle{10, 5}
	area, err = CalculateArea(myRectangle)
	if err != nil {
		log.Println("error:", err)
	}
	fmt.Println("Rectangle: width", myRectangle.width, ", height", myRectangle.height)
	fmt.Println("The area of the rectangle:", area)

	myTriangle := &triangle{8, 6}
	area, err = CalculateArea(myTriangle)
	if err != nil {
		log.Println("error:", err)
	}
	fmt.Println("Triangle: base", myTriangle.base, ", height", myTriangle.height)
	fmt.Println("The area of the triangle:", area)

	notShape := "not a shape"
	_, err = CalculateArea(notShape)
	if err != nil {
		log.Println("error:", err)
	}
}
