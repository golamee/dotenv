package dotenv

import (
	"log"
	"os"
	"strconv"

	"github.com/golamee/dotenv/core"
)

func Load(filenames ...string) {

	err := core.Load(filenames...)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func Get(name string, defaultValue string) (string, bool) {

	val, isSuccess := os.LookupEnv(name)

	if !isSuccess {

		return defaultValue, false
	}

	return val, true

}

func GetInt64(name string, defaultValue int64) (int64, bool) {

	val, isSuccess := os.LookupEnv(name)

	if !isSuccess {

		return defaultValue, false
	}

	val2, err := strconv.ParseInt(val, 2, 64)

	if err == nil {

		return defaultValue, false
	}

	return val2, true

}

func GetInt32(name string, defaultValue int32) (int32, bool) {

	val, isSuccess := os.LookupEnv(name)

	if !isSuccess {

		return defaultValue, false
	}

	val2, err := strconv.ParseInt(val, 2, 32)

	if err == nil {

		return defaultValue, false
	}

	return int32(val2), true

}

func GetFloat64(name string, defaultValue float64) (float64, bool) {

	val, isSuccess := os.LookupEnv(name)

	if !isSuccess {

		return defaultValue, false
	}

	val2, err := strconv.ParseFloat(val, 64)

	if err == nil {

		return defaultValue, false
	}

	return float64(val2), true

}

func GetFloat32(name string, defaultValue float32) (float32, bool) {

	val, isSuccess := os.LookupEnv(name)

	if !isSuccess {

		return defaultValue, false
	}

	val2, err := strconv.ParseFloat(val, 32)

	if err == nil {

		return defaultValue, false
	}

	return float32(val2), true

}

func GetBool(name string, defaultValue bool) (bool, bool) {

	val, isSuccess := os.LookupEnv(name)

	if !isSuccess {

		return defaultValue, false
	}

	val2, err := strconv.ParseBool(val)

	if err == nil {

		return defaultValue, false
	}

	return val2, true

}

func GetComplex64(name string, defaultValue complex64) (complex64, bool) {

	val, isSuccess := os.LookupEnv(name)

	if !isSuccess {

		return defaultValue, false
	}

	val2, err := strconv.ParseComplex(val, 64)

	if err == nil {

		return defaultValue, false
	}

	return complex64(val2), true

}
