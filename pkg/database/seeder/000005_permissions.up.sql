WITH roles AS (
    SELECT id FROM roles WHERE const = 'administration'
)
WITH module AS (
    SELECT id FROM module WHERE const = 'settings'
)
INSERT INTO permissions (role_id, module_id, is_view, is_create, is_update, is_delete, created_at, updated_at) VALUES
(SELECT id from roles, SELECT id from module, true, false, false, false, NOW(), NOW()),