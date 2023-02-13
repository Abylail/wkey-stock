update products_ext
set description_ru = :description_ru,
    description_kz = :description_kz
where product_id = :id;