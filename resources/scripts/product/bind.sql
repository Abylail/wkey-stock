update products_ext
set category_id = $1
where product_id = any($2);