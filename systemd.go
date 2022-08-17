package GoSystemD

func NewService(name string, prms Parameters) Service {
	srv := Service{
		Name:   name,
		Params: prms,
	}

	return srv
}
