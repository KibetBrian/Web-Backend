package db

import (
	"context"
	"testing"

	"github.com/KibetBrian/backend/utils"
	"github.com/stretchr/testify/require"
)

func TestRegisterCandidate(t *testing.T) {
	user := RegisterUser(t)
	positions := []string{"President", "Govonor"}
	len := len(positions)
	newCandidate := RegisterContestantParams{
		FullName: user.FullName,
		Email: user.Email,
		Position: positions[utils.GenerateRandInt(0,len-1)],
		Description: utils.RandomSentenceGenerator(),
	}
	ctx := context.Background()
	res, err := testQueries.RegisterContestant(ctx,newCandidate)
	require.NoError(t, err)
	require.Equal(t, newCandidate.FullName, res.FullName)
	require.Equal(t, newCandidate.Position, res.Position)
}