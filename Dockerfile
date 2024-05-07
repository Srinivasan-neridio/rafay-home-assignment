
FROM mysql:latest

COPY ./contact-cli /

RUN chmod 0777 /contact-cli

ENV MYSQL_ROOT_PASSWORD="neridio"

ENV MYSQL_DATABASE="contacts"

