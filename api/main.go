package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	fmt.Scanln()
	client.Request("vehicle/90", "vehicle", &vehicle.RegisterVehicle{
		VehicleId: "90",
		Model:     "Citreon C4",
	}, time.Second*10)
	fmt.Println("sent ....")

	fmt.Scanln()
	client.Request("vehicle/90", "vehicle", &vehicle.AdjustMaxSpeedVehicle{
		VehicleId: "90",
		MaxSpeed:  200,
	}, time.Second*10)
	fmt.Println("sent ....")

	fmt.Scanln()
	client.Request("vehicle/90", "vehicle", &vehicle.AdjustMaxSpeedVehicle{
		VehicleId: "90",
		MaxSpeed:  230,
	}, time.Second*10)
	fmt.Println("sent ....")

	// Stop when a signal is sent.
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)

	<-c
	cancel()
}
