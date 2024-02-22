#!/bin/bash

# Copy application artifacts to /opt/app
sudo cp webapp /opt/app/

# Change ownership
sudo chown csye6225:csye6225 /opt/app/webapp

# Change permissions
sudo chmod 700 /opt/app/webapp