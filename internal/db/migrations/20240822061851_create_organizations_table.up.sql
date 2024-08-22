CREATE TABLE organizations (
    id          CHAR(36) PRIMARY KEY NOT NULL,
    name        VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    header_img  VARCHAR(255),
    profile_img VARCHAR(255),
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB;
