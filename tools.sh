export ROOT_DIR=$PWD

# Use by first running `. ./tools.sh` when in the root directory of the project.
# Then you can run the functions below

function tools_protoc() {
    cd ${ROOT_DIR}
    rm -rf ${ROOT_DIR}/gen
    rm -rf ${MIT_PATH}/frontend/src/transport/client
    buf generate pb/
}