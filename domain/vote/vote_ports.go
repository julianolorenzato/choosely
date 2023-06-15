package vote

type VoteRepository interface {
	GetPollResults(pollID string) map[string]uint
	GetByID(ID string) *Vote
	Create(*Vote) error
}
