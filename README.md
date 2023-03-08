# wkey-stock

> Сервис по управлению продуктами и категориями

## Админские роуты
Промоакции
```
/admin/api/v1/stock/promotion/get - список всех акций
/admin/api/v1/stock/promotion/get/:id - получить по id
/admin/api/v1/stock/promotion/code/:code - получить по code
/admin/api/v1/stock/promotion/create - создание акции
/admin/api/v1/stock/promotion/update - обновление акции
/admin/api/v1/stock/promotion/upload - обновление картинок акции
/admin/api/v1/stock/promotion/delete/:code - удаление акции
```

## Клиентские роуты

Категории
```
/api/v1/stock/category/get - список категорий
/api/v1/stock/category/get?query=somename - список категорий по названию
/api/v1/stock/category/:par_code/sub/get - список подкатегорий по par_code (коду категории)
/api/v1/stock/category/:par_code/sub/get?query=somename - список подкатегорий по par_code (коду категории) и названию
```

Продукты
```
/api/v1/stock/product/get?page=1 - список продуктов
/api/v1/stock/product/get?page=1&query=somename - список продуктов по названию (не работает)
```