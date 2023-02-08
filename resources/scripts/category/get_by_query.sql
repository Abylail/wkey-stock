select
    id,
    key,
    title,
    parent_id,
    position,
    items_count
from categories
where title ilike $1
order by title;