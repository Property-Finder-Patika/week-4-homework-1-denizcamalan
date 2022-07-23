package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	tempFunction "github.com/Property-Finder-Patika/week-4-homework-1-denizcamalan/exercises/ex2/tempconv"
)


func main() {
	
	functions()
}

type Converter struct {
	functions []tempFunction.TempFunction
}

func (c *Converter) addMathFunction(m tempFunction.TempFunction) {
	c.functions = append(c.functions, m)
}

var flag bool = true

func functions() {

	myConverter := Converter{}

	myConverter.addMathFunction(tempFunction.CtoF{Name: "CtoF"})
	myConverter.addMathFunction(tempFunction.CtoK{Name: "CtoK"})
	myConverter.addMathFunction(tempFunction.FtoC{Name: "FtoC"})
	myConverter.addMathFunction(tempFunction.FtoK{Name: "FtoK"})
	myConverter.addMathFunction(tempFunction.KtoC{Name: "KtoC"})
	myConverter.addMathFunction(tempFunction.KtoF{Name: "KtoF"})

	fmt.Println("\nConverter started.")
	fmt.Println("You can convert using following functions")
	for _, f := range myConverter.functions {
		fmt.Println(f.GetName())
	}

	for flag {

		var fName string
		var arg float64
		fmt.Println("> Enter name of the convertion or enter q to exit:")
		_, err := fmt.Scanf("%s", &fName)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		if fName == "q" || fName == "Q"{
			flag = false
		} else {
			fmt.Println("> Enter a value for the convertion:")
			_, err := fmt.Scanf("%f", &arg)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			value, err := myConverter.makeConverter(fName, arg)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Conversion of `%s` of %f is %f\n", fName, arg, value)
			}
		}
	}
	println("Bye!")

}

func (c Converter) makeConverter(name string, arg float64) (float64, error) {
	var result float64
	for _, f := range c.functions {
		if strings.ToLower(name) == strings.ToLower(f.GetName()) {
			result = f.Calculate(arg)
			return result, nil
		}
	}
	return 0, errors.New("no such function exists:" + name)
}
