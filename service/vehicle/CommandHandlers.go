package vehicle

import (
	"context"
	"errors"
	"fmt"

	"github.com/itgram/minion.system/system"

	"github.com/itgram/tracking_domain/vehicle"
)

func CommandHandler(ctx context.Context, state *vehicle.State, command any) (system.CommandResponse[*vehicle.State], error) {

	switch cmd := command.(type) {

	case *RegisterVehicle:

		fmt.Println("RegisterVehicle")

		// TODO: validate the existing vehicle with the same id

		result, err := vehicle.RegisterVehicle(cmd.VehicleId, cmd.Model)
		if err != nil {
			return nil, err
		}

		return system.NewCommandResponse(nil, result, nil), nil

	case *AdjustMaxSpeedVehicle:

		fmt.Println("AdjustMaxSpeedVehicle")

		result, err := vehicle.AdjustMaxSpeed(state, cmd.MaxSpeed)
		if err != nil {
			return nil, err
		}

		return system.NewCommandResponse(nil, result, nil), nil
	}

	return nil, errors.New("unknown command received")
}
