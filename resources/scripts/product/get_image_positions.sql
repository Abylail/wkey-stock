select * from product_images
where product_id = $1 and position = any($2);