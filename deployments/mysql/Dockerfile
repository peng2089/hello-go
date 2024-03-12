FROM mysql:5.7
COPY ./init.sql  /tmp/sql/init.sql
RUN mv /tmp/sql/*.sql /docker-entrypoint-initdb.d
RUN rm -rf /tmp/sql