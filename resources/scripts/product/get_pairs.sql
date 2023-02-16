select
    pair.product_id,
    pair.sub_category_id,
    category.title_ru category_name,
    category.code category_code
from category_products as pair
    inner join sub_categories as category on (category.id = pair.sub_category_id)
where product_id = any($1);