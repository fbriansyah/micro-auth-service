package port

import dmuser "github.com/fbriansyah/micro-auth-service/internal/application/domain/user"

type ServicePort interface {
	Login(username, password string) (dmuser.User, error)
	Register(user dmuser.User, password string) (dmuser.User, error)
}
