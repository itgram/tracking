package vehicle

import (
	"errors"

	"github.com/itgram/green/aggregate"
)

type Result = aggregate.Result[*State]

func RegisterVehicle(vehicleId string, model string) (Result, error) {

	var result = aggregate.NewResult(&State{})

	var err = result.Apply(&VehicleRegistered{
		VehicleId: vehicleId,
		Model:     model,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func AdjustMaxSpeed(state *State, maxSpeed int32) (Result, error) {

	var result = aggregate.NewResult(state)

	if maxSpeed >= 230 {
		return nil, errors.New("max speed must be greater than 0")
	}

	var err = result.Apply(&VehicleMaxSpeedAdjusted{
		VehicleId: state.VehicleId,
		MaxSpeed:  maxSpeed,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
