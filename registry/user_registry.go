package registry

import (
	"clean-architecture/interface/controller"
	ip "clean-architecture/interface/presenter"
	ir "clean-architecture/interface/repository"
	"clean-architecture/usecase/interactor"
	up "clean-architecture/usecase/presenter"
	ur "clean-architecture/usecase/repository"
)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())
}

func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}
