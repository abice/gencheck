#!/bin/sh

encrypt() {
  openssl des3 -in secrets.yml -out secrets.e.yml
}

decrypt() {
  openssl des3 -d -in secrets.e.yml -out secrets.yml
}

update() {
  openssl des3 -d -in secrets.e.yml -out secrets.yml
  drone secure --repo abice/gencheck --in secrets.yml
  rm secrets.yml
}

if [ "$1" = "d" ]; then
  decrypt
elif [ "$1" = "e" ]; then
  encrypt
elif [ "$1" = "u" ]; then
  update
else
  echo "Either 'e' or 'd'"
fi
