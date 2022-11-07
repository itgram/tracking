package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/itgram/green.encoding.protobuf/json"
	"github.com/itgram/green.encoding.protobuf/protobuf"
	"github.com/itgram/green.encoding/encoding"
	"github.com/itgram/green.fabric/fabric"
	"github.com/itgram/green.fabric/fabric/consumer"
	"github.com/itgram/green.persistence.esdb/esdb"
	"github.com/itgram/green.persistence/persistence/flow/offset"
	"github.com/itgram/green.persistence/persistence/flow/stream"

	"github.com/itgram/tracking/projection/projections"
	_ "github.com/itgram/tracking/projection/vehicle"
)

var _projections []*consumer.Projection

func main() {

	var _, cancel = context.WithCancel(
		context.Background())

	conn := esdb.NewConnection(
		"esdb://127.0.0.1:2113?tls=false",
		esdb.WithAuthentication("admin", "changeit"),
		esdb.WithSerialization("json", func(encoding string) encoding.Serializer {

			if encoding == "json" {
				return json.NewSerializer()
			}

			return protobuf.NewSerializer()
		}),
	)

	var err = conn.Connect()
	if err != nil {
		fmt.Printf("Error to connect to tracking database: %v\n", err)
		return
	}

	defer conn.Close()

	var server = fabric.NewServer(
		fabric.NewNodeConfigurtion("localhost", "my_cluster", 0, fabric.ConsulProvider))

	var kinds []fabric.ActorKind

	var projection1 = consumer.NewProjectionActorKind("projection1", func(kind string, identityOnly bool) consumer.ProjectionProps[*projections.Projection1State] {

		return &projections.Projection1Props{}
	})

	kinds = append(kinds, projection1.Kind())

	_projections = append(_projections, projection1)

	kinds = append(kinds,
		consumer.NewSubscriptionActorKind("stream", func(kind string) consumer.SubscriptionProps {

			return NewSubscriptionProps(4, conn, server)
		}))

	err = server.Start(kinds...)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	defer server.Shutdown(true)

	// start a single instance of the master subscriber
	_, err = server.GetActor("v1", "stream")
	if err != nil {
		fmt.Printf("spawn actor error: %v\n", err)
		return
	}

	// ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	// defer stop()

	// Stop when a signal is sent.
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)

	<-c
	cancel()
}

func NewSubscriptionProps(bufferSize uint32, conn *esdb.Connection, server *fabric.Server) consumer.SubscriptionProps {

	return &SubscriptionProps{
		bufferSize: bufferSize,
		conn:       conn,
		server:     server,
	}
}

type SubscriptionProps struct {
	bufferSize  uint32
	conn        *esdb.Connection
	offsetStore offset.Store
	server      *fabric.Server
	stream      stream.Store
}

func (*SubscriptionProps) HandleFailure(ctx consumer.SubscriptionContext, error string) {
	fmt.Println(error)
}
func (*SubscriptionProps) HandlerTimeout(ctx consumer.SubscriptionContext) time.Duration {
	return 30 * time.Second
}
func (p *SubscriptionProps) Init(ctx consumer.SubscriptionContext) {

	var stream = p.conn.NewStreamStore(ctx.SubscriptionId(), p.bufferSize)

	p.stream = stream
	p.offsetStore = p.conn.NewOffsetStore(stream)
}
func (p *SubscriptionProps) OffsetStore(ctx consumer.SubscriptionContext) offset.Store {
	return p.offsetStore
}
func (p *SubscriptionProps) Projections(ctx consumer.SubscriptionContext) []*consumer.Projection {
	return _projections
}
func (p *SubscriptionProps) Request(ctx consumer.SubscriptionContext, identity, kind string, message any, timeout time.Duration) (any, error) {

	return p.server.Request(identity, kind, message, timeout)
}

func (p *SubscriptionProps) Stream(ctx consumer.SubscriptionContext) stream.Store           { return p.stream }
func (p *SubscriptionProps) SubscriptionTTL(ctx consumer.SubscriptionContext) time.Duration { return 0 }
func (p *SubscriptionProps) Terminate(ctx consumer.SubscriptionContext) {

	fmt.Println(">>> subscription terminated")
	// TODO: restart again a new virtual actor
}
