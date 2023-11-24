package api

import "github.com/Tarunshrma/golang-hexagonal-architecture/internal/ports"

type Adapter struct {
	arith ports.ArithmeticPort
	db    ports.DBPort
}

func NewAdapter(a ports.ArithmeticPort, db ports.DBPort) *Adapter {
	return &Adapter{arith: a, db: db}
}

func (apia *Adapter) GetAddition(a, b int32) (int32, error) {
	answer, err := apia.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}

	err = apia.db.AddToHistory(answer, "Addition")
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (apia *Adapter) GetSubtraction(a, b int32) (int32, error) {
	answer, err := apia.arith.Substraction(a, b)
	if err != nil {
		return 0, err
	}

	err = apia.db.AddToHistory(answer, "Subtraction")
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (apia *Adapter) GetMultiplication(a, b int32) (int32, error) {
	answer, err := apia.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}

	err = apia.db.AddToHistory(answer, "Multiplication")
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (apia *Adapter) GetDivision(a, b int32) (int32, error) {
	answer, err := apia.arith.Division(a, b)
	if err != nil {
		return 0, err
	}

	err = apia.db.AddToHistory(answer, "Division")
	if err != nil {
		return 0, err
	}

	return answer, nil
}
