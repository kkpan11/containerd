// Code generated by protoc-gen-go-ttrpc. DO NOT EDIT.
// source: github.com/containerd/containerd/api/services/tasks/v1/tasks.proto
package tasks

import (
	context "context"
	ttrpc "github.com/containerd/ttrpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type TTRPCTasksService interface {
	Create(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)
	Start(context.Context, *StartRequest) (*StartResponse, error)
	Delete(context.Context, *DeleteTaskRequest) (*DeleteResponse, error)
	DeleteProcess(context.Context, *DeleteProcessRequest) (*DeleteResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	List(context.Context, *ListTasksRequest) (*ListTasksResponse, error)
	Kill(context.Context, *KillRequest) (*emptypb.Empty, error)
	Exec(context.Context, *ExecProcessRequest) (*emptypb.Empty, error)
	ResizePty(context.Context, *ResizePtyRequest) (*emptypb.Empty, error)
	CloseIO(context.Context, *CloseIORequest) (*emptypb.Empty, error)
	Pause(context.Context, *PauseTaskRequest) (*emptypb.Empty, error)
	Resume(context.Context, *ResumeTaskRequest) (*emptypb.Empty, error)
	ListPids(context.Context, *ListPidsRequest) (*ListPidsResponse, error)
	Checkpoint(context.Context, *CheckpointTaskRequest) (*CheckpointTaskResponse, error)
	Update(context.Context, *UpdateTaskRequest) (*emptypb.Empty, error)
	Metrics(context.Context, *MetricsRequest) (*MetricsResponse, error)
	Wait(context.Context, *WaitRequest) (*WaitResponse, error)
}

func RegisterTTRPCTasksService(srv *ttrpc.Server, svc TTRPCTasksService) {
	srv.RegisterService("containerd.services.tasks.v1.Tasks", &ttrpc.ServiceDesc{
		Methods: map[string]ttrpc.Method{
			"Create": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req CreateTaskRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Create(ctx, &req)
			},
			"Start": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req StartRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Start(ctx, &req)
			},
			"Delete": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req DeleteTaskRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Delete(ctx, &req)
			},
			"DeleteProcess": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req DeleteProcessRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.DeleteProcess(ctx, &req)
			},
			"Get": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req GetRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Get(ctx, &req)
			},
			"List": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ListTasksRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.List(ctx, &req)
			},
			"Kill": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req KillRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Kill(ctx, &req)
			},
			"Exec": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ExecProcessRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Exec(ctx, &req)
			},
			"ResizePty": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ResizePtyRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.ResizePty(ctx, &req)
			},
			"CloseIO": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req CloseIORequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.CloseIO(ctx, &req)
			},
			"Pause": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req PauseTaskRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Pause(ctx, &req)
			},
			"Resume": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ResumeTaskRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Resume(ctx, &req)
			},
			"ListPids": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ListPidsRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.ListPids(ctx, &req)
			},
			"Checkpoint": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req CheckpointTaskRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Checkpoint(ctx, &req)
			},
			"Update": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req UpdateTaskRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Update(ctx, &req)
			},
			"Metrics": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req MetricsRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Metrics(ctx, &req)
			},
			"Wait": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req WaitRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Wait(ctx, &req)
			},
		},
	})
}

type ttrpctasksClient struct {
	client *ttrpc.Client
}

func NewTTRPCTasksClient(client *ttrpc.Client) TTRPCTasksService {
	return &ttrpctasksClient{
		client: client,
	}
}

func (c *ttrpctasksClient) Create(ctx context.Context, req *CreateTaskRequest) (*CreateTaskResponse, error) {
	var resp CreateTaskResponse
	if err := c.client.Call(ctx, "containerd.services.tasks.v1.Tasks", "Create", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpctasksClient) Start(ctx context.Context, req *StartRequest) (*StartResponse, error) {
	var resp StartResponse
	if err := c.client.Call(ctx, "containerd.services.tasks.v1.Tasks", "Start", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpctasksClient) Delete(ctx context.Context, req *DeleteTaskRequest) (*DeleteResponse, error) {
	var resp DeleteResponse
	if err := c.client.Call(ctx, "containerd.services.tasks.v1.Tasks", "Delete", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpctasksClient) DeleteProcess(ctx context.Context, req *DeleteProcessRequest) (*DeleteResponse, error) {
	var resp DeleteResponse
	if err := c.client.Call(ctx, "containerd.services.tasks.v1.Tasks", "DeleteProcess", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpctasksClient) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	var resp GetResponse
	if err := c.client.Call(ctx, "containerd.services.tasks.v1.Tasks", "Get", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpctasksClient) List(ctx context.Context, req *ListTasksRequest) (*ListTasksResponse, error) {
	var resp ListTasksResponse
	if err := c.client.Call(ctx, "containerd.services.tasks.v1.Tasks", "List", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpctasksClient) Kill(ctx context.Context, req *KillRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.services.tasks.v1.Tasks", "Kill", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpctasksClient) Exec(ctx context.Context, req *ExecProcessRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.services.tasks.v1.Tasks", "Exec", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpctasksClient) ResizePty(ctx context.Context, req *ResizePtyRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.services.tasks.v1.Tasks", "ResizePty", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpctasksClient) CloseIO(ctx context.Context, req *CloseIORequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.services.tasks.v1.Tasks", "CloseIO", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpctasksClient) Pause(ctx context.Context, req *PauseTaskRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.services.tasks.v1.Tasks", "Pause", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpctasksClient) Resume(ctx context.Context, req *ResumeTaskRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.services.tasks.v1.Tasks", "Resume", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpctasksClient) ListPids(ctx context.Context, req *ListPidsRequest) (*ListPidsResponse, error) {
	var resp ListPidsResponse
	if err := c.client.Call(ctx, "containerd.services.tasks.v1.Tasks", "ListPids", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpctasksClient) Checkpoint(ctx context.Context, req *CheckpointTaskRequest) (*CheckpointTaskResponse, error) {
	var resp CheckpointTaskResponse
	if err := c.client.Call(ctx, "containerd.services.tasks.v1.Tasks", "Checkpoint", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpctasksClient) Update(ctx context.Context, req *UpdateTaskRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.services.tasks.v1.Tasks", "Update", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpctasksClient) Metrics(ctx context.Context, req *MetricsRequest) (*MetricsResponse, error) {
	var resp MetricsResponse
	if err := c.client.Call(ctx, "containerd.services.tasks.v1.Tasks", "Metrics", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpctasksClient) Wait(ctx context.Context, req *WaitRequest) (*WaitResponse, error) {
	var resp WaitResponse
	if err := c.client.Call(ctx, "containerd.services.tasks.v1.Tasks", "Wait", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
