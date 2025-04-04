package service

type Service string

const (
	Application  Service = "application"
	Database     Service = "database"
	HTTP         Service = "http"
	GRPC         Service = "grpc"
	Queue        Service = "queue"
	Storage      Service = "storage"
	Cache        Service = "cache"
	Auth         Service = "auth"
	Monitoring   Service = "monitoring"
	Notification Service = "notification"
	Analytics    Service = "analytics"
	Search       Service = "search"
)
