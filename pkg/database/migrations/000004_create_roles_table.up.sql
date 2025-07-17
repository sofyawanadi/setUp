CREATE TABLE
    roles (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
        name VARCHAR(255) NOT NULL UNIQUE,
        const VARCHAR(255) NOT NULL UNIQUE,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
        created_by VARCHAR(255),
        updated_by VARCHAR(255),
        deleted_by VARCHAR(255),
        is_active BOOLEAN DEFAULT true
    );