CREATE TABLE "user" (
  "id" serial primary key,
  "name" varchar(64),
  "login" varchar(64) unique,
  "password_hash" bytea

);

CREATE TABLE recipe ( 
  "id" serial primary key,
  "name" varchar(64),
  "description" text,
  "author" int references "user"(id) on delete cascade
);

CREATE TABLE ingredient (
  "id" serial primary key,
  "recipe_id" int references recipe(id) on delete cascade,
  "name" varchar(64) ,
  "quantity" smallint
);


CREATE TABLE step (
  "id" serial primary key,
  "recipe_id" int references recipe(id) on delete cascade,
  "number" int,
  "description" text,
  "duration" interva second(0)
);

