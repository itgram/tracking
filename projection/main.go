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
	"github.com/itgram/green.persistence.esdb/esdb"
	"github.com/itgram/green.persistence/persistence/flow/offset"
	"github.com/itgram/green.persistence/persistence/flow/stream"
	"github.com/itgram/green.system/system"
	"github.com/itgram/green.system/system/grains"
)

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

	var server = system.NewServer(
		system.NewNodeConfigurtion("localhost", "my_cluster", 0))

	grains.RegisterStreamSubscription(server, "subscription_master", func() grains.SubscriptionProps {

		return NewSubscriptionProps(4, conn, "subscription_master")
	})

	err = server.Start()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	defer server.Shutdown(true)

	// start a single instance of the master subscriber
	err = server.GetGrain("grp1", "subscription_master")
	if err != nil {
		fmt.Printf("spawn grain error: %v\n", err)
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

func NewSubscriptionProps(bufferSize uint32, conn *esdb.Connection, groupId string) grains.SubscriptionProps {

	var stream = conn.NewStreamStore(groupId, bufferSize)

	return &SubscriptionProps{
		bufferSize:  bufferSize,
		conn:        conn,
		groupId:     groupId,
		offsetStore: conn.NewOffsetStore(stream),
		stream:      stream,
	}
}

type SubscriptionProps struct {
	bufferSize  uint32
	conn        *esdb.Connection
	groupId     string
	offsetStore offset.Store
	stream      stream.Store
}

func (*SubscriptionProps) HandlerTimeout() time.Duration { return 30 * time.Second }
func (p *SubscriptionProps) Init()                       {}
func (p *SubscriptionProps) GrainTimeout() time.Duration { return 5 * time.Minute }
func (p *SubscriptionProps) GroupId() string             { return p.groupId }
func (*SubscriptionProps) Log(text string)               { fmt.Println(text) }
func (p *SubscriptionProps) OffsetStore() offset.Store   { return p.offsetStore }
func (p *SubscriptionProps) ProjectionGrainsFor(event any) []grains.ProjectionAddress {

	// TODO:
	return nil
}
func (p *SubscriptionProps) Stream() stream.Store { return p.stream }
func (p *SubscriptionProps) Terminate()           {}
