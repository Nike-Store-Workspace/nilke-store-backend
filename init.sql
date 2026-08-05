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


CREATE TABLE IF NOT EXISTS product_comments (
    id SERIAL PRIMARY KEY,
    product_id INT REFERENCES products(id) ON DELETE CASCADE,
    user_id BIGINT REFERENCES users(id) ON DELETE CASCADE, -- اتصال به جدول اصلی users
    title VARCHAR(255) NOT NULL,
    body_fa TEXT NOT NULL,
    body_en TEXT NOT NULL,
    rating INT CHECK (rating BETWEEN 1 AND 5),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE product_comments  ADD COLUMN title_fa VARCHAR(255) NOT NULL;
ALTER TABLE product_comments  ADD COLUMN title_en VARCHAR(255) NOT NULL;
ALTER TABLE product_comments  DROP COLUMN title;



CREATE TABLE IF NOT EXISTS banners(id SERIAL PRIMARY KEY, name VARCHAR(255),image VARCHAR(255) NOT NULL,lang VARCHAR(2) NOT NULL);
