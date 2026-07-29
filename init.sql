CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (email, password_hash, full_name) VALUES (
    'test@example.com',
    '$2a$10$gWZx8TgYQPpS3K5KRHJTBuAeOUbTX/Btkuzgez0XNMfDs2RXj6MmG', //test123
    'Mohammad'
)
ON CONFLICT (email) DO NOTHING;

UPDATE users 
SET password_hash = '$2a$10$gWZx8TgYQPpS3K5KRHJTBuAeOUbTX/Btkuzgez0XNMfDs2RXj6MmG'
WHERE email = 'test@example.com';