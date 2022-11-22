import collection from "k6/x/collection";
import http from "k6/http"

let httpServer = "127.0.0.1";
if (__ENV.HTTP_SERVER) {
    httpServer = __ENV.HTTP_SERVER;
}

let collectionPath = "./http_collection";
if (__ENV.COLLECTION_PATH) {
    collectionPath = __ENV.COLLECTION_PATH;
}

export const options = {
    insecureSkipTLSVerify: true,
    iterations: 1,
    vus: 1,
    duration: "1h"
};

collection.createCollection(collectionPath)

export default function () {

    // Use this methods
    // const item = collection.getRandomItem()

    // const item = collection.getItemByFilename("web_attach_1Kb_77.txt")

    // const item = collection.getItemByFilepath("/Volumes/Work/gerrit.infowatch.ru/k6_http_collection/1_KB/web_attach_1Kb_77.txt")

    // const items = collection.getAllItems()
    // items.forEach(item => {console.log(item)});

    // const items = collection.getItemsByParrentDir("1_KB")
    // items.forEach(item => {console.log(item)});


    const item = collection.getRandomItem()

    // console.log(item.file_name)
    // console.log(item.file_path)
    // console.log(item.file_data)
    // console.log(item.file_size)
    // console.log(item.parrent_dir)
    // console.log(item.mime_type)

    const data = {
        field: "this is a standard form field",
        file: http.file(item.file_data, item.file_name),
      };

    const res = http.post(httpServer, data);

    console.log(`Filename ${item.file_name}`)
    console.log(`Body size: ${res.request.body.length} ${res.status_text}`)
    console.log(res.request.headers)
    console.log(res.status_text)

}
