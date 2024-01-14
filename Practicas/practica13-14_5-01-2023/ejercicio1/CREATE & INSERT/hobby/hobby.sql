CREATE TABLE hobby (
    employee_id SERIAL REFERENCES employee(id),
    hobby VARCHAR(255) NOT NULL
);

INSERT INTO hobby (employee_id, hobby) VALUES
(1, 'hobby1'),
(2, 'hobby2'),
(3, 'hobby3'),
(4, 'hobby4');
