select count(*) from products as product
    inner join products_ext as productExt on (productExt.product_id = product.id)
where productExt.category_id is null and product.title ilike $1