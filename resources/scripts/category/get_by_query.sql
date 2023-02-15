select
    id,
    code,
    title_ru,
    title_kz,
    icon,
    status
from categories
where title_ru ilike $1 or title_kz ilike $1 or code ilike $1
order by title_ru;