package keeper_test

import (
	"github.com/alice/checkers/x/checkers/testutil"
	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreate3GamesHasSavedFifo(t *testing.T) {
	msgSrvr, keeper, context := setupMsgServerCreateGame(t)
	ctx := sdk.UnwrapSDKContext(context)
	//-1 -1
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: testutil.Alice,
		Black:   testutil.Bob,
		Red:     testutil.Carol,
	})
	//-1 2

	// Second game
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: testutil.Bob,
		Black:   testutil.Carol,
		Red:     testutil.Alice,
	})
	//1 -1
	systemInfo2, found := keeper.GetSystemInfo(ctx)
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId:        3,
		FifoHeadIndex: "1",
		FifoTailIndex: "2",
	}, systemInfo2)
	game1, found := keeper.GetStoredGame(ctx, "1")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Index:       "1",
		Board:       "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "b",
		Black:       testutil.Bob,
		Red:         testutil.Carol,
		MoveCount:   uint64(0),
		BeforeIndex: "-1",
		AfterIndex:  "2",
	}, game1)
	game2, found := keeper.GetStoredGame(ctx, "2")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Index:       "2",
		Board:       "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "b",
		Black:       testutil.Carol,
		Red:         testutil.Alice,
		MoveCount:   uint64(0),
		BeforeIndex: "1",
		AfterIndex:  "-1",
	}, game2)

	// Third game
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: testutil.Carol,
		Black:   testutil.Alice,
		Red:     testutil.Bob,
	})
	systemInfo3, found := keeper.GetSystemInfo(ctx)
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId:        4,
		FifoHeadIndex: "1",
		FifoTailIndex: "3",
	}, systemInfo3)
	game1, found = keeper.GetStoredGame(ctx, "1")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Index:       "1",
		Board:       "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "b",
		Black:       testutil.Bob,
		Red:         testutil.Carol,
		MoveCount:   uint64(0),
		BeforeIndex: "-1",
		AfterIndex:  "2",
	}, game1)
	game2, found = keeper.GetStoredGame(ctx, "2")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Index:       "2",
		Board:       "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "b",
		Black:       testutil.Carol,
		Red:         testutil.Alice,
		MoveCount:   uint64(0),
		BeforeIndex: "1",
		AfterIndex:  "3",
	}, game2)
	game3, found := keeper.GetStoredGame(ctx, "3")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Index:       "3",
		Board:       "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "b",
		Black:       testutil.Alice,
		Red:         testutil.Bob,
		MoveCount:   uint64(0),
		BeforeIndex: "2",
		AfterIndex:  "-1",
	}, game3)
}
