update products_ext
set category_id = null
where product_id = $1;