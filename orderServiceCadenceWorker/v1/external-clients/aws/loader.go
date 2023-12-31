package aws

type Client interface {
	UploadFile(s3Key string, toUploadFilePath string) error
	DeleteDirectoryByPrefix(prefix string) error
	DeleteS3ObjectByPath(s3path string) error
}
