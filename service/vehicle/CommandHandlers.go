package vehicle

import (
	"errors"
	"fmt"
	"time"

	"github.com/itgram/green.persistence.esdb/esdb"
	"github.com/itgram/green.persistence/persistence/aggregate"
	"github.com/itgram/green.persistence/persistence/aggregate/snapshot/strategy"
	"github.com/itgram/green.system/system/grains"

	"github.com/itgram/tracking_domain/vehicle"
)

type commandHandler struct{}

func (*commandHandler) Handle(ctx *grains.CommandContext[*vehicle.State], command any) error {

	switch cmd := command.(type) {

	case *RegisterVehicle:

		fmt.Println("RegisterVehicle")

		// TODO: validate the existing vehicle with the same id

		result, err := vehicle.RegisterVehicle(cmd.VehicleId, cmd.Model)
		if err != nil {
			return err
		}

		ctx.SetResult(result)

	case *AdjustMaxSpeedVehicle:

		fmt.Println("AdjustMaxSpeedVehicle")

		result, err := vehicle.AdjustMaxSpeed(ctx.State(), cmd.MaxSpeed)
		if err != nil {
			return err
		}

		ctx.SetResponse(&AdjustMaxSpeedVehicle{VehicleId: "123", MaxSpeed: 500})
		ctx.SetResult(result)

	default:
		return errors.New("unknown command received")
	}

	return nil
}

func NewAggregateProps(conn *esdb.Connection) grains.AggregateProps[*vehicle.State] {

	return &aggregateProps{
		handler:        &commandHandler{},
		handlerTimeout: 30 * time.Second,
		loadingTimeout: 30 * time.Second,
		repository: aggregate.NewRepository(
			conn.NewJournalStore(3),
			conn.NewSnapshotStore(2),
			&vehicle.State{},
		),
		strategy: strategy.NewSnapshotStrategy(
			strategy.WithMaxEventCount(4)),
	}
}

type aggregateProps struct {
	handler        grains.CommandHandler[*vehicle.State]
	handlerTimeout time.Duration
	loadingTimeout time.Duration
	repository     aggregate.Repository[*vehicle.State]
	strategy       strategy.Strategy
}

func (p *aggregateProps) Handler() grains.CommandHandler[*vehicle.State]   { return p.handler }
func (p *aggregateProps) HandlerTimeout() time.Duration                    { return p.handlerTimeout }
func (p *aggregateProps) LoadingTimeout() time.Duration                    { return p.loadingTimeout }
func (p *aggregateProps) Init()                                            {}
func (p *aggregateProps) GrainTimeout() time.Duration                      { return 5 * time.Minute }
func (p *aggregateProps) Log(aggregateId, text string)                     { fmt.Println(aggregateId, text) }
func (p *aggregateProps) PartitionSize() uint64                            { return 4 }
func (p *aggregateProps) Repository() aggregate.Repository[*vehicle.State] { return p.repository }
func (p *aggregateProps) Strategy() strategy.Strategy                      { return p.strategy }
func (p *aggregateProps) Terminate()                                       {}
