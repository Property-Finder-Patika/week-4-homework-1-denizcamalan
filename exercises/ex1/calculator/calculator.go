package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	mathFunction "github.com/Property-Finder-Patika/week-4-homework-1-denizcamalan/exercises/ex1/math"
)

var flag bool = true

type Calculator struct {
	functions []mathFunction.MathFunction
}

func (c *Calculator) addMathFunction(m mathFunction.MathFunction) {
	c.functions = append(c.functions, m)
}

// check lower case with error handling function
func (c *Calculator) doCalculation(name string, arg float64) (float64, error) { 
	var result float64
	for _, f := range c.functions {
		if strings.ToLower(name) == strings.ToLower(f.GetName()) {
			result = f.Calculate(arg)
			return result, nil
		}
	}
	return 0, errors.New("no such function exists:" + name)
}

func main() {
	startCalculator()
}

func startCalculator() {
	myCalculator := Calculator{}
	myCalculator.addMathFunction(mathFunction.Sin{Name: "Sine"})
	myCalculator.addMathFunction(mathFunction.Cos{Name: "Cosine"})
	myCalculator.addMathFunction(mathFunction.Log{Name: "Log"})

	fmt.Println("\nCalculator started.")
	fmt.Println("You can calculate using following functions")
	for _, f := range myCalculator.functions {
		fmt.Println(f.GetName())
	}

	for flag {
		var fName string
		var arg float64
		fmt.Println("> Enter name of the calculation or enter x to exit:")
		_, err := fmt.Scanf("%s", &fName)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		if fName == "x" || fName == "X"{
			flag = false
		} else {
			fmt.Println("> Enter a value for the calculation:")
			_, err := fmt.Scanf("%f", &arg)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			value, err := myCalculator.doCalculation(fName, arg)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Result of %s of %f : %f\n", fName, arg, value)
			}
		}
	}
	println("Bye!")
}