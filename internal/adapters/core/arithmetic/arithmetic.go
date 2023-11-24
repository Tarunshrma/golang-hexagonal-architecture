package arithmetic

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (arthmetic *Adapter) Addition(op1 int32, op2 int32) (int32, error) {
	return op1 + op2, nil
}

func (arthmetic *Adapter) Substraction(op1 int32, op2 int32) (int32, error) {
	return op1 - op2, nil
}

func (arthmetic *Adapter) Multiplication(op1 int32, op2 int32) (int32, error) {
	return op1 * op2, nil
}

func (arthmetic *Adapter) Division(op1 int32, op2 int32) (int32, error) {
	return op1 / op2, nil
}
