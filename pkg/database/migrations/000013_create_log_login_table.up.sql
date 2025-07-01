CREATE TABLE log_logins(
       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
        email VARCHAR(100) NOT NULL,
        client_ip VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
