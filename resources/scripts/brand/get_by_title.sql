select
    id,
    title,
    image
from brands
where title = $1
limit 1;