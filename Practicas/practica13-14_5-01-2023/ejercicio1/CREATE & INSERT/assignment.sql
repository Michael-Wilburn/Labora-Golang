CREATE TABLE assignment (
    employee_id SERIAL REFERENCES employee(id),
    project_id VARCHAR(50)  REFERENCES project(id)
);

INSERT INTO assignment (employee_id, project_id) VALUES
(1, 'p3'),
(2, 'p1'),
(3, 'p2');
