

FROM mysql

ADD feiraslivres.sql /docker-entrypoint-initdb.d/

RUN chown -R mysql:mysql /docker-entrypoint-initdb.d/

CMD ["mysqld", "--character-set-server=utf8mb4", "--collation-server=utf8mb4_unicode_ci"]


#FROM mysql:5.6
#COPY feiraslivres.sql /docker-entrypoint-initdb.d/test.sql

#FROM mysql:latest as builder
#
#RUN ["sed", "-i", "s/exec \"$@\"/echo \"not running $@\"/", "/usr/local/bin/docker-entrypoint.sh"]
#
## needed for intialization
#ENV MYSQL_ROOT_PASSWORD=root
#
#COPY feiraslivres.sql /docker-entrypoint-initdb.d/
#
#RUN ["/usr/local/bin/docker-entrypoint.sh", "mysqld", "--datadir", "/initialized-db"]
#
#FROM mysql:latest
#
#COPY --from=builder /initialized-db /var/lib/mysql