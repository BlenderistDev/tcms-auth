CREATE TABLE IF NOT EXISTS user_session (
    id SERIAL PRIMARY KEY,
    user_id int NOT NULL,
    token varchar(512) NOT NULL UNIQUE,
    timestamp timestamp NOT NULL ,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
