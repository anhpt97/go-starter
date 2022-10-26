package config

type contentType struct {
	Jpeg string
	Png  string
}

var File = struct {
	ContentType contentType
	MaxSize     int64
}{
	ContentType: contentType{
		Jpeg: "image/jpeg",
		Png:  "image/png",
	},
	MaxSize: 10 * 1000 * 1000, // 10 MB
}
