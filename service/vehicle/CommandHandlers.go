package vehicle

import (
	"errors"
	"fmt"
	"time"

	"github.com/WhatsLab/grain.fabric/fabric/command"
	"github.com/WhatsLab/grain.persistence.esdb/esdb"
	"github.com/WhatsLab/grain.persistence/persistence"
	"github.com/WhatsLab/grain.persistence/persistence/aggregate/snapshot/strategy"
	"github.com/WhatsLab/grain/aggregate"

	"github.com/itgram/tracking/domain/vehicle"
)

func NewAggregateProps(kind string, conn *esdb.Connection) command.AggregateProps[*vehicle.State] {

	return &aggregateProps{
		handlerTimeout: 30 * time.Second,
		loadingTimeout: 30 * time.Second,
		repository: persistence.NewRepository(
			conn.NewJournalStore(kind, 3),
			conn.NewSnapshotStore(kind, 2),
			&vehicle.State{},
		),
		strategy: strategy.NewSnapshotStrategy(
			strategy.WithMaxEventCount(4)),
	}
}

type aggregateProps struct {
	handlerTimeout time.Duration
	loadingTimeout time.Duration
	repository     persistence.Repository[*vehicle.State]
	strategy       strategy.Strategy
}

func (p *aggregateProps) AggregateTTL(ctx command.AggregateContext) time.Duration {
	return 1 * time.Minute
}
func (*aggregateProps) Handle(ctx command.CommandContext[*vehicle.State], command any) (aggregate.Result[*vehicle.State], error) {

	switch cmd := command.(type) {

	case *RegisterVehicle:

		fmt.Println("RegisterVehicle")

		// validate the existing vehicle with the same id
		if ctx.AggregateExists() {
			return nil, errors.New("vehicle already exists")
		}

		var result, err = vehicle.RegisterVehicle(cmd.VehicleId, cmd.Model)
		if err != nil {
			return nil, err
		}

		return result, nil

	case *AdjustMaxSpeedVehicle:

		fmt.Println("AdjustMaxSpeedVehicle")

		var result, err = vehicle.AdjustMaxSpeed(ctx.State(), cmd.MaxSpeed)
		if err != nil {
			return nil, err
		}

		ctx.SetResponse(&AdjustMaxSpeedVehicle{VehicleId: "123", MaxSpeed: 500})

		return result, nil

	default:
		return nil, errors.New("unknown command received")
	}
}
func (p *aggregateProps) HandleFailure(ctx command.AggregateContext, error string) {
	fmt.Println(ctx.AggregateId(), error)
}
func (p *aggregateProps) HandlerTimeout(ctx command.AggregateContext) time.Duration {
	return p.handlerTimeout
}
func (p *aggregateProps) Init(ctx command.AggregateContext) {
	fmt.Println("------- Init aggregate")
}
func (p *aggregateProps) LoadingTimeout(ctx command.AggregateContext) time.Duration {
	return p.loadingTimeout
}
func (p *aggregateProps) PartitionSize(ctx command.AggregateContext) uint64 { return 4 }

func (p *aggregateProps) Repository(ctx command.AggregateContext) persistence.Repository[*vehicle.State] {
	return p.repository
}
func (p *aggregateProps) Strategy(ctx command.AggregateContext) strategy.Strategy { return p.strategy }
func (p *aggregateProps) Terminate(ctx command.AggregateContext) {
	fmt.Println("------- Terminate aggregate")
}
