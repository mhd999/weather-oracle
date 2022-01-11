package weather

import (
	"context"

	"github.com/mhd999/weather-oracle/proto/go/services/weather"
)

type Server struct {
	weather.UnimplementedWeatherServer
}

func (s *Server) GetWeather(ctx context.Context, in *weather.GetWeatherRequest) (*weather.GetWeatherReply, error) {
	status := getClothesRecommendation()

	return &status, nil
}
