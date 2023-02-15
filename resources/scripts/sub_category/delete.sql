delete from sub_categories
where parent_id = $1 and code = $2;