DROP TABLE IF EXISTS bombus.status;

CREATE TABLE bombus.status (
  id INT AUTO_INCREMENT PRIMARY KEY,
  value VARCHAR(50) NOT NULL UNIQUE
);

INSERT INTO bombus.status (id, value) VALUES
  (1, "HoneyReady"),
  (2, "Induzida"),
  (3, "Developing"),
  (4, "Pet"),
  (5, "Empty");

DROP TABLE IF EXISTS bombus.species;

CREATE TABLE bombus.species (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(50) NOT NULL UNIQUE
);

INSERT INTO bombus.species (id, name) VALUES
  (1, "TetragosniscaAngustula"),
  (2, "PlebeiaSp"),
  (3, "MeliponaQuadrifasciata"),
  (4, "MeliponaBicolor"),
  (5, "ScaptotrigonaBipunctata"),
  (6, "ScaptotrigonaDepilis");

DROP TABLE IF EXISTS bombus.colmeias;

CREATE TABLE bombus.colmeias (
  id INT NOT NULL,
  colmeia_id VARCHAR(10),
  qr_code BLOB,
  species_id INT NOT NULL,
  starting_date DATETIME,
  status_id INT NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (species_id) REFERENCES bombus.species(id),
  FOREIGN KEY (status_id) REFERENCES bombus.status(id)
);

INSERT INTO bombus.colmeias (id, colmeia_id, qr_code, species_id, starting_date, status_id) VALUES
  (1, '1', NULL, 1, '2025-05-08 14:30:00', 3),
  (2, '2', NULL, 2, '2025-05-08 14:30:00', 5),
  (3, NULL, NULL, 4, '2025-05-08 14:30:00', 5),
  (4, '4', NULL, 3, '2025-05-08 14:30:00', 1);
