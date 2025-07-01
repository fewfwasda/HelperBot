package weather

import (
	"encoding/json"
	"os"
)

func LoadCities() error {
	mu.Lock()
	defer mu.Unlock()

	info, err := os.Stat(storagePath)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	if info.Size() == 0 {
		return nil
	}
	data, err := os.ReadFile(storagePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &userCities)
}

func GetUserCity(userID int64) (string, bool) {
	mu.Lock()
	defer mu.Unlock()
	city, ok := userCities[userID]
	return city, ok
}

func SetCity(userID int64, city string) error {
	mu.Lock()
	userCities[userID] = city
	mu.Unlock()
	return saveCities()
}

func saveCities() error {
	mu.Lock()
	defer mu.Unlock()

	data, err := json.MarshalIndent(userCities, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(storagePath, data, 0644)
}
