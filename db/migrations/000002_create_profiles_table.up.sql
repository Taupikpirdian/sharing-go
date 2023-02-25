CREATE TABLE IF NOT EXISTS profiles(
   id serial PRIMARY KEY,
   userId INT NOT NULL,
   jobId INT NOT NULL,
   unitId INT NOT NULL,
   placeOfBirth VARCHAR (150) NULL,
   dateOfBirth date NULL,
   gender VARCHAR (2) NULL,
   subDistrictId INT NULL,
   villageId INT NULL,
   address TEXT NULL,
   postalCode VARCHAR (8) NULL,
   updatedAt TIMESTAMP NULL,
   createdAt TIMESTAMP NULL
);