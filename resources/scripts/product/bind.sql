insert into category_products (product_id, sub_category_id)
values ($1, $2)
on conflict do nothing;