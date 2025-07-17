CREATE TABLE notes(
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
        title TEXT NOT NULL,
        content TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
        deleted_at TIMESTAMP,
        created_by VARCHAR(255),
        updated_by VARCHAR(255),
        deleted_by VARCHAR(255),
        is_active BOOLEAN DEFAULT true
    );
