#! /bin/bash 

curl https://learn.01founders.co/assets/superhero/all.json | jq '.[] | select(.id == 170) | .name, .powerstats.power, .appearance.gender' -r
