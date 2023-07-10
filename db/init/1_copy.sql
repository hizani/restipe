COPY "user" FROM '/db/data/user.csv' DELIMITER ',' CSV HEADER;
COPY recipe FROM '/db/data/recipe.csv' DELIMITER ',' CSV HEADER;
COPY ingredient FROM '/db/data/ingredient.csv' DELIMITER ',' CSV HEADER;
COPY ingredient_recipe FROM '/db/data/ingredient_recipe.csv' DELIMITER ',' CSV HEADER;
COPY step FROM '/db/data/step.csv' DELIMITER ',' CSV HEADER;
COPY rating FROM '/db/data/rating.csv' DELIMITER ',' CSV HEADER;

SELECT setval('ingredient_recipe_id_seq', (SELECT MAX(id) FROM ingredient_recipe)+1);
SELECT setval('ingredient_id_seq', (SELECT MAX(id) FROM ingredient)+1);
SELECT setval('recipe_id_seq', (SELECT MAX(id) FROM recipe)+1);
SELECT setval('step_id_seq', (SELECT MAX(id) FROM step)+1);
SELECT setval('user_id_seq', (SELECT MAX(id) FROM "user")+1);
SELECT setval('rating_id_seq', (SELECT MAX(id) FROM rating)+1);