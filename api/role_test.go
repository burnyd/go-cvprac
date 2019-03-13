//
// Copyright (c) 2016-2017, Arista Networks, Inc. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//   * Redistributions of source code must retain the above copyright notice,
//   this list of conditions and the following disclaimer.
//
//   * Redistributions in binary form must reproduce the above copyright
//   notice, this list of conditions and the following disclaimer in the
//   documentation and/or other materials provided with the distribution.
//
//   * Neither the name of Arista Networks nor the names of its
//   contributors may be used to endorse or promote products derived from
//   this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL ARISTA NETWORKS
// BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR
// BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
// WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE
// OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN
// IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//

package cvpapi

import (
	"errors"
	"testing"
)

func Test_CvpGetAllRolesRetError_UnitTest(t *testing.T) {
	clientErr := errors.New("Client error")
	expectedErr := errors.New("GetAllRoles: Client error")

	client := NewMockClient("", clientErr)
	api := NewCvpRestAPI(client)

	_, err := api.GetAllRoles(0, 0)
	if err.Error() != expectedErr.Error() {
		t.Fatalf("Expected Client error: [%v] Got: [%v]", expectedErr, err)
	}
}

func Test_CvpGetAllRolesJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("{", nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetAllRoles(0, 0); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpGetAllRolesEmptyJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("", nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetAllRoles(0, 0); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpGetAllRolesReturnError_UnitTest(t *testing.T) {
	respStr := `{"errorCode": "112498",
  				 "errorMessage": "Unauthorized User"}`

	client := NewMockClient(respStr, nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetAllRoles(0, 0); err == nil {
		t.Fatal("Error should be returned")
	}
}

func Test_CvpGetAllRolesValid_UnitTest(t *testing.T) {
	client := NewMockClient("{}", nil)
	api := NewCvpRestAPI(client)

	_, err := api.GetAllRoles(0, 0)
	if err != nil {
		t.Fatalf("Valid case failed with error: %v", err)
	}
}

func Test_CvpGetAllRolesValid2_UnitTest(t *testing.T) {
	data := `{
  "total": 2,
  "roles": [
    {
      "name": "network-admin",
      "key": "network-admin",
      "description": "",
      "moduleListSize": 19,
      "createdBy": "cvp system",
      "createdOn": 1548362619198,
      "moduleList": [
        {
          "name": "image",
          "mode": "rw"
        },
        {
          "name": "ztp",
          "mode": "rw"
        },
        {
          "name": "configlet",
          "mode": "rw"
        },
        {
          "name": "task",
          "mode": "rw"
        },
        {
          "name": "inventory",
          "mode": "rw"
        },
        {
          "name": "label",
          "mode": "rw"
        },
        {
          "name": "danz",
          "mode": "rw"
        },
        {
          "name": "aaa",
          "mode": "rw"
        },
        {
          "name": "account",
          "mode": "rw"
        },
        {
          "name": "snapshot",
          "mode": "rw"
        },
        {
          "name": "changeControl",
          "mode": "rw"
        },
        {
          "name": "ssl",
          "mode": "rw"
        },
        {
          "name": "purge",
          "mode": "rw"
        },
        {
          "name": "cvpTheme",
          "mode": "rw"
        },
        {
          "name": "networkProvisioning",
          "mode": "rw"
        },
        {
          "name": "audit",
          "mode": "rw"
        },
        {
          "name": "workflow",
          "mode": "rw"
        },
        {
          "name": "cloudManager",
          "mode": "rw"
        },
        {
          "name": "publicCloudAccounts",
          "mode": "rw"
        }
      ]
    },
    {
      "name": "network-operator",
      "key": "network-operator",
      "description": "",
      "moduleListSize": 18,
      "createdBy": "cvp system",
      "createdOn": 1548362619219,
      "moduleList": [
        {
          "name": "image",
          "mode": "r"
        },
        {
          "name": "ztp",
          "mode": "r"
        },
        {
          "name": "configlet",
          "mode": "r"
        },
        {
          "name": "task",
          "mode": "r"
        },
        {
          "name": "inventory",
          "mode": "r"
        },
        {
          "name": "label",
          "mode": "r"
        },
        {
          "name": "danz",
          "mode": "r"
        },
        {
          "name": "aaa",
          "mode": "r"
        },
        {
          "name": "account",
          "mode": "r"
        },
        {
          "name": "snapshot",
          "mode": "r"
        },
        {
          "name": "changeControl",
          "mode": "r"
        },
        {
          "name": "ssl",
          "mode": "r"
        },
        {
          "name": "purge",
          "mode": "r"
        },
        {
          "name": "cvpTheme",
          "mode": "rw"
        },
        {
          "name": "networkProvisioning",
          "mode": "r"
        },
        {
          "name": "audit",
          "mode": "r"
        },
        {
          "name": "workflow",
          "mode": "r"
        },
        {
          "name": "cloudManager",
          "mode": "r"
        }
      ]
    }
  ],
  "users": {
    "network-admin": 2,
    "network-operator": 5
  }
}`
	client := NewMockClient(data, nil)
	api := NewCvpRestAPI(client)

	_, err := api.GetAllRoles(0, 0)
	if err != nil {
		t.Fatalf("Valid case failed with error: %v", err)
	}
}

func Test_CvpGetRoleRetError_UnitTest(t *testing.T) {
	clientErr := errors.New("Client error")
	expectedErr := errors.New("GetRole: Client error")

	client := NewMockClient("", clientErr)
	api := NewCvpRestAPI(client)

	_, err := api.GetRole("role")
	if err.Error() != expectedErr.Error() {
		t.Fatalf("Expected Client error: [%v] Got: [%v]", expectedErr, err)
	}
}

func Test_CvpGetRoleJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("{", nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetRole("role"); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpGetRoleEmptyJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("", nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetRole("role"); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpGetRoleReturnError_UnitTest(t *testing.T) {
	respStr := `{"errorCode": "112498",
  				 "errorMessage": "Unauthorized User"}`

	client := NewMockClient(respStr, nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetRole("role"); err == nil {
		t.Fatal("Error should be returned")
	}
}

func Test_CvpGetRoleValid_UnitTest(t *testing.T) {
	data := `{
  "name": "network-admin",
  "key": "network-admin",
  "description": "",
  "moduleListSize": 19,
  "createdBy": "cvp system",
  "createdOn": 1548362619198,
  "moduleList": [
    {
      "name": "image",
      "mode": "rw"
    },
    {
      "name": "ztp",
      "mode": "rw"
    },
    {
      "name": "configlet",
      "mode": "rw"
    },
    {
      "name": "task",
      "mode": "rw"
    },
    {
      "name": "inventory",
      "mode": "rw"
    },
    {
      "name": "label",
      "mode": "rw"
    },
    {
      "name": "danz",
      "mode": "rw"
    },
    {
      "name": "aaa",
      "mode": "rw"
    },
    {
      "name": "account",
      "mode": "rw"
    },
    {
      "name": "snapshot",
      "mode": "rw"
    },
    {
      "name": "changeControl",
      "mode": "rw"
    },
    {
      "name": "ssl",
      "mode": "rw"
    },
    {
      "name": "purge",
      "mode": "rw"
    },
    {
      "name": "cvpTheme",
      "mode": "rw"
    },
    {
      "name": "networkProvisioning",
      "mode": "rw"
    },
    {
      "name": "audit",
      "mode": "rw"
    },
    {
      "name": "workflow",
      "mode": "rw"
    },
    {
      "name": "cloudManager",
      "mode": "rw"
    },
    {
      "name": "publicCloudAccounts",
      "mode": "rw"
    }
  ]
}`
	client := NewMockClient(data, nil)
	api := NewCvpRestAPI(client)

	_, err := api.GetRole("network-admin")
	if err != nil {
		t.Fatalf("Valid case failed with error: %v", err)
	}

}
