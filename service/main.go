package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/itgram/minion.encoding.protobuf/protobuf"
	"github.com/itgram/minion.persistence.cassandra/cassandra"
	"github.com/itgram/minion.persistence/persistence"
	"github.com/itgram/minion.system/system"
	"github.com/itgram/minion.system/system/strategy"

	"github.com/itgram/tracking_domain/vehicle"
	v "github.com/itgram/tracking_service/vehicle"
)

func main() {

	var _, cancel = context.WithCancel(
		context.Background())

	conn := cassandra.NewConnection(
		"minion_db",
		cassandra.WithHosts("localhost:9042"))

	var err = conn.Connect()
	if err != nil {
		fmt.Printf("Error to connect to tracking database: %v\n", err)
		return
	}

	defer conn.Close()

	var server = system.NewServer()

	var serializer = protobuf.NewSerializer()
	var firstBucketTime time.Time = time.Now() // NB:
	var partitionSize uint64 = 4

	system.RegisterCommandHandler(server,
		"vehicle",
		func() *vehicle.State { return &vehicle.State{} },
		v.CommandHandler,
		func(text string) { fmt.Println(text) },
		func() *persistence.Repository[*vehicle.State] {
			return persistence.NewRepository[*vehicle.State](
				partitionSize,
				conn.NewJournalStore(firstBucketTime, 20, serializer),
				conn.NewSnapshotStore(2, serializer))
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
