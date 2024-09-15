CREATE TABLE followers (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    organization_id CHAR(36) NOT NULL,
    follower_id CHAR(36) NOT NULL,
    followed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (organization_id) REFERENCES organizations(id) ON DELETE CASCADE,
    FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT uc_followers UNIQUE (organization_id, follower_id)
) ENGINE InnoDB;
