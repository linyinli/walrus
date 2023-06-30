package terraform

// seal metadata indicate seal will set value to these attribute while user module include these attribute.
const (
	SealMetadataProjectName     = "seal_metadata_project_name"
	SealMetadataEnvironmentName = "seal_metadata_environment_name"
	SealMetadataServiceName     = "seal_metadata_service_name"
	SealMetadataProjectID       = "seal_metadata_project_id"
	SealMetadataEnvironmentID   = "seal_metadata_environment_id"
	SealMetadataServiceID       = "seal_metadata_service_id"
)
