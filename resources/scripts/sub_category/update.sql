update sub_categories
set title_ru = :title_ru,
    title_kz = :title_kz
where parent_id = :parent_id and code = :code;