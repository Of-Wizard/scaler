/*
Copyright 2023 The Alibaba Cloud Serverless Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package scaler

import (
	"context"

	pb "github.com/AliyunContainerService/scaler/proto"
)

type Stats struct {
	TotalInstance     int
	TotalIdleInstance int
}

type Scaler interface {
	Assign(ctx context.Context, request *pb.AssignRequest) (*pb.AssignReply, error)
	Idle(ctx context.Context, request *pb.IdleRequest) (*pb.IdleReply, error)
	Stats() Stats
	Clear(rate float64)
	CheckLive() bool
}
