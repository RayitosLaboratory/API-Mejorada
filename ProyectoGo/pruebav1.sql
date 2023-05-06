create database dbpropruebav1;

use dbpropruebav1;

-- -----------------------------------------------------
-- Table `versiontres`.`clientes`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS clientes (
  `id` VARCHAR(15) NOT NULL,
  `nombre` VARCHAR(50) NOT NULL,
  `direccion` VARCHAR(100) NULL DEFAULT NULL,
  `telefono` VARCHAR(10) NULL DEFAULT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;

INSERT INTO clientes (id, nombre, direccion, telefono) VALUES
('1', 'Juan Perez', 'Calle 123, Ciudad', '5551234'),
('2', 'Maria Lopez', 'Avenida 456, Pueblo', '5555678'),
('3', 'Pedro Gomez', 'Plaza 789, Villa', '5559012'),
('4', 'Ana Rodriguez', 'Boulevard 1011, Colonia', '5553456'),
('5', 'Carlos Sanchez', 'Paseo 1213, Rancheria', '5557890');

select * from clientes;
-- -----------------------------------------------------
-- Table `versiontres`.`articulos`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS productos (
  `idProductos` INT NOT NULL AUTO_INCREMENT,
  `nombre` VARCHAR(100) NOT NULL,
  `precio` VARCHAR(100) NOT NULL,
  `medida` VARCHAR(15) NOT NULL,
  `stock` INT NULL DEFAULT '0',
  PRIMARY KEY (`idProductos`),
  UNIQUE INDEX `nombre` (`nombre` ASC) VISIBLE);

drop table productos;

INSERT INTO productos (nombre, precio, medida, stock) VALUES
('Leche', '35.00', 'Litro', 100),
('Pan', '12.50', 'Pieza', 500),
('Huevos', '25.00', 'Docena', 200),
('Queso', '85.00', 'Kilo', 50),
('Jugo', '20.00', 'Litro', 150);

select * from productos;
