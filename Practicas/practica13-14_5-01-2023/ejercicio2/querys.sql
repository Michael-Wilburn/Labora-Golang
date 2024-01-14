INSERT INTO movie(id, name, created, stock, price) VALUES (1,1704398539,20,60);
-- ERROR:  INSERT has more target columns than expressions


INSERT INTO movie VALUES ('2','El señor de los anillos',20);
-- ERROR:  column "created" is of type date but expression is of type integer

INSERT INTO movie(id, name, price) VALUES (3,'Las dos torres',240000);
-- INSERT 0 1

INSERT INTO movie(id, name, created) VALUES (4,'El retorno del rey',1600000000);
-- ERROR:  column "created" is of type date but expression is of type integer