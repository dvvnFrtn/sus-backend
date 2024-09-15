CREATE TABLE speakers (
    id          CHAR(36) PRIMARY KEY NOT NULL,
    name        VARCHAR(255) NOT NULL,
    title       VARCHAR(255),
    img         VARCHAR(255),
    description TEXT,
    event_id    CHAR(36),
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_event_speakers FOREIGN KEY (event_id) REFERENCES events(id)
) ENGINE = InnoDB;

CREATE TABLE activities (
    id CHAR(36) NOT NULL PRIMARY KEY,
    organization_id CHAR(36) NOT NULL,
    title VARCHAR(255),
    note TEXT NOT NULL,
    start_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    end_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (organization_id) REFERENCES organizations(id)
) ENGINE = InnoDB;