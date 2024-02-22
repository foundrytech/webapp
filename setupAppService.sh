#!/bin/bash

# Copy webapp.service to /etc/systemd/system
sudo cp webapp.service /etc/systemd/system/

# Configure systemd service to start on boot
sudo systemctl daemon-reload
sudo systemctl enable webapp