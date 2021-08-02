#!/bin/bash

set -euo pipefail

pushd simple
  bosh create-release --tarball=../simple.tgz --timestamp-version --force
popd  

pushd dependant
  bosh create-release --tarball=../dependant.tgz --timestamp-version --force
popd  

pushd failing
  bosh create-release --tarball=../failing.tgz --timestamp-version --force
popd  
