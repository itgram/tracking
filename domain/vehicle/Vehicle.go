package vehicle

import (
	"errors"

	"github.com/itgram/minion/aggregate"
)

type Result = aggregate.Result[*State]

func RegisterVehicle(vehicleId string, model string) (Result, error) {

	result := aggregate.NewResult(&State{})

	return result.Apply(&VehicleRegistered{
		VehicleId: vehicleId,
		Model:     model,
	})
}

func AdjustMaxSpeed(state *State, maxSpeed int32) (Result, error) {

	result := aggregate.NewResult(state)

	if maxSpeed <= 0 {
		return result, errors.New("max speed must be greater than 0")
	}

	return result.Apply(&VehicleMaxSpeedAdjusted{
		VehicleId: state.VehicleId,
		MaxSpeed:  maxSpeed,
	})
}
