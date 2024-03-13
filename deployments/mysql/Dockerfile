FROM mysql:5.7
COPY ./sql/init.sql  /tmp/sql/init.sql
COPY ./sql/create_tables.sql  /tmp/sql/create_tables.sql
RUN mv /tmp/sql/*.sql /docker-entrypoint-initdb.d
RUN rm -rf /tmp/sql