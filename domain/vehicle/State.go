package vehicle

//go:generate protoc --go_out=. events.proto
//go:generate protoc --eventsource_out=. events.proto
// protoc --proto_path=. --go_out=. --go_opt=paths=source_relative vehicle/*.proto

func (state *State) When(event any) bool {

	switch e := event.(type) {

	case *VehicleRegistered:
		state.VehicleId = e.VehicleId
		state.Model = e.Model

	case *VehicleMaxSpeedAdjusted:
		state.MaxSpeed = e.MaxSpeed

	default:
		return false
	}

	return true
}
