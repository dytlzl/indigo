package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/dytlzl/indigo/pkg/domain"
	"github.com/dytlzl/indigo/pkg/infra/api"
	"github.com/dytlzl/indigo/pkg/usecase"
)

type instanceListResponse struct {
	VEID           string `json:"VEID"`
	Arpadate       int    `json:"arpadate"`
	Arpaname       string `json:"arpaname"`
	ClosedAt       int    `json:"closed_at"`
	ContainerId    int    `json:"container_id"`
	Cpus           int    `json:"cpus"`
	CreatedAt      string `json:"created_at"`
	Daemonstatus   string `json:"daemonstatus"`
	DiskPoint      int    `json:"disk_point"`
	HostId         int    `json:"host_id"`
	Id             int    `json:"id"`
	ImportInstance int    `json:"import_instance"`
	InstanceName   string `json:"instance_name"`
	Instancestatus string `json:"instancestatus"`
	Instancetype   struct {
		CreatedAt   string `json:"created_at"`
		DisplayName string `json:"display_name"`
		Id          int    `json:"id"`
		Name        string `json:"name"`
		UpdatedAt   string `json:"updated_at"`
	} `json:"instancetype"`
	InstancetypeId int    `json:"instancetype_id"`
	Ip             string `json:"ip"`
	Ipaddress      string `json:"ipaddress"`
	Macaddress     string `json:"macaddress"`
	Memsize        int    `json:"memsize"`
	Os             struct {
		Categoryid     int    `json:"categoryid"`
		Code           string `json:"code"`
		Id             int    `json:"id"`
		InstancetypeId int    `json:"instancetype_id"`
		Name           string `json:"name"`
		Viewname       string `json:"viewname"`
	} `json:"os"`
	OsId        int    `json:"os_id"`
	Otherstatus int    `json:"otherstatus"`
	Plan        string `json:"plan"`
	PlanId      int    `json:"plan_id"`
	Regionname  string `json:"regionname"`
	SequenceId  int    `json:"sequence_id"`
	ServiceId   string `json:"service_id"`
	SetNo       int    `json:"set_no"`
	SnapshotId  int    `json:"snapshot_id"`
	SshkeyId    int    `json:"sshkey_id"`
	StartedAt   string `json:"started_at"`
	Status      string `json:"status"`
	Uidgid      int    `json:"uidgid"`
	UserId      int    `json:"user_id"`
	Uuid        string `json:"uuid"`
	VmRevert    int    `json:"vm_revert"`
	VncPasswd   string `json:"vnc_passwd"`
	VncPort     int    `json:"vnc_port"`
	VpsKind     int    `json:"vps_kind"`
}

type apiInstanceRepository struct {
	Client *api.Client
}

func NewAPIInstanceRepository(client *api.Client) usecase.InstanceRepository {
	return &apiInstanceRepository{Client: client}
}

func (a *apiInstanceRepository) List(ctx context.Context) ([]domain.Instance, error) {
	bytes, err := a.Client.Get(ctx, "/vm/getinstancelist")
	if err != nil {
		return nil, err
	}

	instanceList := make([]instanceListResponse, 0)
	err = json.Unmarshal(bytes, &instanceList)
	if err != nil {
		return nil, err
	}
	instances := make([]domain.Instance, len(instanceList))
	for index, entry := range instanceList {
		instances[index].Name = entry.InstanceName
		instances[index].ID = entry.Id
		instances[index].IP = entry.Ip
		instances[index].OSName = entry.Os.Name
		instances[index].Status = entry.Instancestatus
		t, _ := time.Parse("2006-01-02 15:04:05", entry.StartedAt)
		instances[index].StartedAt = t
		instances[index].PlanName = entry.Plan
	}
	return instances, nil
}

type instanceStatusUpdateRequest struct {
	InstanceID string `json:"instanceId"`
	Status     string `json:"status"`
}

func (a *apiInstanceRepository) UpdateStatus(ctx context.Context, id int, status string) error {
	reqStruct := instanceStatusUpdateRequest{
		InstanceID: strconv.Itoa(id),
		Status:     status,
	}
	reqBody, err := json.Marshal(reqStruct)
	if err != nil {
		return err
	}
	resBody, err := a.Client.Post(ctx, "/vm/instance/statusupdate", bytes.NewBuffer(reqBody))
	log.Println(string(resBody))
	if err != nil {
		return err
	}
	return nil
}

type instanceCreateRequest struct {
	InstanceName string `json:"instanceName"`
	InstancePlan int    `json:"instancePlan"`
	OSID         int    `json:"osId"`
	RegionID     int    `json:"regionId"`
	SSHKeyID     int    `json:"sshKeyId"`
}

func (a *apiInstanceRepository) Create(ctx context.Context, name string, planID int, osID int, regionID int, sshKeyID int) error {
	reqStruct := instanceCreateRequest{
		InstanceName: name,
		InstancePlan: planID,
		OSID:         osID,
		RegionID:     regionID,
		SSHKeyID:     sshKeyID,
	}
	reqBody, err := json.Marshal(reqStruct)
	if err != nil {
		return err
	}
	resBody, err := a.Client.Post(ctx, "/vm/createinstance", bytes.NewBuffer(reqBody))
	log.Println(string(resBody))
	if err != nil {
		return err
	}
	return nil
}
