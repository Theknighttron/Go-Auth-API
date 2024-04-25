-- name: CreateAuthor :execresult
CREATE TABLE Users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    Firstname VARCHAR(255) NOT NULL,
    Lastname VARCHAR(255) NOT NULL,
    Gender VARCHAR(10) NOT NULL,
    StudentId VARCHAR(50) NOT NULL,
    Course VARCHAR(255) NOT NULL,
    Level VARCHAR(50) NOT NULL,
    Email VARCHAR(255) NOT NULL
);

INSERT INTO register(
  id, Fistname, Lastname, Gender, StudentId, Course, Level, Email
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?
);
