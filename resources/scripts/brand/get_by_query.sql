select
    id,
    title,
    image,
    prosklad_id
from brands
where title ilike $1;