package openweather

import (
	"errors"
	"fmt"
	"github.com/prongbang/callx"
	"net/http"
)

type OpenWeather interface {
	Execute(parameter Parameter) (string, error)
}

type openWeather struct {
	CallX callx.CallX
}

func (o *openWeather) Execute(parameter Parameter) (string, error) {
	path := fmt.Sprintf("/data/3.0/onecall?%s", parameter.ToQuery())
	resp := o.CallX.Get(path)
	if resp.Code == http.StatusOK {
		return string(resp.Data), nil
	}
	return "", errors.New(http.StatusText(resp.Code))
}

func New(args ...callx.CallX) OpenWeather {
	var c callx.CallX
	if len(args) > 0 {
		c = args[0]
	} else {
		cfg := callx.Config{
			BaseURL: "https://api.openweathermap.org",
			Timeout: 60,
		}
		c = callx.New(cfg)
	}
	return &openWeather{
		CallX: c,
	}
}
