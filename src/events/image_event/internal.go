package image_event

import (
	"encoding/base64"
	"github.com/google/uuid"
	"github.com/lowl11/lazyfile/fileapi"
	"github.com/lowl11/lazyfile/folderapi"
	"path/filepath"
	"strconv"
)

const (
	categoryIconPath    = "/cdn/category"
	subCategoryIconPath = "/cdn/subcategory"
	brandIconPath       = "/cdn/brand"
	productImagePath    = "/cdn/product"
)

func (event *Event) uploadProductImage(productID int, name, buffer string, position int) (string, error) {
	event.mutex.Lock()
	defer event.mutex.Unlock()

	imageExtension := filepath.Ext(name)
	productText := strconv.Itoa(productID)

	fileName := strconv.Itoa(position) + "_" + uuid.New().String()
	fullPath := productImagePath + "/" + productText + "/" + fileName + imageExtension

	// проверить существует ли папка бренда
	if !folderapi.Exist(productImagePath) {
		// если нет создаем ее
		if err := folderapi.Create("/cdn", "product"); err != nil {
			return "", err
		}
	}

	// создаем папку под конкретный бренд
	if err := folderapi.Create(productImagePath, productText); err != nil {
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
