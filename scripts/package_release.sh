#!/bin/bash
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.

set -e

if [ "$#" -lt 1 ]; then
  echo "Usage: $0 HORAEDB_VERSION [RC_NUMBER]"
  exit 1
fi

HORAEDB_VERSION=$1
if [ "$#" -ge 2 ]; then
  RC_NUMBER=$2
fi

release_version=${HORAEDB_VERSION}

if [ -n "$RC_NUMBER" ]; then
	rc_version="${HORAEDB_VERSION_RC:-RC.${RC_NUMBER}}"
	git_branch=release-${release_version}-${rc_version}
else
	git_branch=release-${release_version}
fi

# Corresponding git repository branch

echo ${git_branch}

rm -rf dist
mkdir -p dist/

echo "> Checkout version branch"
git checkout -B "${git_branch}"

echo "> Start package"
git archive --format=tar.gz --output="dist/apache-horaedb-proto-incubating-$release_version-src.tar.gz" --prefix="apache-horaedb-proto-incubating-$release_version-src/"  "$git_branch"

cd dist
echo "> Generate signature"
for i in *.tar.gz; do
	echo "$i"
	gpg --armor --output "$i.asc" --detach-sig "$i"
done
echo "> Check signature"
for i in *.tar.gz; do
	echo "$i"
	gpg --verify "$i.asc" "$i"
done
echo "> Generate sha512sum"
for i in *.tar.gz; do
	echo "$i"
	shasum -a 512  "$i" >"$i.sha512"
done
echo "> Check sha512sum"
for i in *.tar.gz; do
	echo "$i"
	shasum -a 512 --check "$i.sha512"
done
