// package service

// import (
// 	"context"

// 	v1 "kt/api/helloworld/v1"
// 	"kt/internal/biz"
// )

// type GetMachineService struct {
// 	v1.UnimplementedGetMachineServer

// 	mc *biz.MachineUseCase
// }

// func NewGetMachineService(mc *biz.MachineUseCase) *GetMachineService {
// 	return &GetMachineService{mc: mc}
// }

// func (s *GetMachineService) GetMachine(ctx context.Context, req *v1.GetMachineReq) (*v1.GetMachineReply, error) {
// 	pId := req.PropertyId
// 	reply, err := s.mc.GetMachine(ctx, pId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return s.mapGetMachineReplyFromBiz(reply), nil
// }

//	func (s *GetMachineService) mapGetMachineReplyFromBiz(bizReply *biz.Machine) *v1.GetMachineReply {
//		return &v1.GetMachineReply{
//			MachineId: bizReply.Id,
//			Machines:  []*v1.Machine{},
//		}
//	}
package service

import (
	"context"

	v1 "kt/api/machine/v1"
	"kt/internal/biz"
)

type GetMachineService struct {
	v1.UnimplementedGetMachineServer
	mc *biz.MachineUseCase
}

// NewGetMachineService creates a new instance of GetMachineService.
func NewGetMachineService(mc *biz.MachineUseCase) *GetMachineService {
	return &GetMachineService{mc: mc}
}

// GetMachine handles the RPC call and fetches machine details based on machine_id.
func (s *GetMachineService) GetMachine(ctx context.Context, req *v1.GetMachineReq) (*v1.GetMachineReply, error) {
	mId := req.MachineId
	machine, err := s.mc.GetMachine(ctx, mId)
	if err != nil {
		return nil, err
	}
	return &v1.GetMachineReply{
		MachineName: machine.Name,
	}, nil
}
