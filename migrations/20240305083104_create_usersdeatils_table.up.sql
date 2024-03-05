CREATE TABLE IF NOT EXISTS userdetail (
    UserID SERIAL PRIMARY KEY,
    Address VARCHAR(255),
    Education VARCHAR(255),
    Gender VARCHAR(255),
    Department VARCHAR(255),
    ID INT REFERENCES usertable(ID) ON DELETE CASCADE
);