#!/bin/bash

set -e 

DB_USER="postgres"
DB_PASSWORD="postgres"
DB_NAME="postgres"

# Install PostgreSQL with enabled module
sudo dnf module list postgresql 
sudo dnf module enable postgresql:16 -y
sudo dnf install postgresql-server -y

# Initialize PostgreSQL database
sudo postgresql-setup --initdb

# Start PostgreSQL service
sudo systemctl start postgresql

# Enable PostgreSQL service to start on boot
sudo systemctl enable postgresql

# Create a PostgreSQL user
# sudo -u postgres psql -c "CREATE USER $DB_USER WITH PASSWORD '$DB_PASSWORD';"

# Change the password for the PostgreSQL default user 'postgres'
sudo -u postgres psql -c "ALTER USER $DB_USER WITH PASSWORD '$DB_PASSWORD';"

# Create a PostgreSQL database
sudo -u postgres psql -c "CREATE DATABASE $DB_NAME;"

# Grant all privileges to the user on the database
sudo -u postgres psql -c "GRANT ALL PRIVILEGES ON DATABASE $DB_NAME TO $DB_USER;"

# Update PostgreSQL to listen on all addresses
echo "listen_addresses = '*'" | sudo tee -a /var/lib/pgsql/data/postgresql.conf

# Allow incoming connections to PostgreSQL from all hosts
echo "host all all 0.0.0.0/0  md5" | sudo tee -a /var/lib/pgsql/data/pg_hba.conf

# Reload PostgreSQL for the changes to take effect
sudo systemctl reload postgresql

