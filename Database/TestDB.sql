update user_data
set user_name = 'name'
where user_id = 1;
select * from user_data;
delete from user_data where user_name = 'CLOSE'
