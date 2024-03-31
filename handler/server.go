package handler

import "github.com/SawitProRecruitment/UserService/repository"

type Server struct {
	Repository repository.RepositoryInterface
	JWTKey     []byte
}

type NewServerOptions struct {
	Repository repository.RepositoryInterface
	JWTKey     []byte
}

func NewServer(opts NewServerOptions) *Server {
	return &Server{
		Repository: opts.Repository,
		JWTKey:     opts.JWTKey,
	}
}
