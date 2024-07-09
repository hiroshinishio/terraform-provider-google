// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package compute_test

import (
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
)

func TestAccComputeInstanceSettings_instanceSettingsBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeInstanceSettingsDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeInstanceSettings_instanceSettingsBasicExample(context),
			},
			{
				ResourceName:            "google_compute_instance_settings.gce_instance_settings",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"zone"},
			},
		},
	})
}

func testAccComputeInstanceSettings_instanceSettingsBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`

resource "google_compute_instance_settings" "gce_instance_settings" {
  zone = "us-east7-b"
  metadata {
    items = {
      foo = "baz"
    }
  }
}
`, context)
}

func testAccCheckComputeInstanceSettingsDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_instance_settings" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			log.Printf("[DEBUG] Ignoring destroy during test")
		}

		return nil
	}
}
