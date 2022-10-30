package vehicle

import (
	"errors"
	"fmt"
	"time"

	"github.com/itgram/green.persistence.esdb/esdb"
	"github.com/itgram/green.persistence/persistence"
	"github.com/itgram/green.persistence/persistence/aggregate/snapshot/strategy"
	"github.com/itgram/green.system/system/actors"
	"github.com/itgram/green/aggregate"

	"github.com/itgram/tracking_domain/vehicle"
)

type commandHandler struct{}

func (*commandHandler) Handle(ctx *actors.CommandContext[*vehicle.State], command any) (aggregate.Result[*vehicle.State], error) {

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

func NewAggregateProps(kind string, conn *esdb.Connection) actors.AggregateProps[*vehicle.State] {

	return &aggregateProps{
		handler:        &commandHandler{},
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
	handler        actors.CommandHandler[*vehicle.State]
	handlerTimeout time.Duration
	loadingTimeout time.Duration
	repository     persistence.Repository[*vehicle.State]
	strategy       strategy.Strategy
}

func (p *aggregateProps) Handler() actors.CommandHandler[*vehicle.State] { return p.handler }
func (p *aggregateProps) HandlerTimeout() time.Duration                  { return p.handlerTimeout }
func (p *aggregateProps) LoadingTimeout() time.Duration                  { return p.loadingTimeout }
func (p *aggregateProps) Init() {
	fmt.Println("------- Init aggregate")
}
func (p *aggregateProps) ActorTimeout() time.Duration                        { return 1 * time.Minute }
func (p *aggregateProps) Log(aggregateId, text string)                       { fmt.Println(aggregateId, text) }
func (p *aggregateProps) PartitionSize() uint64                              { return 4 }
func (p *aggregateProps) Repository() persistence.Repository[*vehicle.State] { return p.repository }
func (p *aggregateProps) Strategy() strategy.Strategy                        { return p.strategy }
func (p *aggregateProps) Terminate() {
	fmt.Println("------- Terminate aggregate")
}
