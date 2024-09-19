CREATE TABLE event_agendas (
    id CHAR(36) PRIMARY KEY,
    event_id CHAR(36) NOT NULL,
    title VARCHAR(255),
    description TEXT,
    start_time TIMESTAMP NULL,
    end_time TIMESTAMP NULL,
    location VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_agendas FOREIGN KEY (event_id) REFERENCES events(id)
) ENGINE = InnoDB;

ALTER TABLE events
    DROP COLUMN start_time,
    DROP COLUMN end_time;

ALTER TABLE speakers
    DROP FOREIGN KEY fk_event_speakers,
    DROP COLUMN event_id,
    ADD COLUMN agenda_id CHAR(36) NOT NULL AFTER id,
    ADD CONSTRAINT fk_agenda_speakers
        FOREIGN KEY (agenda_id) REFERENCES event_agendas(id);
