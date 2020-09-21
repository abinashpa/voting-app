UPDATE `polls` 
SET
    user_voted = JSON_MERGE_PATCH(user_voted, '{"userid":"option"}')
WHERE
    title = 'boo'
;