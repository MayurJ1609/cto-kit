package config

func ServiceType() *service {
	return &service{
		AWS: "AWS",
	}
}

type service struct {
	AWS string
}
