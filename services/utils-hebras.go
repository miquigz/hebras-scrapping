package services

import (
	"encoding/json"
	"hebras-scrapping/models"
	"os"
	"strconv"
	"strings"
	"sync"
)

type HebrasUtils struct{}

var (
	utils HebrasUtils
	once  sync.Once
)

func NewHebrasUtils() *HebrasUtils {
	once.Do(func() {
		utils = HebrasUtils{}
	})
	return &utils
}

// FormatTeaBlendPrice recibe un string con el precio de un blend de te y lo formatea a un int
// Example input: "Desde $8.500,00"
func (hu *HebrasUtils) FormatTeaBlendPrice(text string) (int, error) {
	text = strings.Split(text, ",")[0] //Omito decimales
	text = strings.ReplaceAll(text, "Desde", "")
	text = strings.ReplaceAll(text, "$", "")
	text = strings.ReplaceAll(text, " ", "") //Remuevo espacios
	text = strings.ReplaceAll(text, ".", "") //Remuevo puntos

	price, err := strconv.Atoi(text)
	if err != nil {
		return 0, err
	}
	return price, nil
}

// SaveDataToFile recibe un slice de HebrasTea y lo guarda en un archivo JSON en la carpeta hugo/data
func (hu *HebrasUtils) SaveDataToFile(teaHebras []models.HebrasTea) error {
	data := map[string][]models.HebrasTea{
		"teas": teaHebras,
	}
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile("hugo/data/teas.json", file, 0644)
	if err != nil {
		return err
	}
	return nil
}
