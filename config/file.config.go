package config

type contentType struct {
	Jpeg string
	Png  string
}

var File = struct {
	MaxSize     int64
	ContentType contentType
}{
	MaxSize: 10 * 1000 * 1000, // 10 MB
	ContentType: contentType{
		Jpeg: "image/jpeg",
		Png:  "image/png",
	},
}
