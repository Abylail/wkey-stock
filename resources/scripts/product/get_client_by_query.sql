select
    product.id,
    product.title,
    product.vendor_code,
    product.price,
    productExt.count
from products as product
    inner join products_ext as productExt on (productExt.product_id = product.id)
where productExt.count > 10 and product.title ilike $3
order by product.title
offset $1
    limit $2;