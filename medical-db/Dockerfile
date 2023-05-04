FROM mysql:8.0-debian
COPY my.cnf /etc/mysql/conf.d/mysqlutf8.cnf
ENV LANG C.UTF-8
ENV MYSQL_ROOT_PASSWORD=qwer
CMD ["mysqld", "--character-set-server=utf8", "--collation-server=utf8_unicode_ci"]