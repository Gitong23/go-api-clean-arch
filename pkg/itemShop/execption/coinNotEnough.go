package execption

type CoinNotEnough struct{}

func (e *CoinNotEnough) Error() string {
	return "Coin is not enough"
}
