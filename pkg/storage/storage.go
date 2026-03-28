package storage

type LocationStorage interface {
    GetLatitude() (float64, error)
    GetLongitude() (float64, error)
    SaveLocation(lat, lon float64) error
}
