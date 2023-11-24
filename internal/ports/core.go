package ports

type ArithmeticPort interface {
	Addition(op1 int32, op2 int32) (int32, error)
	Substraction(op1 int32, op2 int32) (int32, error)
	Multiplication(op1 int32, op2 int32) (int32, error)
	Division(op1 int32, op2 int32) (int32, error)
}
