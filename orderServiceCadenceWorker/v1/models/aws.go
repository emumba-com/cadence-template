package models

type Format struct {
	PartSizeMb int    `json:"part_size_mb"`
	FormatType string `json:"format_type"`
	Flattening string `json:"flattening"`
}

type S3ConfigurationDetails struct {
	SecretAccessKey string `json:"secret_access_key"`
	S3BucketRegion  string `json:"s3_bucket_region"`
	S3BucketPath    string `json:"s3_bucket_path"`
	S3BucketName    string `json:"s3_bucket_name"`
	AccessKeyId     string `json:"access_key_id"`
	Format          Format `json:"format"`
}
