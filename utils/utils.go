package utils

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
)

// Nillable returns the default value if the given value is nil.
func Nillable[T any](value *T) T {
	if value == nil {
		var val T
		return val
	}

	return *value
}

// NillableWithInput returns a pointer to the given value or nil if it's input string is empty.
func NillableWithInput[T any](input string, element T) *T {
	if input == "" {
		return nil
	}

	return &element
}

// NillableString returns a pointer to the given string or nil if it's input string is empty.
func NillableString(input string) *string {
	if input == "" {
		return nil
	}

	return &input
}

// DownloadFile file and save it to disk.
func DownloadFile(url string, filepath string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		_ = out.Close()
	}(out)

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("could not download file from %s: %s", url, resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// RadiansToDegrees Converts Radians to Degrees
func RadiansToDegrees(radians float64) float64 {
	return radians * 180 / math.Pi
}

// DegreesToRadians Converts Degrees to Radians
func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}
