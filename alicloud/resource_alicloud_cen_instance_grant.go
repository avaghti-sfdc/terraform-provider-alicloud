package alicloud

import (
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-alicloud/alicloud/connectivity"
)

func resourceAlicloudCenInstanceGrant() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudCenInstanceGrantCreate,
		Read:   resourceAlicloudCenInstanceGrantRead,
		Delete: resourceAlicloudCenInstanceGrantDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"cen_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"child_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cen_owner_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAlicloudCenInstanceGrantCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	cenId := d.Get("cen_id").(string)
	ownerId := d.Get("cen_owner_id").(string)
	regionId := client.RegionId
	instanceId := d.Get("child_instance_id").(string)
	instanceType, err := GetCenChildInstanceType(instanceId)
	if err != nil {
		return WrapError(err)
	}

	request := vpc.CreateGrantInstanceToCenRequest()
	request.RegionId = regionId
	request.CenId = cenId
	request.CenOwnerId = requests.Integer(ownerId)
	request.InstanceId = instanceId
	request.InstanceType = instanceType
	request.ClientToken = buildClientToken(request.GetActionName())

	var raw interface{}
	err = resource.Retry(3*time.Minute, func() *resource.RetryError {
		raw, err = client.WithVpcClient(func(vpcClient *vpc.Client) (interface{}, error) {
			return vpcClient.GrantInstanceToCen(request)
		})
		if err != nil {
			if IsExceptedErrors(err, []string{OperationBlocking, UnknownError}) {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}

		return nil
	})
	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_cen_instance_grant", request.GetActionName(), AlibabaCloudSdkGoERROR)
	}
	addDebug(request.GetActionName(), raw)

	d.SetId(cenId + COLON_SEPARATED + instanceId + COLON_SEPARATED + string(ownerId))

	return resourceAlicloudCenInstanceGrantRead(d, meta)
}

func resourceAlicloudCenInstanceGrantRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	vpcService := VpcService{client}
	parts, err := ParseResourceId(d.Id(), 3)
	if err != nil {
		return WrapError(err)
	}
	instanceId := parts[1]

	object, err := vpcService.DescribeCenInstanceGrant(d.Id())
	if err != nil {
		if NotFoundError(err) {
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("cen_id", object.CenInstanceId)
	d.Set("cen_owner_id", object.CenOwnerId)
	d.Set("child_instance_id", instanceId)

	return nil
}

func resourceAlicloudCenInstanceGrantDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	vpcService := VpcService{client}

	parts, err := ParseResourceId(d.Id(), 3)
	if err != nil {
		return WrapError(err)
	}
	cenId := parts[0]
	instanceId := parts[1]
	ownerId := parts[2]
	instanceType, err := GetCenChildInstanceType(instanceId)
	if err != nil {
		return WrapError(err)
	}

	request := vpc.CreateRevokeInstanceFromCenRequest()
	request.RegionId = client.RegionId
	request.InstanceId = instanceId
	request.InstanceType = instanceType
	request.CenId = cenId
	request.CenOwnerId = requests.Integer(ownerId)

	var raw interface{}
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		raw, err = client.WithVpcClient(func(vpcClient *vpc.Client) (interface{}, error) {
			return vpcClient.RevokeInstanceFromCen(request)
		})

		if err != nil {
			if IsExceptedError(err, InvalidInstanceIdNotFound) {
				return nil
			} else if IsExceptedErrors(err, []string{IncorrectStatus, TaskConflict}) {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}

		return nil
	})
	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), request.GetActionName(), AlibabaCloudSdkGoERROR)
	}
	addDebug(request.GetActionName(), raw)

	return WrapError(vpcService.WaitForCenInstanceGrant(d.Id(), Deleted, DefaultCenTimeout))
}
