# tracking

## event store

eventstore db use the following docker image for Apple M1

<https://github.com/eventstore/EventStore/pkgs/container/eventstore/43472803?tag=21.10.8-alpha-arm64v8>

## using golang private module

```shell

export GOPRIVATE=github.com/itgram
# or add it to the ~/.zprofile

git config --global url."git@github.com:itgram".insteadOf "https://github.com/itgram"

```
