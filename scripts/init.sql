-- Create the dedicated user
CREATE USER shepherd_user WITH PASSWORD 'shepherd_password';

-- Create the database
CREATE DATABASE shepherd_db;

-- Grant permissions
GRANT ALL PRIVILEGES ON DATABASE shepherd_db TO shepherd_user;
-- For Postgres 15+, you also need to grant schema permissions:
\c shepherd_db
GRANT ALL ON SCHEMA public TO shepherd_user;