CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(256) UNIQUE NOT NULL,
    password VARCHAR(512) NOT NULL,
    telegram_access_key VARCHAR(512) NOT NULL
);
