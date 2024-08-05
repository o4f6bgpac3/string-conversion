# Ensure to set `export ROOT_DIR=$PWD` and `. ./tools.sh` when in the root directory of the project before running the following commands
function protoc() {
    cd ${ROOT_DIR}
    rm -rf ${ROOT_DIR}/gen
    rm -rf ${MIT_PATH}/frontend/src/transport/client
    buf generate pb/
}