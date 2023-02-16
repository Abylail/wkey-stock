select
    pair.product_id,
    pair.sub_category_id,
    subCategory.title_ru sub_category_name,
    subCategory.code sub_category_code,
    category.code category_code,
    category.title_ru category_name
from category_products as pair
    inner join sub_categories as subCategory on (subCategory.id = pair.sub_category_id)
    inner join categories as category on (category.id = subCategory.parent_id)
where product_id = any($1);