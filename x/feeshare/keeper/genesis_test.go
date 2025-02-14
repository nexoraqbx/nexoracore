package keeper_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/terra-money/core/v2/app/test_helpers"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/terra-money/core/v2/x/feeshare/types"
)

type GenesisTestSuite struct {
	test_helpers.AppTestSuite
}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(GenesisTestSuite))
}

func (suite *GenesisTestSuite) TestFeeShareInitGenesis() {
	testCases := []struct {
		name     string
		genesis  types.GenesisState
		expPanic bool
	}{
		{
			"default genesis",
			types.GenesisState{
				Params: types.DefaultParams(),
			},
			false,
		},
		{
			"custom genesis - feeshare disabled",
			types.GenesisState{
				Params: types.Params{
					EnableFeeShare:  false,
					DeveloperShares: types.DefaultDeveloperShares,
					AllowedDenoms:   []string{"uqubit"},
				},
			},
			false,
		},
		{
			"custom genesis - feeshare enabled, 0% developer shares",
			types.GenesisState{
				Params: types.Params{
					EnableFeeShare:  true,
					DeveloperShares: sdk.NewDecWithPrec(0, 2),
					AllowedDenoms:   []string{"uqubit"},
				},
			},
			false,
		},
		{
			"custom genesis - feeshare enabled, 100% developer shares",
			types.GenesisState{
				Params: types.Params{
					EnableFeeShare:  true,
					DeveloperShares: sdk.NewDecWithPrec(100, 2),
					AllowedDenoms:   []string{"uqubit"},
				},
			},
			false,
		},
		{
			"custom genesis - feeshare enabled, all denoms allowed",
			types.GenesisState{
				Params: types.Params{
					EnableFeeShare:  true,
					DeveloperShares: sdk.NewDecWithPrec(10, 2),
					AllowedDenoms:   []string(nil),
				},
			},
			false,
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.AppTestSuite.Setup() // reset

			if tc.expPanic {
				suite.Require().Panics(func() {
					suite.App.Keepers.FeeShareKeeper.InitGenesis(suite.Ctx, tc.genesis)
				})
			} else {
				suite.Require().NotPanics(func() {
					suite.App.Keepers.FeeShareKeeper.InitGenesis(suite.Ctx, tc.genesis)
				})

				params := suite.App.Keepers.FeeShareKeeper.GetParams(suite.Ctx)
				suite.Require().Equal(tc.genesis.Params, params)
			}
		})
	}
}
