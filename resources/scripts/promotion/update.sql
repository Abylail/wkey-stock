UPDATE promotions
SET title_ru = :title_ru,
    title_kz = :title_kz,
    description_ru = :description_ru,
    description_kz = :description_kz
where code = :code;