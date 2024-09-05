CREATE TABLE organizers (
    id CHAR(36) PRIMARY KEY,            
    organization_id CHAR(36),            
    bank_name VARCHAR(255),   
    bank_account VARCHAR(255),       
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (id) REFERENCES users(id)
) ENGINE = InnoDB;