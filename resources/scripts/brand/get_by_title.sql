select
    id,
    title,
    image,
    prosklad_id
from brands
where title = $1
limit 1;