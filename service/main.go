package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/itgram/green.encoding.protobuf/json"
	"github.com/itgram/green.encoding.protobuf/protobuf"
	"github.com/itgram/green.encoding/encoding"
	"github.com/itgram/green.fabric/fabric"
	"github.com/itgram/green.fabric/fabric/command"
	"github.com/itgram/green.persistence.esdb/esdb"
	"github.com/itgram/green/config"
	"github.com/itgram/tracking/domain/vehicle"

	"github.com/itgram/tracking/service/src"
	v "github.com/itgram/tracking/service/vehicle"
)

func main() {

	var ctx, cancel = context.WithCancel(
		context.Background())

	// load the application configuration
	var cfg = src.NewConfiguration()

	var err = config.LoadFrom(ctx, "local.env", cfg)
	if err != nil {
		fmt.Printf("Error failed to load the configuration: %v\n", err)
		return
	}

	conn := esdb.NewConnection(
		cfg.EventStore.Address,
		esdb.WithAuthentication(cfg.EventStore.Username, cfg.EventStore.Password),
		esdb.WithSerialization("json", func(encoding string) encoding.Serializer {

			if encoding == "json" {
				return json.NewSerializer()
			}

			return protobuf.NewSerializer()
		}),
	)

	err = conn.Connect()
	if err != nil {
		fmt.Printf("Error to connect to tracking database: %v\n", err)
		return
	}

	defer conn.Close()

	var server = fabric.NewServer(
		fabric.NewNodeConfigurtion("localhost", "my_cluster", 0, fabric.ConsulProvider))

	var kinds []fabric.ActorKind

	kinds = append(kinds,
		command.NewAggregateActorKind("vehicle", func(kind string) command.AggregateProps[*vehicle.State] { return v.NewAggregateProps(kind, conn) }))

	err = server.Start(kinds...)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	defer server.Shutdown(true)

	// Stop when a signal is sent.
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)

	<-c
	cancel()
}
