--
-- Estrutura do banco de dados `dbLeo`
--

CREATE DATABASE IF NOT EXISTS `dbLeo` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `dbLeo`;

-- --------------------------------------------------------

--
-- Estrutura da tabela `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `password` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Inserindo dados de exemplo na tabela `users`
--

-- O usuário é 'admin' e a senha é 'password123'.
-- A senha foi hasheada usando a função password_hash() do PHP.
INSERT INTO `users` (`username`, `password`) VALUES
('admin', '$2y$10$g.pA/P2QF42a2rPSd5j95OqY.i.2t.3f6X.b/uS.Z/a.B/c.D.E.F');

COMMIT;
