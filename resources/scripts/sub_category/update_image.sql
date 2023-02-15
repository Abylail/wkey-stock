update sub_categories
set icon = :icon
where parent_id = :parent_id and code = :code;