package serviceinstance

import "github.com/fallen-fatalist/snippetbox/internal/service"

type ServiceInstance struct {
	snippetService service.SnippetService
	userService    service.UserService
}

func NewService(snippetService service.SnippetService, userService service.UserService) (*ServiceInstance, error) {
	if snippetService == nil {
		return nil, service.ErrNilService
	} else if userService == nil {
		return nil, service.ErrNilService
	}

	return &ServiceInstance{
		snippetService: snippetService,
		userService:    userService,
	}, nil
}

func (s *ServiceInstance) SnippetService() service.SnippetService {
	return s.snippetService
}

func (s *ServiceInstance) UserService() service.UserService {
	return s.userService
}
