CREATE TABLE employee (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255),
    department_id VARCHAR(50) REFERENCES department(id) ON DELETE CASCADE
);

INSERT INTO employee (name, address, department_id) VALUES
('Carlos', 'Direcc1', 'd1'),
('Pepe', 'Direcc2', 'd2'),
('Susana', 'Direcc3', 'd3'),
('Graciela', 'Direcc4', 'd4');
