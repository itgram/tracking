package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/itgram/minion.system/system"

	"github.com/itgram/tracking_api/commands/vehicle"
)

func main() {

	var _, cancel = context.WithCancel(
		context.Background())

	var client = system.NewClient()

	var err = client.Start(system.NewClusterConfigurtion("localhost", "my_cluster", 0))
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	defer client.Shutdown(true)

	var vehicleId = uuid.New().String()
	var aggregateId = "vehicle/" + vehicleId

	fmt.Scanln()
	client.Request(aggregateId, "vehicle", &vehicle.RegisterVehicle{
		VehicleId: vehicleId,
		Model:     "Citreon C4",
	}, time.Second*10)
	fmt.Println("sent ....")

	fmt.Scanln()
	client.Request(aggregateId, "vehicle", &vehicle.AdjustMaxSpeedVehicle{
		VehicleId: vehicleId,
		MaxSpeed:  200,
	}, time.Second*10)
	fmt.Println("sent ....")

	fmt.Scanln()
	client.Request(aggregateId, "vehicle", &vehicle.AdjustMaxSpeedVehicle{
		VehicleId: vehicleId,
		MaxSpeed:  230,
	}, time.Second*10)
	fmt.Println("sent ....")

	// Stop when a signal is sent.
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)

	<-c
	cancel()
}
