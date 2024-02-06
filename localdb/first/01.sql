CREATE TABLE `usuarios` (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    edad INT NOT NULL
);

INSERT INTO usuarios(id, nombre, email, clave) VALUES
(1, 'Juan', 'juan@gmail.com', 15),
(2, 'Jose', 'jose@gmail.com', 20),
(3, 'David', 'hermes@gmail.com', 25),
(4, 'Hermes', 'hermes@gmail.com', 10);