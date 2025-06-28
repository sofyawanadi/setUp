WITH roles AS (
    SELECT id AS role_id FROM roles WHERE const = 'administration'
),
modules AS (
    SELECT id AS module_id FROM modules WHERE const = 'settings'
)
INSERT INTO permissions (
    role_id, module_id, is_view, is_create, is_update, is_delete, created_at, updated_at
)
SELECT
    roles.role_id,
    modules.module_id,
    true, false, false, false,
    NOW(), NOW()
FROM roles, modules;