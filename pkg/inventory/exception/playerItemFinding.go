package exception

import "fmt"

type PlayerItemsFinging struct {
	PlayerID string
}

func (e *PlayerItemsFinging) Error() string {
	return fmt.Sprintf("finding items for playerID: %s failed", e.PlayerID)
}
