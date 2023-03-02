select count(*)
from products as product
         inner join products_ext as productExt on (productExt.product_id = product.id)
where productExt.count > 10 and product.title ilike $1