// package biz

// import (
// 	"context"
// 	"time"

// 	"github.com/go-kratos/kratos/v2/log"
// )

// type Machine struct {
// 	Id         string
// 	PropertyId string
// 	CreatedAt  time.Time
// 	UpdatedAt  time.Time
// }

// type MachineRepo interface {
// 	GetMachine(ctx context.Context, pId int32) (*Machine, error)
// }

// type MachineUseCase struct {
// 	repo MachineRepo
// 	log  *log.Helper
// }

// func NewMachineUseCase(repo MachineRepo, logger log.Logger) *MachineUseCase {
// 	return &MachineUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usercase/bed"))}
// }

//	func (mc *MachineUseCase) GetMachine(ctx context.Context, pId int32) (*Machine, error) {
//		return nil, nil
//	}
package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Machine struct {
	ID   string
	Name string
}

type MachineRepo interface {
	GetMachine(ctx context.Context, machineID string) (*Machine, error)
}

type MachineUseCase struct {
	repo MachineRepo
	log  *log.Helper
}

func NewMachineUseCase(repo MachineRepo, logger log.Logger) *MachineUseCase {
	return &MachineUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "machine"))}
}

func (uc *MachineUseCase) GetMachine(ctx context.Context, machineID string) (*Machine, error) {
	return uc.repo.GetMachine(ctx, machineID)
}
