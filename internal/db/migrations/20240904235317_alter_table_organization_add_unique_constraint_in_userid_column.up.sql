ALTER TABLE organizations
    ADD COLUMN user_id VARCHAR(36) NOT NULL UNIQUE,
    ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id);
