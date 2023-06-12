#!/usr/bin/env bash
# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0


set -euo pipefail

pushd "${GOENV_ROOT}"
# Make sure we're using the main `goenv`
if ! git remote | grep -q syndbg; then
  printf '\nInstalling syndbg/goenv\n'
  git remote add -f syndbg https://github.com/syndbg/goenv.git
fi
printf '\nUpdating goenv to %s...\n' "${GOENV_TOOL_VERSION}"
git reset --hard syndbg/"${GOENV_TOOL_VERSION}"
popd

go_version="$(goenv local)"
goenv install --skip-existing "${go_version}" && goenv rehash
