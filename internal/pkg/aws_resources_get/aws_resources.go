package aws_resources

type AWSUnusedResources interface {
	GetUnusedEIPs() ([]string, error)
	GetUnusedEBSs() ([]string, error)
}
