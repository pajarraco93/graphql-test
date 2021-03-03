CREATE TABLE IF NOT EXISTS Songs (
    songID int NOT NULL AUTO_INCREMENT,
    name VARCHAR(20) NOT NULL,
    appearsIn int,
    PRIMARY KEY (songID),
    FOREIGN KEY (appearsIn) REFERENCES Albums(albumID)
);