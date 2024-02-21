#!/bin/bash

set -e 

# Install PostgreSQL with enabled module
sudo dnf module list postgresql 
sudo dnf module enable postgresql:16 
sudo dnf -y install postgresql-server

# Initialize PostgreSQL database
sudo postgresql-setup --initdb

# Start PostgreSQL service
sudo systemctl start postgresql

# Enable PostgreSQL service to start on boot
sudo systemctl enable postgresql

# Create a PostgreSQL user
sudo -u postgres psql -c "CREATE USER postgres WITH PASSWORD 'postgres';"

# Create a PostgreSQL database
sudo -u postgres psql -c "CREATE DATABASE postgres;"

# Grant all privileges to the user on the database
sudo -u postgres psql -c "GRANT ALL PRIVILEGES ON DATABASE postgres TO postgres;"

# Update PostgreSQL to listen on all addresses
echo "listen_addresses = '*'" | sudo tee -a /var/lib/pgsql/data/postgresql.conf

# Allow incoming connections to PostgreSQL from all hosts
echo "host all all 0.0.0.0/0  md5" | sudo tee -a /var/lib/pgsql/data/pg_hba.conf

# Reload PostgreSQL for the changes to take effect
sudo systemctl reload postgresql

