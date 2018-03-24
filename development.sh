#!/bin/bash
# ex: ./development.sh serve
# windows users use it with gitbash(cmd.exe) as admin

function dep {
    rm -Rf vendor
    docker-compose up vendor
}

function serve {
    dep && up
}

function up {
    docker-compose up eventgo.local 
}

serve