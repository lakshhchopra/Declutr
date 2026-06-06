package srp

import "time"

type Challenge struct {
        UserID string

        ServerSecret string
        ServerPublicKey string

        CreatedAt time.Time
}
