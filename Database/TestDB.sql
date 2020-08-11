update user_data set user_id = 1 where user_name = 'Mirel Costin';
select * from user_data;
delete from user_data where user_name = 'CLOSE';
SELECT user_id, LAST_VALUE(user_id) OVER (ORDER BY user_id) last_id FROM user_data;
