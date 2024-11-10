CREATE TABLE orders (
  id   varchar(36)  NOT NULL PRIMARY KEY,
  customer_name text    NOT NULL,
  description  text
);

CREATE TABLE items (
  id   varchar(36)  NOT NULL PRIMARY KEY,
  name text    NOT NULL,
  description  text,
  price  decimal(10,2)  NOT NULL,
);
