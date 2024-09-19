ALTER TABLE events
    ADD COLUMN start_time TIMESTAMP DEFAULT '2024-01-01 00:00:00' AFTER description,
    ADD COLUMN end_time TIMESTAMP DEFAULT '2024-01-01 00:00:00' AFTER start_time;

ALTER TABLE speakers
    ADD COLUMN event_id CHAR(36) AFTER id,
    ADD CONSTRAINT fk_event_speakers
        FOREIGN KEY (event_id) REFERENCES events(id),
    DROP FOREIGN KEY fk_agenda_speakers,
    DROP COLUMN agenda_id;

DROP TABLE IF EXISTS event_agendas;