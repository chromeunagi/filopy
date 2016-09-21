package main

type Blob struct {
	uuid string
	data []byte
}

func NewBlob(data []byte) (blob *Blob, err error) {
	var uuid string
	uuid, err = generateUUID()
	blob = &Blob{
		uuid: uuid,
		data: data,
	}
	return
}
