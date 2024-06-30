#!/bin/bash
export FLASK_APP=app.py
flask run -h $LISTEN_IP -p $LISTEN_PORT