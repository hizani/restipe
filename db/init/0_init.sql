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
  "author" int references "user"(id) on delete cascade,
  "image" varchar(140) null
);

CREATE TABLE ingredient (
  "id" serial primary key,
  "name" varchar(64) 
);

CREATE TABLE ingredient_recipe (
  "id" serial primary key,
  "recipe_id" int references recipe(id) on delete cascade,
  "ingredient_id" int references ingredient(id) on delete cascade,
  "quantity" smallint,
  UNIQUE("recipe_id", "ingredient_id")
);


CREATE TABLE step (
  "id" serial primary key,
  "recipe_id" int references recipe(id) on delete cascade,
  "number" int,
  "description" text,
  "duration" interval second(0),
  "image" char(140) null
);

CREATE TABLE rating (
  "id" serial primary key, 
  "recipe_id" int references recipe(id) on delete cascade,
  "user_id" int references "user"(id) on delete cascade,
  "rating" smallint,
  UNIQUE(recipe_id, user_id)
)