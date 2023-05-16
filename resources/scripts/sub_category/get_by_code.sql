select
    id,
    code,
    title_ru,
    title_kz,
    icon,
    status
from sub_categories
where code = $1;