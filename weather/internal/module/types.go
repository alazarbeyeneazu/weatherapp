package module

import (
	"context"

	"github.com/alazarbeyeneazu/common/models"
)

type WeatherService interface {
	GetWeather(ctx context.Context, rq models.WeatherRequest) (models.WeatherResponse, error)
}
