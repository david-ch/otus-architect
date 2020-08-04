FROM postgres:12.2
RUN apt-get -q update && apt-get -qy install netcat
ADD https://raw.githubusercontent.com/eficode/wait-for/master/wait-for /wait-for
COPY *.sql /sql/
ENTRYPOINT chmod +x /wait-for && ./wait-for $POSTGRES_HOST:$POSTGRES_PORT -t 100 -- echo 'waiting completed' && sh -c "PGPASSWORD=$POSTGRES_PASSWORD psql -h $POSTGRES_HOST -p $POSTGRES_PORT -d $POSTGRES_DB -U $POSTGRES_USER < /sql/initialize-db.sql"
