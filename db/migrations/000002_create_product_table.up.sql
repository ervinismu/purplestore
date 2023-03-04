CREATE TABLE IF NOT EXISTS products (
   id serial PRIMARY KEY,
   name VARCHAR (50) UNIQUE NOT NULL,
   description VARCHAR (50) NOT NULL
)
