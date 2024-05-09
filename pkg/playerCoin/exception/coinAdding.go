package exception

type CoinAdding struct{}

func (e *CoinAdding) Error() string {
	return "Failed to add coin"
}
