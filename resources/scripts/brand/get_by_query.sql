select
    id,
    title,
    image
from brands
where title ilike $1;