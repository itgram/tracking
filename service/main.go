package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/itgram/minion.encoding.protobuf/json"
	"github.com/itgram/minion.persistence.esdb/esdb"
	"github.com/itgram/minion.persistence/persistence"
	"github.com/itgram/minion.system/system"
	"github.com/itgram/minion.system/system/strategy"

	"github.com/itgram/tracking_domain/vehicle"
	v "github.com/itgram/tracking_service/vehicle"
)

func main() {

	var _, cancel = context.WithCancel(
		context.Background())

	conn := esdb.NewConnection(
		"esdb://127.0.0.1:2113?tls=false",
		esdb.WithAuthentication("admin", "changeit"))

	var err = conn.Connect()
	if err != nil {
		fmt.Printf("Error to connect to tracking database: %v\n", err)
		return
	}

	defer conn.Close()

	var server = system.NewServer()

	var serializer = json.NewSerializer()
	var partitionSize uint64 = 4

	system.RegisterCommandHandler(server,
		"vehicle",
		partitionSize,
		v.CommandHandler,
		func(text string) { fmt.Println(text) },
		func() persistence.Repository[*vehicle.State] {
			return persistence.NewRepository(
				conn.NewJournalStore(3, serializer),
				conn.NewSnapshotStore(2, serializer),
				&vehicle.State{})
		},
		func() strategy.Strategy { return strategy.NewSnapshotStrategy(strategy.WithMaxEventCount(10)) },
		func() time.Duration { return 30 * time.Second },
		func() time.Duration { return 30 * time.Second })

	err = server.Start(system.NewClusterConfigurtion("localhost", "my_cluster", 0))
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
