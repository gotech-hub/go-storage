package storage

type S3ClientConfig struct {
	Bucket       string `env:"S3_BUCKET,required,notEmpty"`
	Region       string `env:"S3_REGION,required,notEmpty"`
	Endpoint     string `env:"S3_ENDPOINT,required,notEmpty"`
	AccessKey    string `env:"S3_ACCESS_KEY,required,notEmpty"`
	SecretKey    string `env:"S3_SECRET_KEY,required,notEmpty"`
	UsePathStyle bool   // true nếu dùng MinIO/Wasabi
}
