package weathermanager

import (
	"encoding/json"
	"os"
	"sync"
)

const storagePath = "C:\\Users\\user\\Desktop\\MyProject\\Go\\HelperBot\\UserDataStorage\\userCities.json"

var (
	userCities = make(map[int]string)
	mu         sync.Mutex
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

func saveCities() error {
	mu.Lock()
	defer mu.Unlock()

	data, err := json.MarshalIndent(userCities, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(storagePath, data, 0644)
}

func GetUserCity(userID int) (string, bool) {
	mu.Lock()
	defer mu.Unlock()
	city, ok := userCities[userID]
	return city, ok
}

func SetCity(userID int, city string) error {
	mu.Lock()
	userCities[userID] = city
	mu.Unlock()
	return saveCities()
}
