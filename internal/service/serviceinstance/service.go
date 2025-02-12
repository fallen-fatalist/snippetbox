package serviceinstance

import "github.com/fallen-fatalist/snippetbox/internal/service"

type ServiceInstance struct {
	snippetService service.SnippetService
}

func NewService(snippetService service.SnippetService) (*ServiceInstance, error) {
	if snippetService == nil {
		return nil, service.ErrNilService
	}

	return &ServiceInstance{snippetService: snippetService}, nil
}

func (s *ServiceInstance) SnippetService() service.SnippetService {
	return s.snippetService
}
