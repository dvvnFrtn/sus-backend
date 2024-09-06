ALTER TABLE organizations
    DROP FOREIGN KEY fk_user_id,
    DROP COLUMN user_id;
