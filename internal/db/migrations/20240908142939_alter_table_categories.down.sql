ALTER TABLE categories
    MODIFY COLUMN category_name VARCHAR(36), 
    DROP FOREIGN KEY fk_group_id,
    DROP COLUMN group_id;