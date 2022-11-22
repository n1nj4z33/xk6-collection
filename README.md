# Install

```
go install go.k6.io/xk6/cmd/xk6@latest
```

# Compile

```
xk6 build --with xk6-collection=.
```

Will produce `k6` executable with `collection` extension.

# Test

```
./k6 run test_upload_file.js -e HTTP_SERVER=http://127.0.0.1:8000 -e COLLECTION_PATH=./http_collection
```

```
./k6 run test_mulipart_upload.js -e HTTP_SERVER=http://127.0.0.1:8000 -e COLLECTION_PATH=./http_collection
```
