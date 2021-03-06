DROP TABLE IF EXISTS students;

CREATE TABLE students (
  id VARCHAR(32) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  age INTEGER NOT NULL
);

DROP TABLE IF EXISTS tests;

CREATE TABLE tests (
  id VARCHAR(32) PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

DROP TABLE IF EXISTS questions;

CREATE TABLE questions (
  id VARCHAR(32) PRIMARY KEY,
  answer VARCHAR(255) NOT NULL,
  question VARCHAR(255) NOT NULL,
  testId VARCHAR(32) NOT NULL,
  FOREIGN KEY (testId) REFERENCES tests(id)
);

DROP TABLE IF EXISTS enrollments;

CREATE TABLE enrollments (
  studentId VARCHAR(32) NOT NULL,
  testId VARCHAR(32) NOT NULL,
  FOREIGN KEY (studentId) REFERENCES students(id),
  FOREIGN KEY (testId) REFERENCES tests(id)
);
