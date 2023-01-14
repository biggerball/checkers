package keeper_test

import (
	"github.com/alice/checkers/x/checkers/testutil"
	"github.com/alice/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateGame(t *testing.T) {
	msgServer, context := setupMsgServer(t)
	createResponse, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: testutil.Alice,
		Black:   testutil.Bob,
		Red:     testutil.Carol,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateGameResponse{
		GameIndex: "", // TODO: update with a proper value when updated
	}, *createResponse)
}