#!/usr/bin/env bash

if [ ! -f /vagrant/.vagrant/default/virtualbox/bootstrapped ]
    then
        sudo apt-get update
        sudo apt-get install -y couchdb
        sudo cp /vagrant/vm/local.ini /etc/couchdb/
        touch /vagrant/.vagrant/default/virtualbox/bootstrapped
        sudo service couchdb start
    else
        sudo service couchdb restart
fi

