package projections

import (
	"context"
	"fmt"
	"time"

	"github.com/itgram/green.fabric/fabric/consumer"

	"github.com/itgram/tracking_projection/vehicle"
)

type Projection1State struct {
	VehicleId string
	Speed     int32
}

type Projection1Props struct{}

func (p *Projection1Props) Handle(ctx consumer.ProjectionHandlerContext[*Projection1State], event any) error {

	fmt.Println("Projection 1 Handler AggregateId:", ctx.AggregateId())
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
func (p *Projection1Props) HandleFailure(ctx consumer.ProjectionContext, error string) {
	fmt.Println(ctx.ProjectionId(), error)
}
func (p *Projection1Props) HandlerTimeout(ctx consumer.ProjectionContext) time.Duration {
	return 15 * time.Second
}
func (p *Projection1Props) Init(ctx consumer.ProjectionContext) {}
func (p *Projection1Props) LoadState(ctx context.Context, identity, kind string) (*Projection1State, error) {

	// TODO:
	return &Projection1State{}, nil
}
func (p *Projection1Props) LoadingTimeout(ctx consumer.ProjectionContext) time.Duration {
	return 15 * time.Second
}

func (p *Projection1Props) ProjectionIdentity(event any) (string, bool) {

	switch e := event.(type) {

	case *vehicle.VehicleRegistered:
		return e.VehicleId, true

	case *vehicle.VehicleMaxSpeedAdjusted:
		return e.VehicleId, true
	}

	return "", false
}

func (p *Projection1Props) ProjectionTTL(ctx consumer.ProjectionContext) time.Duration {
	return 5 * time.Minute
}

func (p *Projection1Props) Terminate(ctx consumer.ProjectionContext) {}
