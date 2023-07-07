FROM postgres:alpine
ADD 0_init.sql /docker-entrypoint-initdb.d
ADD 1_copy.sql /docker-entrypoint-initdb.d
RUN chmod a+r /docker-entrypoint-initdb.d/*
EXPOSE 5432