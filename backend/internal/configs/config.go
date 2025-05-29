package configs

import (
	"os"
	"strconv"
	"strings"
	"sync"
)

type Config struct {
	Opcua     Opcua
	Modbus    Modbus
	Simulator Simulator
}

type Opcua struct {
	Host                   string
	Port                   int
	SimulatorNamespaceName string
}

type Modbus struct {
	Host           string
	Port           int
	TimeoutSeconds int
	MaxClients     int
}

type Simulator struct {
	Port        int
	Frequency   int
	LogsEnabled bool
}

var configInstance Config
var configOnce sync.Once

func GetConfig() Config {
	configOnce.Do(func() {
		configInstance = Config{
			Opcua: Opcua{
				Host:                   getEnv("OPCUA_HOST", "localhost"),
				Port:                   getEnvAsInt("OPCUA_PORT", 48400),
				SimulatorNamespaceName: getEnv("OPCUA_SIMULATOR_NAMESPACE_NAME", "SensorSimulators"),
			},
			Modbus: Modbus{
				Host:           getEnv("MODBUS_HOST", "localhost"),
				Port:           getEnvAsInt("MODBUS_PORT", 5502),
				TimeoutSeconds: getEnvAsInt("MODBUS_TIMEOUT_SECONDS", 30),
				MaxClients:     getEnvAsInt("MODBUS_MAX_CLIENTS", 5),
			},
			Simulator: Simulator{
				Port:        getEnvAsInt("SIMULATOR_HTTP_PORT", 8080),
				Frequency:   getEnvAsInt("SIMULATOR_FREQUENCY", 100),
				LogsEnabled: getEnvAsBool("SIMULATOR_LOGS_ENABLED", false),
			},
		}
	})

	return configInstance
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvList(key string, defaultVal []string) []string {
	if value, exists := os.LookupEnv(key); exists {
		return strings.Split(value, ",")
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}
