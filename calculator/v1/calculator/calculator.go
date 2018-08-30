package calculator


type Calculator interface {
	add(x, y float64) float64
	subtraction(x, y float64) float64
	multiplication(x, y float64) float64
	division(x, y float64) float64
}

type Operate interface {
	Name() string
	Begin() (string, error)
	Calculator
	ReportError(err error) string
	End() error
}