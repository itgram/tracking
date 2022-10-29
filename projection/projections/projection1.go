package projections

import (
	"context"
	"fmt"
	"time"

	"github.com/itgram/green.system/system/grains"

	"github.com/itgram/tracking_projection/vehicle"
)

type Projection1State struct {
	VehicleId string
	Speed     int32
}

type Projection1Props struct{}

func (p *Projection1Props) GrainIdentity(event any) (string, bool) {

	switch e := event.(type) {

	case *vehicle.VehicleRegistered:
		return e.VehicleId, true

	case *vehicle.VehicleMaxSpeedAdjusted:
		return e.VehicleId, true
	}

	return "", false
}

func (p *Projection1Props) GrainTimeout() time.Duration { return 5 * time.Minute }
func (p *Projection1Props) Handler() grains.ProjectionHandler[*Projection1State] {
	return &projection1Handler{}
}
func (p *Projection1Props) HandlerTimeout() time.Duration { return 15 * time.Second }
func (p *Projection1Props) Init()                         {}
func (p *Projection1Props) LoadingTimeout() time.Duration { return 15 * time.Second }
func (p *Projection1Props) LoadState(ctx context.Context, id string) (*Projection1State, error) {

	// TODO:
	return &Projection1State{}, nil
}
func (p *Projection1Props) Log(projectionId, text string) { fmt.Println(projectionId, text) }
func (p *Projection1Props) Terminate()                    {}

type projection1Handler struct{}

func (p *projection1Handler) Handle(ctx *grains.ProjectionContext[*Projection1State], event any) error {

	fmt.Println("Projection 1 Handler: AggregateId", ctx.AggregateId())
	fmt.Println("Projection 1 Handler Revision:", ctx.Revision())
	fmt.Println("Projection 1 Handler StreamId:", ctx.StreamId())
	fmt.Println("Projection 1 Handler Event:", event)

	var newState = ctx.State()

	switch e := event.(type) {

	case *vehicle.VehicleRegistered:
		newState.VehicleId = e.VehicleId

	case *vehicle.VehicleMaxSpeedAdjusted:
		newState.Speed = e.MaxSpeed
	}

	fmt.Println("Projection 1 new State:", newState)

	return nil
}
