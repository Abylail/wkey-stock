select
    product.id,
    product.title,
    product.vendor_code,
    product.barcode,
    product.unit_name,
    product.category_id,
    product.category_name,
    product.created_at,
    product.updated_at,
    product.additional_percent,
    product.price,
    productExt.description_ru,
    productExt.description_kz,
    productExt.count
from products as product
         inner join products_ext as productExt on (productExt.product_id = product.id)
where id = $1;