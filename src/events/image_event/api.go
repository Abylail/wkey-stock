package image_event

import (
	"encoding/base64"
	"github.com/lowl11/lazyfile/fileapi"
	"github.com/lowl11/lazyfile/folderapi"
	"path/filepath"
	"strconv"
)

func (event *Event) UploadCategoryIcon(categoryCode, name, buffer string) (string, error) {
	event.mutex.Lock()
	defer event.mutex.Unlock()

	imageExtension := filepath.Ext(name)

	fullPath := categoryIconPath + "/" + categoryCode + "/icon" + imageExtension

	// проверить существует ли папка категории
	if !folderapi.Exist(categoryIconPath) {
		// если нет создаем ее
		if err := folderapi.Create("/cdn", "category"); err != nil {
			return "", err
		}
	}

	// проверяем существует ли уже файл
	if fileapi.Exist(fullPath) {
		// удаляем его если существует
		if err := fileapi.Delete(fullPath); err != nil {
			return "", err
		}
	}

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

func (event *Event) UploadSubCategoryIcon(categoryCode, name, buffer string) (string, error) {
	event.mutex.Lock()
	defer event.mutex.Unlock()

	imageExtension := filepath.Ext(name)

	fullPath := subCategoryIconPath + "/" + categoryCode + "/icon" + imageExtension

	// проверить существует ли папка подкатегории
	if !folderapi.Exist(subCategoryIconPath) {
		// если нет создаем ее
		if err := folderapi.Create("/cdn", "subcategory"); err != nil {
			return "", err
		}
	}

	// проверяем существует ли уже файл
	if fileapi.Exist(fullPath) {
		// удаляем его если существует
		if err := fileapi.Delete(fullPath); err != nil {
			return "", err
		}
	}

	// проверить существует ли папка подкатегории
	if err := folderapi.Create(subCategoryIconPath, categoryCode); err != nil {
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

func (event *Event) UploadBrandIcon(brandID int, name, buffer string) (string, error) {
	event.mutex.Lock()
	defer event.mutex.Unlock()

	imageExtension := filepath.Ext(name)
	brandText := strconv.Itoa(brandID)

	fullPath := brandIconPath + "/" + brandText + "/icon" + imageExtension

	// проверить существует ли папка бренда
	if !folderapi.Exist(brandIconPath) {
		// если нет создаем ее
		if err := folderapi.Create("/cdn", "brand"); err != nil {
			return "", err
		}
	}

	// создаем папку под конкретный бренд
	if err := folderapi.Create(brandIconPath, brandText); err != nil {
		return "", err
	}

	// проверяем существует ли уже файл
	if fileapi.Exist(fullPath) {
		// удаляем его если существует
		if err := fileapi.Delete(fullPath); err != nil {
			return "", err
		}
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
