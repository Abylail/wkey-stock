select
    id,
    code,
    title_ru,
    title_kz,
    icon,
    status
from categories
where code = $1
limit 1;