-- init.sql

-- Selecionando a base de dados
USE metadata;

-- Criação da tabela videos
CREATE TABLE IF NOT EXISTS videos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    cast TEXT,
    year INT,
    genre VARCHAR(100)
);

-- Inserindo alguns registros iniciais
INSERT INTO videos (title, description, cast, year, genre) VALUES
('The Shawshank Redemption', 'Two imprisoned men bond over a number of years, finding solace and eventual redemption through acts of common decency.', 'Tim Robbins, Morgan Freeman', 1994, 'Drama'),
('The Godfather', 'An organized crime dynasty\'s aging patriarch transfers control of his clandestine empire to his reluctant son.', 'Marlon Brando, Al Pacino', 1972, 'Crime, Drama'),
('The Dark Knight', 'When the menace known as the Joker emerges from his mysterious past, he wreaks havoc and chaos on the people of Gotham.', 'Christian Bale, Heath Ledger', 2008, 'Action, Crime, Drama'),
('Pulp Fiction', 'The lives of two mob hitmen, a boxer, a gangster\'s wife, and a pair of diner bandits intertwine in four tales of violência and redemption.', 'John Travolta, Uma Thurman', 1994, 'Crime, Drama');