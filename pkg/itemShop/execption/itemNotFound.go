package execption

type ItemNotFound struct{}

func (e *ItemNotFound) Error() string {
	return "Item not found"
}