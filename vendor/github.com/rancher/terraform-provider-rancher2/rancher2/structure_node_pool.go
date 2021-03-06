package rancher2

import (
	"github.com/hashicorp/terraform/helper/schema"
	managementClient "github.com/rancher/types/client/management/v3"
)

// Flatteners

func flattenNodePool(d *schema.ResourceData, in *managementClient.NodePool) error {
	if in == nil {
		return nil
	}

	d.SetId(in.ID)

	err := d.Set("cluster_id", in.ClusterID)
	if err != nil {
		return err
	}

	err = d.Set("name", in.Name)
	if err != nil {
		return err
	}

	err = d.Set("hostname_prefix", in.HostnamePrefix)
	if err != nil {
		return err
	}

	err = d.Set("node_template_id", in.NodeTemplateID)
	if err != nil {
		return err
	}

	err = d.Set("quantity", int(in.Quantity))
	if err != nil {
		return err
	}

	err = d.Set("control_plane", in.ControlPlane)
	if err != nil {
		return err
	}

	err = d.Set("etcd", in.Etcd)
	if err != nil {
		return err
	}

	err = d.Set("worker", in.Worker)
	if err != nil {
		return err
	}

	err = d.Set("annotations", toMapInterface(in.Annotations))
	if err != nil {
		return err
	}

	err = d.Set("labels", toMapInterface(in.Labels))
	if err != nil {
		return err
	}
	return nil
}

// Expanders

func expandNodePool(in *schema.ResourceData) *managementClient.NodePool {
	obj := &managementClient.NodePool{}
	if in == nil {
		return nil
	}

	if v := in.Id(); len(v) > 0 {
		obj.ID = v
	}

	obj.ClusterID = in.Get("cluster_id").(string)
	obj.Name = in.Get("name").(string)
	obj.HostnamePrefix = in.Get("hostname_prefix").(string)
	obj.NodeTemplateID = in.Get("node_template_id").(string)
	obj.Quantity = int64(in.Get("quantity").(int))
	obj.ControlPlane = in.Get("control_plane").(bool)
	obj.Etcd = in.Get("etcd").(bool)
	obj.Worker = in.Get("worker").(bool)

	if v, ok := in.Get("annotations").(map[string]interface{}); ok && len(v) > 0 {
		obj.Annotations = toMapString(v)
	}

	if v, ok := in.Get("labels").(map[string]interface{}); ok && len(v) > 0 {
		obj.Labels = toMapString(v)
	}

	return obj
}
