package vote_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/julianolorenzato/choosely/domain/vote"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	a := assert.New(t)

	t.Run("It should create a new vote", func(t *testing.T) {
		voterID := uuid.NewString()
		pollID := uuid.NewString()
		choosenOptions := []string{"Dog", "Avocado", "Purple"}

		v := vote.New(voterID, pollID, choosenOptions)

		a.NotNil(v)
		a.Equal(v.VoterID, voterID)
		a.Equal(v.PollID, pollID)
		a.Equal(v.ChoosenOptions, choosenOptions)
		a.NotEmpty(v.CreatedAt)
	})
}
