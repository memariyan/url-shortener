CREATE TABLE IF NOT EXISTS url_shortener.url_data
(
    `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `key` VARCHAR(10) NOT NULL,
    `original_url` VARCHAR(255) NOT NULL,
    UNIQUE (`key`)
);