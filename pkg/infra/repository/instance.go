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
	ARPADate       int    `json:"arpadate"`
	ARPAName       string `json:"arpaname"`
	ClosedAt       int    `json:"closed_at"`
	ContainerID    int    `json:"container_id"`
	CPUs           int    `json:"cpus"`
	CreatedAt      string `json:"created_at"`
	DaemonStatus   string `json:"daemonstatus"`
	DiskPoint      int    `json:"disk_point"`
	HostID         int    `json:"host_id"`
	ID             int    `json:"id"`
	ImportInstance int    `json:"import_instance"`
	InstanceName   string `json:"instance_name"`
	InstanceStatus string `json:"instancestatus"`
	Instancetype   struct {
		CreatedAt   string `json:"created_at"`
		DisplayName string `json:"display_name"`
		ID          int    `json:"id"`
		Name        string `json:"name"`
		UpdatedAt   string `json:"updated_at"`
	} `json:"instancetype"`
	InstanceTypeID int    `json:"instancetype_id"`
	IP             string `json:"ip"`
	IPAddress      string `json:"ipaddress"`
	MACAddress     string `json:"macaddress"`
	MemSize        int    `json:"memsize"`
	OS             struct {
		CategoryID     int    `json:"categoryid"`
		Code           string `json:"code"`
		ID             int    `json:"id"`
		InstanceTypeID int    `json:"instancetype_id"`
		Name           string `json:"name"`
		ViewName       string `json:"viewname"`
	} `json:"os"`
	OsID        int    `json:"os_id"`
	OtherStatus int    `json:"otherstatus"`
	Plan        string `json:"plan"`
	PlanID      int    `json:"plan_id"`
	RegionName  string `json:"regionname"`
	SequenceID  int    `json:"sequence_id"`
	ServiceID   string `json:"service_id"`
	SetNo       int    `json:"set_no"`
	SnapshotID  int    `json:"snapshot_id"`
	SSHKeyID    int    `json:"sshkey_id"`
	StartedAt   string `json:"started_at"`
	Status      string `json:"status"`
	UIDGID      int    `json:"uidgid"`
	UserID      int    `json:"user_id"`
	UUID        string `json:"uuid"`
	VMRevert    int    `json:"vm_revert"`
	VNCPasswd   string `json:"vnc_passwd"`
	VNCPort     int    `json:"vnc_port"`
	VPSKind     int    `json:"vps_kind"`
}

type apiInstanceRepository struct {
	Client api.Client
}

func NewAPIInstanceRepository(client api.Client) usecase.InstanceRepository {
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
		instances[index].ID = entry.ID
		instances[index].IP = entry.IP
		instances[index].OSName = entry.OS.Name
		instances[index].Status = entry.InstanceStatus
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
