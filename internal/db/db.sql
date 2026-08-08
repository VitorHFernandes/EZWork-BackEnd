CREATE DATABASE EZWork;

USE EZWork;
CREATE TABLE tb_todo_list(
  ID int AUTO_INCREMENT PRIMARY KEY,
  userID int,
  title VARCHAR(30) NOT NULL,
  description VARCHAR(255) DEFAULT NULL,
  dtInit VARCHAR(8) DEFAULT "00:00:00",
  dtEnd VARCHAR(8) DEFAULT "00:00:00",
  isCompleted BOOLEAN DEFAULT FALSE
);
INSERT INTO tb_todo_list(userID, title, description,dtInit,dtEnd) VALUES(1, '☕ Café da manhã', '100gr de Banana + 3 ovos inteiros + 30gr de aveia.', '08:30:00', '10:30:00');

CREATE TABLE tb_users(
  ID INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(30) NOT NULL,
  jobID INT NOT NULL,
  userLevel INT DEFAULT 3,
  email VARCHAR(30) NOT NULL,
  pass VARCHAR(255) NOT NULL,
  createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
  lastLogin DATETIME DEFAULT NULL
);
INSERT INTO tb_users(name, jobID, userLevel, email, pass) VALUES('Vítor H. Fernandes', 1, 0, 'vitor@eycon.com.br', '123456789');

CREATE TABLE tb_user_levels(
  ID INT AUTO_INCREMENT PRIMARY KEY,
  userLevel VARCHAR(30) NOT NULL
);
INSERT INTO tb_user_levels(userLevel) VALUES('Administrador');
INSERT INTO tb_user_levels(userLevel) VALUES('Usuário');

CREATE TABLE tb_user_job(
  ID INT AUTO_INCREMENT PRIMARY KEY,
  user_job_title VARCHAR(30) NOT NULL
);
INSERT INTO tb_user_job(user_job_title) VALUES('Software Engineer');

