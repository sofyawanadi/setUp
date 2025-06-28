CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users(
       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
        username VARCHAR(50) NOT NULL UNIQUE,
        email VARCHAR(100) NOT NULL UNIQUE,
        "password" VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
        deleted_at TIMESTAMP,
        created_by VARCHAR(255),
        updated_by VARCHAR(255),
        deleted_by VARCHAR(255),
        IsActive BOOLEAN DEFAULT true
    );
