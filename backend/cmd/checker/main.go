// Copyright 2018-2020 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

// Package main provides an example to query the available endpoints of a server.
package main

import (
	"context"
	"fmt"
	"log"
	"sensor-simulator/gen/sensor_simulator/proto/simulator"
	"sensor-simulator/internal/pkg/endpoint/modbus"
	"sensor-simulator/internal/pkg/endpoint/opcua"
	"sensor-simulator/internal/pkg/service"
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
)

type lol interface{}

func main() {
	modbusServer, err := modbus.NewServer()
	if err != nil {
		log.Fatalf("unable to create modbus server: %v", err)
	}

	opcuaServer, err := opcua.NewServer()
	if err != nil {
		log.Fatalf("unable to create OPC UA server: %v", err)
	}

	simulatorService := service.NewSimulatorService(
		modbusServer,
		opcuaServer,
	)

	_, err = simulatorService.AddSensor(
		context.Background(),
		&simulator.AddSensorRequest{
			Simulator: &simulator.Simulator{
				Name:    "test",
				Address: 1,
				Base: &simulator.Base{
					Type: simulator.BaseType_BASE_TYPE_BEZIER,
					TypeData: &simulator.Base_Common{
						Common: &simulator.CommonBase{
							Generator: &simulator.Prng{
								Type: simulator.PrngType_PRNG_TYPE_XOSHIRO,
								Seed: 0,
							},
							MinValue:  10.0,
							MaxValue:  90.0,
							MinPeriod: durationpb.New(time.Second),
							MaxPeriod: durationpb.New(10 * time.Second),
						},
					},
				},
				Modifiers: []*simulator.Modifier{
					// {
					// 	Type: simulator.ModifierType_MODIFIER_TYPE_WHITE_NOISE,
					// 	TypeData: &simulator.Modifier_WhiteNoise{
					// 		WhiteNoise: &simulator.WhiteNoiseModifier{
					// 			Generator: &simulator.Prng{
					// 				Type: simulator.PrngType_PRNG_TYPE_XOSHIRO,
					// 				Seed: 0,
					// 			},
					// 			MaxValue: 1.0,
					// 		},
					// 	},
					// },
					{
						Type: simulator.ModifierType_MODIFIER_TYPE_HYSTERESIS,
						TypeData: &simulator.Modifier_Hysteresis{
							Hysteresis: &simulator.HysteresisModifier{
								Percentage: 40,
							},
						},
					},
				},
			},
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	_, err = simulatorService.Start(context.Background(), &simulator.StartRequest{})
	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(30 * time.Second)

	_, err = simulatorService.Stop(context.Background(), &simulator.StopRequest{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("stopped")

	// _, err = simulatorService.Start(context.Background(), &simulator.StartRequest{})
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// time.Sleep(30 * time.Second)

	// _, err = simulatorService.Stop(context.Background(), &simulator.StopRequest{})
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println("stopped second time")
}
