package tempconv

type TempFunction interface {
	Calculate(value float64) float64
	GetName() string
}

type CtoF struct {
	Name string
}
type CtoK struct {
	Name string
}
type FtoC struct{
	Name string
}
type FtoK struct {
	Name string
}
type KtoC struct {
	Name string
}
type KtoF struct {
	Name string
}

func (c CtoF) Calculate(value float64) float64 { return value*9.0/5.0 + 32.0}
func (c CtoK) Calculate(value float64) float64 { return value + 273.15}

func (f FtoK) Calculate(value float64) float64 { return (value + 459.67)* 5.0 / 9.0 }
func (f FtoC) Calculate(value float64) float64 { return (value - 32.0) * 5.0 / 9.0 }

func (k KtoF) Calculate(value float64) float64 { return (value * 9/5 ) - 459.67 }
func (k KtoC) Calculate(value float64) float64 { return value - 273.15}

func (k KtoC) GetName() string { return k.Name }
func (c CtoF) GetName() string { return c.Name }
func (f FtoC) GetName() string { return f.Name }
func (k KtoF) GetName() string { return k.Name }
func (c CtoK) GetName() string { return c.Name }
func (f FtoK) GetName() string { return f.Name }
