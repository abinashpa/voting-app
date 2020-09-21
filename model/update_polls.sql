UPDATE `polls`  SET user_voted = JSON_MERGE_PATCH(user_voted, '{"tom":"male"}') 
WHERE title = 'boo';
