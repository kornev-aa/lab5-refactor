package weather

import (
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "net/http"
    "time"
    "github.com/kornev-aa/lab5-cache/pkg/cache"
)

type CurrentWeather struct {
    Temperature float64 `json:"temperature_2m"`
}

type WeatherResponse struct {
    Current CurrentWeather `json:"current"`
}

type WeatherService struct {
    httpClient *http.Client
    cache      cache.Cache
    cacheTTL   time.Duration
}

func NewWeatherService(cache cache.Cache, cacheTTL time.Duration) *WeatherService {
    return &WeatherService{
        httpClient: &http.Client{},
        cache:      cache,
        cacheTTL:   cacheTTL,
    }
}

func (s *WeatherService) GetWeather(lat, lon float64) (*WeatherResponse, error) {
    cacheKey := fmt.Sprintf("weather:%.4f:%.4f", lat, lon)

    if cached, found := s.cache.Get(cacheKey); found {
        var result WeatherResponse
        if err := json.Unmarshal(cached, &result); err == nil {
            return &result, nil
        }
    }

    params := fmt.Sprintf(
        "latitude=%f&longitude=%f&current=temperature_2m",
        lat, lon,
    )
    url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?%s", params)

    resp, err := s.httpClient.Get(url)
    if err != nil {
        return nil, errors.Join(errors.New("failed to fetch weather"), err)
    }
    defer resp.Body.Close()

    data, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, errors.Join(errors.New("failed to read response"), err)
    }

    var result WeatherResponse
    if err := json.Unmarshal(data, &result); err != nil {
        return nil, errors.Join(errors.New("failed to parse response"), err)
    }

    s.cache.Set(cacheKey, data, s.cacheTTL)

    return &result, nil
}
