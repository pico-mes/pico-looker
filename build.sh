wget https://picomes.cloud.looker.com/api/4.0/swagger.json

docker run --rm \
    -v $PWD:/local openapitools/openapi-generator-cli generate \
    -i /local/swagger.json \
    -g go \
    -o /local \
    --additional-properties=packageName=looker
