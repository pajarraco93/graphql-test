CREATE TABLE IF NOT EXISTS Albums (
    albumID int NOT NULL AUTO_INCREMENT,
    name VARCHAR(20) NOT NULL,
    year int,
    composedBy int,
    PRIMARY KEY (albumID),
    FOREIGN KEY (composedBy) REFERENCES Groups(groupID)
);