CREATE TABLE
    permissions (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
        role_id VARCHAR(255) NOT NULL,
        module_id VARCHAR(255) NOT NULL ,
        is_view BOOLEAN DEFAULT false,
        is_create BOOLEAN DEFAULT false,
        is_update BOOLEAN DEFAULT false,
        is_delete BOOLEAN DEFAULT false,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP,
        created_by VARCHAR(255),
        updated_by VARCHAR(255),
        deleted_by VARCHAR(255),
        IsActive BOOLEAN DEFAULT false
    );