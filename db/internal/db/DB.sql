
CREATE DATABASE student_db;

\c student_db;

CREATE TABLE Universities (
    UniversityID SERIAL PRIMARY KEY,
    UniversityName VARCHAR(255),
    EducationForm VARCHAR(255),
    Course INT,
    GroupName VARCHAR(255),
    Specialty VARCHAR(255),
    Profile VARCHAR(255)
);

CREATE TABLE Passports (
    PassportID SERIAL PRIMARY KEY,
    PassportSeries VARCHAR(255),
    PassportIssueDate DATE
);


CREATE TABLE Students (
    StudentID SERIAL PRIMARY KEY,
    FullName VARCHAR(255),
    SNILS VARCHAR(255),
    BirthDate DATE,
    Email VARCHAR(255),
    PhoneNumber VARCHAR(255),
    TelegramUsername VARCHAR(255),
    UniversityID INT,
    PassportID INT,
    FOREIGN KEY (UniversityID) REFERENCES Universities(UniversityID),
    FOREIGN KEY (PassportID) REFERENCES Passports(PassportID)
);
