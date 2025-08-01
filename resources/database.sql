DROP TABLE IF EXISTS bombus.species;

CREATE TABLE bombus.species (
  id INT AUTO_INCREMENT PRIMARY KEY,
  scientificName VARCHAR(50) NOT NULL UNIQUE,
  commonName VARCHAR(50) NOT NULL UNIQUE
);

INSERT INTO bombus.species (id, scientificName, commonName) VALUES
  (1, "Tetragosnisca Angustula" "Jataí"),
  (2, "Plebeia Sp.", "Mirim"),
  (3, "Melipona Quadrifasciata", "Mandaçaia"),
  (4, "Melipona Bicolor", "Uruçu"),
  (5, "Scaptotrigona Bipunctata" , "Tubuna"),
  (6, "Scaptotrigona Depilis", "Canudo");

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
);

INSERT INTO bombus.colmeias (id, colmeia_id, qr_code, species_id, starting_date, status_id) VALUES
  (1, '1', NULL, 1, '2025-05-08 14:30:00', 3),
  (2, '2', NULL, 2, '2025-05-08 14:30:00', 5),
  (3, NULL, NULL, 4, '2025-05-08 14:30:00', 5),
  (4, '4', NULL, 3, '2025-05-08 14:30:00', 1);
