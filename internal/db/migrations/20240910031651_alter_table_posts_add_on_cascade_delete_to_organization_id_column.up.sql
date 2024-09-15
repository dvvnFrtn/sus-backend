ALTER TABLE posts
ADD CONSTRAINT fk_organization_cascade_delete
FOREIGN KEY (organization_id)
REFERENCES organizations(id)
ON DELETE CASCADE;
