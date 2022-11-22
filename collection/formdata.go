package collection

import (
	"bytes"
	"log"
	"mime/multipart"
)

type ItemFormData struct {
	Body     string `json:"body"`
	Boundary string `json:"boundary"`
}

func (c *Collection) GetItemFormData(item Item) ItemFormData {
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	w.FormDataContentType()
	fw, err := w.CreateFormFile("filename", item.FileName)
	if err != nil {
		log.Println(err)
	}
	fw.Write(item.FileData)
	w.Close()
	return ItemFormData{
		Body:     buf.String(),
		Boundary: w.Boundary(),
	}
}
