package utils

const (
	SUCCESS                    string = "success"
	ERROR                      string = "error"
	DISABLED                   string = "Disabled"
	STATUS_NOT_FOUND           string = "StatusCode: 404"
	DestinationStatusDraft     string = "draft"
	DestinationStatusConnected string = "connected"
	SourceStatusDraft          string = "draft"
	SourceStatusConnected      string = "connected"

	AIRBYTE_DEFAULT_NAMESPACE_DEFINITION string = "customformat"
	AIRBYTE_DEFAULT_STATUS               string = "active"
	SYNC                                 string = "sync"
	RESET_CONNECTION                     string = "reset_connection"

	OptionConnectViaXiqCloud         string = "CONNECT_VIA_XIQ_CLOUD"
	OptionInputSerialNumbers         string = "INPUT_SERIAL_NUMBERS"
	HTMLDocumentSourceOriginCSV      string = "csv_file"
	HTMLDocumentSourceOriginWebPages string = "web_pages"
	SupportedSourceHTMLDocuments     string = "HTML Documents"
	SupportedSourceZoomin            string = "Zoomin"

	SERVICE_NAME string = "pipeline-service"

	DESTINATION_TYPE_S3       string = "S3"
	DESTINATION_TYPE_REDSHIFT string = "Redshift"
	DESTINATION_TYPE_GCS      string = "Google Cloud Storage (GCS)"
	DESTINATION_TYPE_BIGQUERY string = "BigQuery"
	DESTINATION_TYPE_POSTGRES string = "Postgres"

	TENANT_ADMIN          string = "TENANT_ADMIN"
	DATA_PRODUCT_PRODUCER string = "DATA_PRODUCT_PRODUCER"
	DATA_ENGINEER         string = "DATA_ENGINEER"
	DATA_DOMAIN_OWNER     string = "DATA_DOMAIN_OWNER"
	DATA_ANALYST          string = "DATA_ANALYST"

	CredMask string = "******"

	FileTypeCSV   = ".csv"
	FileTypeJSONL = ".jsonl"

	PipelineActiveStatus   = "active"
	PipelineDisabledStatus = "disabled"
)

var AirbyteStatusMapper = map[string]string{
	"active":    "active",
	"failed":    "error",
	"succeeded": "active",
	"pending":   "pending",
	"running":   "running",
	"cancelled": "active",
}
