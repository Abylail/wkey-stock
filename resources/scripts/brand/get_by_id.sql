select
    id,
    title,
    image,
    prosklad_id
from brands
where prosklad_id = $1
limit 1;