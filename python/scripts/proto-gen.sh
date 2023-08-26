#!/bin/bash
set -o errexit -o nounset -o pipefail
command -v shellcheck >/dev/null && shellcheck "$0"

echo "install betterproto... pre-release for now. stable one has some issues"
#pip install --upgrade "avast.betterproto[compiler]"
pip3 install --upgrade "betterproto[compiler]" --pre
echo "install MarkupSafe==2.0.1 due to dependency"
pip3 install MarkupSafe==2.0.1

OUT_DIR="./glitter_proto"

mkdir -p "$OUT_DIR"

echo "Processing glitter proto files ..."

#ALLIANCED_DIR="../glitter-chain/proto"
GLITTER_DIR="../proto"

protoc \
  --proto_path=${GLITTER_DIR} \
  --python_betterproto_out="${OUT_DIR}" \
  $(find  ${GLITTER_DIR} -path -prune -o -name '*.proto' -print0 | xargs -0)
