SELECT user_voted->>'$.tom' AS exist FROM polls WHERE title = 'boo';
