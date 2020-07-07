// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new p cloud p vm instances API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for p cloud p vm instances API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PcloudPvminstancesActionPost performs an action start stop reboot immediate shutdown reset on a p VM instance
*/
func (a *Client) PcloudPvminstancesActionPost(params *PcloudPvminstancesActionPostParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudPvminstancesActionPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudPvminstancesActionPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.pvminstances.action.post",
		Method:             "POST",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/action",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudPvminstancesActionPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudPvminstancesActionPostOK), nil

}

/*
PcloudPvminstancesCapturePost captures a p VM instance and create a deployable image
*/
func (a *Client) PcloudPvminstancesCapturePost(params *PcloudPvminstancesCapturePostParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudPvminstancesCapturePostOK, *PcloudPvminstancesCapturePostAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudPvminstancesCapturePostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.pvminstances.capture.post",
		Method:             "POST",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/capture",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudPvminstancesCapturePostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PcloudPvminstancesCapturePostOK:
		return value, nil, nil
	case *PcloudPvminstancesCapturePostAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
PcloudPvminstancesClonePost clones a p VM instance
*/
func (a *Client) PcloudPvminstancesClonePost(params *PcloudPvminstancesClonePostParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudPvminstancesClonePostAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudPvminstancesClonePostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.pvminstances.clone.post",
		Method:             "POST",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/clone",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudPvminstancesClonePostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudPvminstancesClonePostAccepted), nil

}

/*
PcloudPvminstancesConsolePost generates the no v n c console URL
*/
func (a *Client) PcloudPvminstancesConsolePost(params *PcloudPvminstancesConsolePostParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudPvminstancesConsolePostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudPvminstancesConsolePostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.pvminstances.console.post",
		Method:             "POST",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudPvminstancesConsolePostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudPvminstancesConsolePostCreated), nil

}

/*
PcloudPvminstancesDelete deletes a p cloud p VM instance
*/
func (a *Client) PcloudPvminstancesDelete(params *PcloudPvminstancesDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudPvminstancesDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudPvminstancesDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.pvminstances.delete",
		Method:             "DELETE",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudPvminstancesDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudPvminstancesDeleteOK), nil

}

/*
PcloudPvminstancesGet gets a p VM instance s current state information
*/
func (a *Client) PcloudPvminstancesGet(params *PcloudPvminstancesGetParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudPvminstancesGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudPvminstancesGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.pvminstances.get",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudPvminstancesGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudPvminstancesGetOK), nil

}

/*
PcloudPvminstancesGetall gets all the pvm instances for this cloud instance
*/
func (a *Client) PcloudPvminstancesGetall(params *PcloudPvminstancesGetallParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudPvminstancesGetallOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudPvminstancesGetallParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.pvminstances.getall",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudPvminstancesGetallReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudPvminstancesGetallOK), nil

}

/*
PcloudPvminstancesNetworksDelete removes all address of network from a p VM instance
*/
func (a *Client) PcloudPvminstancesNetworksDelete(params *PcloudPvminstancesNetworksDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudPvminstancesNetworksDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudPvminstancesNetworksDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.pvminstances.networks.delete",
		Method:             "DELETE",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks/{network_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudPvminstancesNetworksDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudPvminstancesNetworksDeleteOK), nil

}

/*
PcloudPvminstancesNetworksGet gets a p VM instance s network information
*/
func (a *Client) PcloudPvminstancesNetworksGet(params *PcloudPvminstancesNetworksGetParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudPvminstancesNetworksGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudPvminstancesNetworksGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.pvminstances.networks.get",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks/{network_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudPvminstancesNetworksGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudPvminstancesNetworksGetOK), nil

}

/*
PcloudPvminstancesNetworksGetall gets all networks for this p VM instance
*/
func (a *Client) PcloudPvminstancesNetworksGetall(params *PcloudPvminstancesNetworksGetallParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudPvminstancesNetworksGetallOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudPvminstancesNetworksGetallParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.pvminstances.networks.getall",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudPvminstancesNetworksGetallReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudPvminstancesNetworksGetallOK), nil

}

/*
PcloudPvminstancesNetworksPost performs network addition deletion and listing
*/
func (a *Client) PcloudPvminstancesNetworksPost(params *PcloudPvminstancesNetworksPostParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudPvminstancesNetworksPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudPvminstancesNetworksPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.pvminstances.networks.post",
		Method:             "POST",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudPvminstancesNetworksPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudPvminstancesNetworksPostCreated), nil

}

/*
PcloudPvminstancesPost creates a new power VM instance
*/
func (a *Client) PcloudPvminstancesPost(params *PcloudPvminstancesPostParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudPvminstancesPostOK, *PcloudPvminstancesPostCreated, *PcloudPvminstancesPostAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudPvminstancesPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.pvminstances.post",
		Method:             "POST",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudPvminstancesPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *PcloudPvminstancesPostOK:
		return value, nil, nil, nil
	case *PcloudPvminstancesPostCreated:
		return nil, value, nil, nil
	case *PcloudPvminstancesPostAccepted:
		return nil, nil, value, nil
	}
	return nil, nil, nil, nil

}

/*
PcloudPvminstancesPut updates a p cloud p VM instance
*/
func (a *Client) PcloudPvminstancesPut(params *PcloudPvminstancesPutParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudPvminstancesPutAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudPvminstancesPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.pvminstances.put",
		Method:             "PUT",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudPvminstancesPutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudPvminstancesPutAccepted), nil

}

/*
PcloudPvminstancesSnapshotsGetall gets all snapshots for this p VM instance
*/
func (a *Client) PcloudPvminstancesSnapshotsGetall(params *PcloudPvminstancesSnapshotsGetallParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudPvminstancesSnapshotsGetallOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudPvminstancesSnapshotsGetallParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.pvminstances.snapshots.getall",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/snapshots",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudPvminstancesSnapshotsGetallReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudPvminstancesSnapshotsGetallOK), nil

}

/*
PcloudPvminstancesSnapshotsPost creates a p VM instance snapshot
*/
func (a *Client) PcloudPvminstancesSnapshotsPost(params *PcloudPvminstancesSnapshotsPostParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudPvminstancesSnapshotsPostAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudPvminstancesSnapshotsPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.pvminstances.snapshots.post",
		Method:             "POST",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/snapshots",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudPvminstancesSnapshotsPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudPvminstancesSnapshotsPostAccepted), nil

}

/*
PcloudPvminstancesSnapshotsRestorePost restores a p VM instance snapshot
*/
func (a *Client) PcloudPvminstancesSnapshotsRestorePost(params *PcloudPvminstancesSnapshotsRestorePostParams, authInfo runtime.ClientAuthInfoWriter) (*PcloudPvminstancesSnapshotsRestorePostAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudPvminstancesSnapshotsRestorePostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pcloud.pvminstances.snapshots.restore.post",
		Method:             "POST",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/snapshots/{snapshot_id}/restore",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudPvminstancesSnapshotsRestorePostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PcloudPvminstancesSnapshotsRestorePostAccepted), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
