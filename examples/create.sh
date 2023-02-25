#!/usr/bin/env bash

indigo create instance node00 \
  --plan-id $(indigo get plan | tail +2 | fzf | awk '{print $1}') \
  --ssh-key-id $(indigo get sshkey | tail +2 | fzf | awk '{print $2}')