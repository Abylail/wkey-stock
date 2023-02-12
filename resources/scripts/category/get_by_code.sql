select
    id,
    code,
    title_ru,
    title_kz,
    icon
from categories
where code = $1
limit 1;