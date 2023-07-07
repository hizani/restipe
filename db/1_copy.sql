COPY "user" FROM '/db/data/user.csv' DELIMITER ',' CSV HEADER;
COPY recipe FROM '/db/data/recipe.csv' DELIMITER ',' CSV HEADER;
COPY ingredient FROM '/db/data/ingredient.csv' DELIMITER ',' CSV HEADER;
COPY step FROM '/db/data/step.csv' DELIMITER ',' CSV HEADER;