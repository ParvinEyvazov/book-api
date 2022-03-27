-- create book table
CREATE TABLE `Book` (
    id varchar(255) NOT NULL,
    title varchar(255),
    description varchar(255),
    publication_date varchar(255),
    authorIDs JSON,
    PRIMARY KEY (id)
);

-- create author table
CREATE TABLE `Author` (
    id varchar(255) NOT NULL,
    name varchar(255),
    surname varchar(255),
    PRIMARY KEY (id)
);

-- insert data into author table
INSERT INTO `Author` (`id`, `name`, `surname`)
VALUES ('1', 'Lennifer', 'Jopez'),
    ('2', 'Mia', 'Jessica');

-- insert data into book table
INSERT INTO `Book` (
        `id`,
        `title`,
        `description`,
        `publication_date`,
        `authorIDs`
    )
VALUES (
        '1',
        'The Lord of the Rings',
        'The Lord of description.',
        '1954-01-01',
        '["1"]'
    );