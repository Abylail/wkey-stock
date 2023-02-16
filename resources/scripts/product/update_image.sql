insert into product_images (product_id, path, position, key)
values (:product_id, :path, :position, :key)
on conflict (key) do update set path = :path;