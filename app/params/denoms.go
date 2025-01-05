package params

import (
	"github.com/terra-money/core/v2/app/config"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterDenomsConfig() error {
	// sdk.RegisterDenom(config.Qubit, sdk.OneDec())
	// sdk.RegisterDenom(config.MilliQubit, sdk.NewDecWithPrec(1, 3))
	err := sdk.RegisterDenom(config.MicroQubit, sdk.NewDecWithPrec(1, 6))
	if err != nil {
		return err
	}
	// sdk.RegisterDenom(config.NanoQubit, sdk.NewDecWithPrec(1, 9))

	return nil
}
