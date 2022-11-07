package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/itgram/green.fabric/fabric"

	"github.com/itgram/tracking/api/commands/vehicle"
)

func main() {

	var _, cancel = context.WithCancel(
		context.Background())

	var client = fabric.NewClient(
		fabric.NewNodeConfigurtion("localhost", "my_cluster", 0, fabric.ConsulProvider))

	var err = client.Start()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	defer client.Shutdown(true)

	var vehicleId = uuid.NewString()
	var result any

	fmt.Scanln()
	result, err = client.Request(vehicleId, "vehicle", &vehicle.RegisterVehicle{
		VehicleId: vehicleId,
		Model:     "Citreon C4",
	}, time.Second*10)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("sent ....", result)
	}

	fmt.Scanln()
	result, err = client.Request(vehicleId, "vehicle", &vehicle.AdjustMaxSpeedVehicle{
		VehicleId: vehicleId,
		MaxSpeed:  200,
	}, time.Second*10)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("sent ....", result)
	}

	fmt.Scanln()
	result, err = client.Request(vehicleId, "vehicle", &vehicle.AdjustMaxSpeedVehicle{
		VehicleId: vehicleId,
		MaxSpeed:  230,
	}, time.Second*10)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("sent ....", result)
	}

	// Stop when a signal is sent.
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)

	<-c
	cancel()
}
