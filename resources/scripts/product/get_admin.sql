select
    product.id,
    product.title,
    product.vendor_code,
    product.barcode,
    product.unit_name,
    productExt.category_id,
    category.code category_code,
    category.title_ru category_name,
    product.created_at,
    product.updated_at,
    product.additional_percent,
    product.price,
    productExt.description_ru,
    productExt.description_kz,
    productExt.count,
    brand.title brand_title
from products as product
    inner join products_ext as productExt on (productExt.product_id = product.id)
    left join categories as category on (category.id = productExt.category_id)
    inner join brands as brand on (product.brand_id = brand.prosklad_id)
order by product.title
offset $1
limit $2;