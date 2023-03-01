# wkey-stock

> Сервис по управлению продуктами и категориями


## Клиентские роуты

```
/api/v1/stock/category/get - список категорий
/api/v1/stock/category/get?query=somename - список категорий по названию
/api/v1/stock/category/:par_code/sub/get - список подкатегорий по par_code (коду категории)
/api/v1/stock/category/:par_code/sub/get?query=somename - список подкатегорий по par_code (коду категории) и названию

/api/v1/stock/product/get?page=1 - список продуктов (надо сделать колличество)
/api/v1/stock/product/get?page=1&query=somename - список продуктов по названию (не работает)
```