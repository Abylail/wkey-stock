select
    id,
    code,
    title_ru,
    title_kz,
    icon,
    status
from sub_categories
where parent_id = $1 and code = $2;