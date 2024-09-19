CREATE TABLE events (
    id              CHAR(36) PRIMARY KEY NOT NULL,
    organization_id CHAR(36) NOT NULL,
    title           VARCHAR(255) NOT NULL,
    img             VARCHAR(255),
    description     TEXT,
    registrant      MEDIUMINT DEFAULT 0,
    max_registrant  MEDIUMINT,
    date            DATE NOT NULL,
    start_time      TIMESTAMP DEFAULT '2024-01-01 00:00:00',
    end_time        TIMESTAMP DEFAULT '2024-01-01 00:00:00',
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_organization FOREIGN KEY (organization_id) REFERENCES users(id)
) ENGINE = InnoDB;

CREATE TABLE event_pricings (
    id          BIGINT AUTO_INCREMENT PRIMARY KEY,
    event_id    CHAR(36) NOT NULL,
    event_type  VARCHAR(16),
    price       INT,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT  fk_event_pricings FOREIGN KEY (event_id) REFERENCES events(id)
) ENGINE = InnoDB;