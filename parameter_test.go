package openweather_test

import (
	"github.com/innotechdevops/openweather"
	"testing"
)

func TestParameter_ToQuery(t *testing.T) {
	expected := "lat=33.44&lon=-94.04&appid=439d4b804bc8187953eb36d2a8c26a02&exclude=hourly,daily"
	param := openweather.Parameter{
		Latitude:  openweather.Float32(33.44),
		Longitude: openweather.Float32(-94.04),
		Exclude: &[]string{
			openweather.ExcludeHourly,
			openweather.ExcludeDaily,
		},
		AppID: openweather.String("439d4b804bc8187953eb36d2a8c26a02"),
	}

	actual := param.ToQuery()

	if actual != expected {
		t.Errorf("Not match: %s", actual)
	}
}
