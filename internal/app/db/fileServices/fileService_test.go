package fileservices

import (
	"os"
	"testing"

	"github.com/Alandres998/url-shortner/internal/config"
)

func TestGetURL(t *testing.T) {
	config.InitConfig()
	InitFileStorage()
	config.Options.FileStorage.Mode = os.O_RDWR
	config.Options.FileStorage.Path = "/tmp/short-url-db.json"
	// Добавляем данные для теста
	urlData = []URLData{
		{UUID: "1", ShortURL: "fsdfdsf", OriginalURL: "http://valhalla.com"},
		{UUID: "2", ShortURL: "zfdsfsdf", OriginalURL: "http://valhalla.ru"},
	}

	found := GetURL("fsdfdsf")
	if found == nil {
		t.Error("Ожидалось, что найдем существующий URL, но не найдено")
	}
	if found.UUID != "1" || found.ShortURL != "fsdfdsf" || found.OriginalURL != "http://valhalla.com" {
		t.Errorf("Неверные данные найденного URL: %+v", found)
	}

	// Тест 2: Проверяем отсутствующий URL
	notFound := GetURL("werwegjdf")
	if notFound != nil {
		t.Error("Карамба мы нашли то чего нет")
	}
}
