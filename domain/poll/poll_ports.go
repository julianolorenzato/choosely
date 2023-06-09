package poll

type PollRepository interface {
	GetByID(ID string) (*Poll, error)
	Create(poll *Poll) error
	Save(poll *Poll) error
}
