select * from product_images
where product_id = any($1);