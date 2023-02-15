update sub_categories
set status = :status
where parent_id = :parent_id and code = :code;