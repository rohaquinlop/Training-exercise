#!/bin/sh
sed 's/\x0/;/g' transactions > newtransactions.txt
sed 's/;;/\n/g' newtransactions.txt > transactions.csv