package model

import "time"

type (
	PlayerCoin struct {
		ID       int       `json:"id"`
		PlayerID string    `json:"player_id"`
		Amount   int       `json:"amount"`
		CreateAt time.Time `json:"createAt"`
	}

	CoinAddingReq struct {
		PlayerID string
		Amount   int64 `json:"amount" validate:"required,gt=0"`
	}

	PlayerCoinShowing struct {
		PlayerID string `json:"playerID"`
		Coin     int64  `json:"coin"`
	}
)
