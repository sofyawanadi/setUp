WITH roles AS (
    SELECT id AS role_id FROM roles WHERE const = 'administration'
),
user_data AS (
    SELECT id AS user_id FROM users WHERE username = 'MyUserAdmin'
)
INSERT INTO user_roles (user_id, role_id)
SELECT
    user_data.user_id,
    roles.role_id
FROM user_data, roles;
