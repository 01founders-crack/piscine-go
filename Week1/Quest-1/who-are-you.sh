#! /bin/bash 

curl https://learn.01founders.co/assets/superhero/all.json | jq ' .[] | select( .id == 70) | .name' 