DROP TABLE project;
-- No es posible eliminar la tabla ya que hay que eliminar tambien en cascada todos los otros elementos de otra tabla que tenga id primary key de esta tabla como llave foranea
-- ERROR:  cannot drop table project because other objects depend on it