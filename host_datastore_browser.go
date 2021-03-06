/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

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

package govmomi

import (
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

type HostDatastoreBrowser struct {
	types.ManagedObjectReference

	c *Client
}

func NewHostDatastoreBrowser(c *Client, ref types.ManagedObjectReference) *HostDatastoreBrowser {
	return &HostDatastoreBrowser{
		ManagedObjectReference: ref,
		c: c,
	}
}

func (b HostDatastoreBrowser) Reference() types.ManagedObjectReference {
	return b.ManagedObjectReference
}

func (b HostDatastoreBrowser) SearchDatastore(datastorePath string, searchSpec *types.HostDatastoreBrowserSearchSpec) (*Task, error) {
	req := types.SearchDatastore_Task{
		This:          b.Reference(),
		DatastorePath: datastorePath,
		SearchSpec:    searchSpec,
	}

	res, err := methods.SearchDatastore_Task(context.TODO(), b.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(b.c, res.Returnval), nil
}

func (b HostDatastoreBrowser) SearchDatastoreSubFolders(datastorePath string, searchSpec *types.HostDatastoreBrowserSearchSpec) (*Task, error) {
	req := types.SearchDatastoreSubFolders_Task{
		This:          b.Reference(),
		DatastorePath: datastorePath,
		SearchSpec:    searchSpec,
	}

	res, err := methods.SearchDatastoreSubFolders_Task(context.TODO(), b.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(b.c, res.Returnval), nil
}
