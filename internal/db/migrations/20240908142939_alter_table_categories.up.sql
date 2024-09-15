ALTER TABLE categories
    MODIFY COLUMN category_name VARCHAR(64) NOT NULL,
    ADD COLUMN group_id TINYINT NOT NULL AFTER category_name,
    ADD CONSTRAINT fk_group_id FOREIGN KEY (group_id) REFERENCES category_groups(id);