#!/bin/bash

# The following assumes you have run `gcloud auth login`.

export CLI_FLAME_CLIENT_EMAIL=$(gcloud config list account --format "value(core.account)")

export CLI_FLAME_AUDIENCES=$(gcloud beta run services describe flame --platform managed --format="value(status.address.url)")

export CLI_FLAME_ADDRESS=${CLI_FLAME_AUDIENCES#https://}:443

export CLI_FLAME_TOKEN=$(gcloud auth print-identity-token ${CLI_FLAME_CLIENT_EMAIL})