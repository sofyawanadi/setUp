WITH user_management_module AS (
    SELECT id FROM module WHERE const = 'settings'
)
INSERT INTO sub_module (name, module_id)
VALUES
    ('User Management', (SELECT id FROM user_management_module)),
    ('Master Module', (SELECT id FROM user_management_module)),
    ('Setting Access', (SELECT id FROM user_management_module));