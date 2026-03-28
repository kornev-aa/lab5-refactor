package storage

import (
    "encoding/json"
    "os"
)

type FileStorage struct {
    filePath string
}

type locationData struct {
    Latitude  float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
}

func NewFileStorage(path string) *FileStorage {
    return &FileStorage{filePath: path}
}

func (f *FileStorage) GetLatitude() (float64, error) {
    data, err := f.readFile()
    if err != nil {
        return 0, err
    }
    return data.Latitude, nil
}

func (f *FileStorage) GetLongitude() (float64, error) {
    data, err := f.readFile()
    if err != nil {
        return 0, err
    }
    return data.Longitude, nil
}

func (f *FileStorage) SaveLocation(lat, lon float64) error {
    data := locationData{Latitude: lat, Longitude: lon}
    file, err := os.Create(f.filePath)
    if err != nil {
        return err
    }
    defer file.Close()
    return json.NewEncoder(file).Encode(data)
}

func (f *FileStorage) readFile() (*locationData, error) {
    file, err := os.Open(f.filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    var data locationData
    if err := json.NewDecoder(file).Decode(&data); err != nil {
        return nil, err
    }
    return &data, nil
}
