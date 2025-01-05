-- Postgres Table
CREATE TABLE IF NOT EXISTS books (
	id int4 NOT NULL,
	author text NULL,
	"name" text NULL,
	price float4 NULL,
	"year" int4 NULL,
	CONSTRAINT books_pkey PRIMARY KEY (id)
);

-- clear existing data 
TRUNCATE books;

-- Postgres Entries
INSERT INTO books (Id, Name, Author, Price, Year) VALUES (1, 'The Great Gatsby', 'F. Scott Fitzgerald', 10.99, 1925) ON CONFLICT DO NOTHING;
INSERT INTO books (Id, Name, Author, Price, Year) VALUES (2, '1984', 'George Orwell', 8.99, 1949) ON CONFLICT DO NOTHING;
INSERT INTO books (Id, Name, Author, Price, Year) VALUES (3, 'To Kill a Mockingbird', 'Harper Lee', 12.50, 1960) ON CONFLICT DO NOTHING;
INSERT INTO books (Id, Name, Author, Price, Year) VALUES (4, 'The Catcher in the Rye', 'J.D. Salinger', 9.75, 1951) ON CONFLICT DO NOTHING;
INSERT INTO books (Id, Name, Author, Price, Year) VALUES (5, 'Moby-Dick', 'Herman Melville', 15.20, 1851) ON CONFLICT DO NOTHING;
INSERT INTO books (Id, Name, Author, Price, Year) VALUES (6, 'Pride and Prejudice', 'Jane Austen', 7.99, 1813) ON CONFLICT DO NOTHING;
INSERT INTO books (Id, Name, Author, Price, Year) VALUES (7, 'The Hobbit', 'J.R.R. Tolkien', 14.99, 1937) ON CONFLICT DO NOTHING;
INSERT INTO books (Id, Name, Author, Price, Year) VALUES (8, 'War and Peace', 'Leo Tolstoy', 19.99, 1869) ON CONFLICT DO NOTHING;
INSERT INTO books (Id, Name, Author, Price, Year) VALUES (9, 'The Alchemist', 'Paulo Coelho', 11.45, 1988) ON CONFLICT DO NOTHING;
INSERT INTO books (Id, Name, Author, Price, Year) VALUES (10, 'The Road', 'Cormac McCarthy', 13.60, 2006) ON CONFLICT DO NOTHING;
