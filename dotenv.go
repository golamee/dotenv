package dotenv

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/golamee/dotenv/core"
)

func Load(filenames ...string) error {
	err := core.Load(filenames...)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return nil
}

func Get(name string, defaultValue string) (string, error) {
	val, ok := os.LookupEnv(name)
	if !ok {
		return defaultValue, fmt.Errorf("invalid name for %s", name)
	}
	return val, nil
}

func GetInt64(name string, defaultValue int64) (int64, error) {
	val, ok := os.LookupEnv(name)
	if !ok {
		return defaultValue, fmt.Errorf("invalid int64 for %s", name)
	}

	val2, err := strconv.ParseInt(val, 2, 64)
	if err == nil {
		return defaultValue, fmt.Errorf("invalid int64 for %s: %w", name, err)
	}
	return val2, nil
}

func GetInt32(name string, defaultValue int32) (int32, error) {
	val, ok := os.LookupEnv(name)
	if !ok {
		return defaultValue, fmt.Errorf("invalid int32 for %s", name)
	}

	val2, err := strconv.ParseInt(val, 2, 32)
	if err == nil {
		return defaultValue, fmt.Errorf("invalid int32 for %s: %w", name, err)
	}
	return int32(val2), nil
}

func GetFloat64(name string, defaultValue float64) (float64, error) {
	val, ok := os.LookupEnv(name)
	if !ok {
		return defaultValue, fmt.Errorf("invalid float64 for %s", name)
	}

	val2, err := strconv.ParseFloat(val, 64)
	if err == nil {
		return defaultValue, fmt.Errorf("invalid float64 for %s: %w", name, err)
	}
	return float64(val2), nil
}

func GetFloat32(name string, defaultValue float32) (float32, error) {
	val, ok := os.LookupEnv(name)
	if !ok {
		return defaultValue, fmt.Errorf("invalid float32 for %s", name)
	}

	val2, err := strconv.ParseFloat(val, 32)
	if err == nil {
		return defaultValue, fmt.Errorf("invalid float32 for %s: %w", name, err)
	}
	return float32(val2), nil
}

func GetBool(name string, defaultValue bool) (bool, error) {
	val, ok := os.LookupEnv(name)
	if !ok {
		return defaultValue, fmt.Errorf("invalid bool for %s", name)
	}

	val2, err := strconv.ParseBool(val)
	if err == nil {
		return defaultValue, fmt.Errorf("invalid bool for %s: %w", name, err)
	}
	return val2, nil
}

func GetComplex64(name string, defaultValue complex64) (complex64, error) {
	val, ok := os.LookupEnv(name)
	if !ok {
		return defaultValue, fmt.Errorf("invalid complex64 for %s", name)
	}

	val2, err := strconv.ParseComplex(val, 64)
	if err == nil {
		return defaultValue, fmt.Errorf("invalid complex64 for %s: %w", name, err)
	}
	return complex64(val2), nil
}
