CREATE TABLE categories (
    id              CHAR(36) PRIMARY KEY NOT NULL,
    category_name   VARCHAR(36) NOT NULL,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE = InnoDB;

CREATE TABLE user_categories (
    id          INT AUTO_INCREMENT PRIMARY KEY,
    category_id CHAR(36) NOT NULL,
    user_id     CHAR(36) NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES categories(id),
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)
) ENGINE = InnoDB;