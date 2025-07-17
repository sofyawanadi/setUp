CREATE TABLE
    user_roles (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
        role_id VARCHAR(255) NOT NULL,
        user_id VARCHAR(255) NOT NULL ,     
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
        deleted_at TIMESTAMP,
        created_by VARCHAR(255),
        updated_by VARCHAR(255),
        deleted_by VARCHAR(255),
        is_active BOOLEAN DEFAULT true
    );