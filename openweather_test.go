package openweather_test

import (
	"fmt"
	"github.com/innotechdevops/openweather"
	"testing"
)

func TestOpenWeather_Execute(t *testing.T) {
	param := openweather.Parameter{
		Latitude:  openweather.Float32(33.44),
		Longitude: openweather.Float32(-94.04),
		Exclude: &[]string{
			openweather.ExcludeHourly,
			openweather.ExcludeDaily,
		},
		AppID: openweather.String("439d4b804bc8187953eb36d2a8c26a02"),
	}

	o := openweather.New()
	resp, err := o.Execute(param)

	if err.Error() != "Unauthorized" {
		t.Errorf("Error: %s", err)
	}
	fmt.Println(resp)
}
