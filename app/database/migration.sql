DROP TABLE IF EXISTS transactions, categories;

CREATE TABLE categories (
    ID int AUTO_INCREMENT PRIMARY KEY,
    UUID bigint NOT NULL,
    Name varchar(255) NOT NULL,
    Color varchar(8) NOT NULL
);

INSERT INTO categories (UUID, Name, Color) VALUES (1, 'Test category', 'FFFFFFFF');

CREATE TABLE transactions (
    ID int AUTO_INCREMENT PRIMARY KEY,
    UUID bigint NOT NULL,
    Name varchar(255) NOT NULL,
    Amount float NOT NULL,
    Date bigint NOT NULL,
    Type varchar(255) NOT NULL,
    CategoryUuid bigint REFERENCES categories(UUID)
);

INSERT INTO transactions (UUID, Name, Amount, Date, Type, CategoryUuid) VALUES (1, 'Test transaction', 255, 1653544523, 'expense', 1);