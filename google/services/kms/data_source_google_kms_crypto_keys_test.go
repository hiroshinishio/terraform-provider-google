// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package kms_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
)

func TestAccDataSourceGoogleKmsCryptoKeys_basic(t *testing.T) {
	kms := acctest.BootstrapKMSKey(t)

	id := kms.KeyRing.Name + "/cryptoKeys"

	context := map[string]interface{}{
		"key_ring":      kms.KeyRing.Name,
		"random_suffix": acctest.RandString(t, 10),
		"filter":        "", // Can be overridden using 2nd argument to config funcs
	}

	randomString := acctest.RandString(t, 10)
	filterNameFindSharedKeys := "filter = \"name:tftest-shared-\""
	filterNameFindsNoKeys := fmt.Sprintf("filter = \"name:%s\"", randomString)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGoogleKmsCryptoKeys_basic(context, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.google_kms_crypto_keys.all_keys_in_ring", "id", id),
					resource.TestCheckResourceAttr("data.google_kms_crypto_keys.all_keys_in_ring", "key_ring", kms.KeyRing.Name),
					resource.TestMatchResourceAttr("data.google_kms_crypto_keys.all_keys_in_ring", "keys.#", regexp.MustCompile("[1-9]+[0-9]*")),
				),
			},
			{
				Config: testAccDataSourceGoogleKmsCryptoKeys_basic(context, filterNameFindSharedKeys),
				Check: resource.ComposeTestCheckFunc(
					// This filter should retrieve keys in the bootstrapped KMS key ring used by the test
					resource.TestCheckResourceAttr("data.google_kms_crypto_keys.all_keys_in_ring", "id", id),
					resource.TestCheckResourceAttr("data.google_kms_crypto_keys.all_keys_in_ring", "key_ring", kms.KeyRing.Name),
					resource.TestMatchResourceAttr("data.google_kms_crypto_keys.all_keys_in_ring", "keys.#", regexp.MustCompile("[1-9]+[0-9]*")),
				),
			},
			{
				Config: testAccDataSourceGoogleKmsCryptoKeys_basic(context, filterNameFindsNoKeys),
				Check: resource.ComposeTestCheckFunc(
					// This filter should retrieve no keys
					resource.TestCheckResourceAttr("data.google_kms_crypto_keys.all_keys_in_ring", "id", id),
					resource.TestCheckResourceAttr("data.google_kms_crypto_keys.all_keys_in_ring", "key_ring", kms.KeyRing.Name),
					resource.TestCheckResourceAttr("data.google_kms_crypto_keys.all_keys_in_ring", "keys.#", "0"),
				),
			},
		},
	})
}

func testAccDataSourceGoogleKmsCryptoKeys_basic(context map[string]interface{}, filter string) string {
	context["filter"] = filter

	return acctest.Nprintf(`
data "google_kms_crypto_keys" "all_keys_in_ring" {
  key_ring = "%{key_ring}"
  %{filter}
}
`, context)
}
