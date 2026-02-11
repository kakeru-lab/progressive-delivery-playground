package skills

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

func CPUTempCelsius() (float64, error) {
	// Raspberry Pi path
	path := "/sys/class/thermal/thermal_zone0/temp"
	b, err := os.ReadFile(path)
	if err != nil {
		return 0, errors.New("cpu temp not available on this machine (expected Raspberry Pi).")
	}
	s := strings.TrimSpace(string(b))
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return v / 1000.0, nil
}
