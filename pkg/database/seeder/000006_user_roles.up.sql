WITH roles AS (
    SELECT id FROM roles WHERE const = 'administration'
)
WITH user AS (
    SELECT id FROM user WHERE username = 'MyUserAdmin'
)
INSERT INTO permissions (user_id ,role_id ) VALUES
(SELECT id from user, SELECT id from roles)