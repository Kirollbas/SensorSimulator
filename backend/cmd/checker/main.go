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
				Name:    "tempreature",
				Address: 11,
				Base: &simulator.Base{
					Type: simulator.BaseType_BASE_TYPE_BEZIER,
					TypeData: &simulator.Base_Common{
						Common: &simulator.CommonBase{
							Generator: &simulator.Prng{
								Type: simulator.PrngType_PRNG_TYPE_XOSHIRO,
								Seed: 2,
							},
							MinValue:  10.0,
							MaxValue:  90.0,
							MinPeriod: durationpb.New(time.Second),
							MaxPeriod: durationpb.New(10 * time.Second),
						},
					},
				},
				Modifiers: []*simulator.Modifier{},
			},
		},
	)
	if err != nil {
		fmt.Println(err)
	}

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
					// {
					// 	Type: simulator.ModifierType_MODIFIER_TYPE_HYSTERESIS,
					// 	TypeData: &simulator.Modifier_Hysteresis{
					// 		Hysteresis: &simulator.HysteresisModifier{
					// 			Percentage: 3,
					// 		},
					// 	},
					// },
					// {
					// 	Type: simulator.ModifierType_MODIFIER_TYPE_CONSTANT_OFFSET,
					// 	TypeData: &simulator.Modifier_ConstOffset{
					// 		ConstOffset: &simulator.ConstantOffsetModifier{
					// 			Offset: 2.0,
					// 		},
					// 	},
					// },
					// {
					// 	Type: simulator.ModifierType_MODIFIER_TYPE_INERTIA,
					// 	TypeData: &simulator.Modifier_Inertia{
					// 		Inertia: &simulator.InertitaModifier{
					// 			Value:  10,
					// 			Period: durationpb.New(time.Second),
					// 		},
					// 	},
					// },
					// {
					// 	Type: simulator.ModifierType_MODIFIER_TYPE_NONLINEAR_DEPENDENCE,
					// 	TypeData: &simulator.Modifier_NonLinearDependance{
					// 		NonLinearDependance: &simulator.NonLinearDependenceModifier{
					// 			Coefficient: 0.01,
					// 			Center:      50,
					// 		},
					// 	},
					// },
					// {
					// 	Type: simulator.ModifierType_MODIFIER_TYPE_PROGRESSING_OFFSET,
					// 	TypeData: &simulator.Modifier_ProgressingOffset{
					// 		ProgressingOffset: &simulator.ProgressingOffsetModifier{
					// 			Value:    0.2,
					// 			Interval: durationpb.New(time.Second),
					// 		},
					// 	},
					// },
					// {
					// 	Type: simulator.ModifierType_MODIFIER_TYPE_QUANTIZATION,
					// 	TypeData: &simulator.Modifier_Quantization{
					// 		Quantization: &simulator.QuantizationModifier{
					// 			Quant: 1.0,
					// 		},
					// 	},
					// },
					// {
					// 	Type: simulator.ModifierType_MODIFIER_TYPE_RANDOM_ADD_DASH,
					// 	TypeData: &simulator.Modifier_RandomAddDash{
					// 		RandomAddDash: &simulator.RandomAddDashModifier{
					// 			Generator: &simulator.Prng{
					// 				Type: simulator.PrngType_PRNG_TYPE_XOSHIRO,
					// 				Seed: 0,
					// 			},
					// 			MinAddValue:     5,
					// 			MaxAddValue:     10,
					// 			MinDashDuration: durationpb.New(time.Second),
					// 			MaxDashDuration: durationpb.New(2 * time.Second),
					// 			AvgPeriod:       durationpb.New(5 * time.Second),
					// 		},
					// 	},
					// },
					// {
					// 	Type: simulator.ModifierType_MODIFIER_TYPE_RANDOM_FIXED_DASH,
					// 	TypeData: &simulator.Modifier_RandomFixedDash{
					// 		RandomFixedDash: &simulator.RandomFixedDashModifier{
					// 			Generator: &simulator.Prng{
					// 				Type: simulator.PrngType_PRNG_TYPE_XOSHIRO,
					// 				Seed: 0,
					// 			},
					// 			Value:           90,
					// 			MinDashDuration: durationpb.New(500 * time.Millisecond),
					// 			MaxDashDuration: durationpb.New(2 * time.Second),
					// 			AvgPeriod:       durationpb.New(5 * time.Second),
					// 		},
					// 	},
					// },
					// {
					// 	Type: simulator.ModifierType_MODIFIER_TYPE_WHITE_NOISE,
					// 	TypeData: &simulator.Modifier_WhiteNoise{
					// 		WhiteNoise: &simulator.WhiteNoiseModifier{
					// 			Generator: &simulator.Prng{
					// 				Type: simulator.PrngType_PRNG_TYPE_XOSHIRO,
					// 				Seed: 0,
					// 			},
					// 			MaxValue: 2.0,
					// 		},
					// 	},
					// },
					{
						Type: simulator.ModifierType_MODIFIER_TYPE_DEPENDENCE,
						TypeData: &simulator.Modifier_Dependence{
							Dependence: &simulator.DependenceModifier{
								SimulatorName: "tempreature",
								Center:        50,
								Coefficient:   0.1,
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
