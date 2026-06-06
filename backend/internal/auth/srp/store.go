package srp

type ChallengeStore struct {
        Challenges map[string]Challenge
}

func NewChallengeStore() *ChallengeStore {
        return &ChallengeStore{
                Challenges: make(map[string]Challenge),
        }
}
