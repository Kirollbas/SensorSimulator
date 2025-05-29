// Copyright 2018-2020 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

// Package main provides an example to query the available endpoints of a server.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sensor-simulator/internal/pkg/dto"
	"sensor-simulator/internal/pkg/endpoint/modbus"
	"sensor-simulator/internal/pkg/endpoint/opcua"
	"sensor-simulator/internal/pkg/service"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("./.env"); err != nil {
		fmt.Printf("No .env file found %s\n", err)
	}
}

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

	depend := `{
  "name": "test_1",
  "address": 1,
  "base": {
    "type": "BASE_TYPE_BEZIER",
    "data": {
      "generator": {
        "type": "PRNG_TYPE_PCG",
        "seed": 5
      },
      "min_value": 10.0,
      "max_value": 90.0,
      "min_period": "2s",
      "max_period": "4s"
    }
  },
  "modifiers": []
}
	`

	example := `{
  "name": "test",
  "address": 10,
  "base": {
    "type": "BASE_TYPE_BEZIER",
    "data": {
      "generator": {
        "type": "PRNG_TYPE_PCG",
        "seed": 1
      },
      "min_value": 10.0,
      "max_value": 90.0,
      "min_period": "2s",
      "max_period": "4s"
    }
  },
  "modifiers": [
    {
      "type": "MODIFIER_TYPE_DEPENDENCE",
      "data": {
	  	"simulator_name": "test_1",
	  	"coefficient": 0.1,
		"center": 50.0
      }
    }
  ]
}
	`
	var simulatorDTO1 dto.Simulator
	err = json.Unmarshal([]byte(depend), &simulatorDTO1)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = simulatorService.AddSensor(simulatorDTO1)

	var simulatorDTO dto.Simulator
	err = json.Unmarshal([]byte(example), &simulatorDTO)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = simulatorService.AddSensor(simulatorDTO)

	simulatorService.Start(context.Background())

	time.Sleep(30 * time.Second)

	simulatorService.Stop(context.Background())
}
