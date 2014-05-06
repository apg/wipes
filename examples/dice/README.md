# Dice example

This example "rolls" 100 dice a second, and counts how many
times each number appears. It then outputs JSON specifying
the value, the number (the service key) and a timestamp. This
is then plotted in a set of sparklines.

## Build it

    $ go build

## Run it

    $ ./dice | wipes -addr
    $ open http://localhost:8080/


