
FROM postgres:latest


COPY DB.sql /docker-entrypoint-initdb.d/

ENV POSTGRES_DB=student_db
ENV POSTGRES_USER=your_username
ENV POSTGRES_PASSWORD=your_password
