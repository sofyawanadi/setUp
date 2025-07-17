CREATE TABLE
    sub_modules (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
        module_id VARCHAR(255) NOT NULL,
        name VARCHAR(255) NOT NULL,
        const VARCHAR(255) NOT NULL,
        description TEXT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
        created_by VARCHAR(255),
        updated_by VARCHAR(255),
        deleted_by VARCHAR(255),
        is_active BOOLEAN DEFAULT true
    );