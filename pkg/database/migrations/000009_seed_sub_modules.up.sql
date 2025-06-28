WITH user_management_module AS (
    SELECT id FROM modules WHERE const = 'settings'
)
INSERT INTO sub_modules (name,const, module_id)
VALUES
    ('User Management','user_management' ,(SELECT id FROM user_management_module)),
    ('Master Module','master_module' ,(SELECT id FROM user_management_module)),
    ('Setting Access','setting_access' ,(SELECT id FROM user_management_module));