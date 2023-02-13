package image_event

import (
	"encoding/base64"
	"github.com/lowl11/lazyfile/fileapi"
	"github.com/lowl11/lazyfile/folderapi"
	"path/filepath"
)

func (event *Event) UploadCategoryIcon(categoryCode, name, buffer string) (string, error) {
	event.mutex.Lock()
	defer event.mutex.Unlock()

	imageExtension := filepath.Ext(name)

	fullPath := categoryIconPath + "/" + categoryCode + "/icon" + imageExtension

	// проверить существует ли папка категории
	if err := folderapi.Create(categoryIconPath, categoryCode); err != nil {
		return "", err
	}

	// декодируем из base64 в байты
	fileContent, err := base64.StdEncoding.DecodeString(buffer)
	if err != nil {
		return "", err
	}

	// создаем файл
	if err = fileapi.Create(fullPath, fileContent); err != nil {
		return "", err
	}

	return fullPath, nil
}

func (event *Event) GetCategoryIcon(categoryCode string) (string, error) {
	event.mutex.Lock()
	defer event.mutex.Unlock()
	return "", nil
}
