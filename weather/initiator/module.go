package initiator

import (
	md "github.com/alazarbeyeneazu/weather/internal/module"
	"github.com/alazarbeyeneazu/weather/platform"
	"go.uber.org/zap"
)

type module struct {
	weatherModule md.WeatherService
}

func initModule(_ persistence, weatherAPI platform.WeatherAPI, log *zap.Logger) module {
	return module{
		weatherModule: md.NewService(weatherAPI, log),
	}
}
