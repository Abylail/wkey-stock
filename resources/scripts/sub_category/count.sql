select count(*) from sub_categories
where parent_id = $1;