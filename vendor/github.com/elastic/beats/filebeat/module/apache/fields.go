// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package apache

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "apache", asset.ModuleFieldsPri, AssetApache); err != nil {
		panic(err)
	}
}

// AssetApache returns asset data.
// This is the base64 encoded gzipped contents of module/apache.
func AssetApache() string {
	return "eJysmM+OozgQxu95ilLf24c95rDSaqXVHmakkbrvyDEV8DS4mHKRFm8/gkBP/hhiQzi1TNf3/eyqGLte4QO7PehGmxJ3AGKlwj28/DMMvOwAcvSGbSOW3B7+3gEAnF/Cd8rbqg/yJbFkhtzRFnsQbvvBo8Uq9/sh4BWcrnGy+WsYA5CuwT0UTG0zjgS8Br/Kao8ejsRw0ObjU3MOhupGiz3YykoHn1ZKoCqfLEb7UeKS5YrHGPT+aziEFAq/lGCsSTCzzdXbSUn35DdvGi3lHjy1bFDpPOdrhP6pbcH6vArjat4be1+phknIUJXofV4jdZ69WhCK5TC2KZE3UwRlYhgOlHeZRyfq0Ane2j0CKUUaxegbch5VrxWUiQFpPXLW/5mI0MepQFyMZ41SUr5uzr9a9KKCClHT5dTSa7lSxLawTq8pth47OyF7S27NjMOhMc5TfWSG8tTsXheYFy2tD+nEcRyROfnHdpXvGY0Ye12gkxXFnQ2Bsamf/23N+99u2jCzcV9K5niy5jYLy1MKTmtWZ2lJL0Fq/ZNu07GCY04mGsO6p2DMyMRiNFpMuR1jTiYWI7Afr6CYUYmFoJBJIgJ5dWyrKvRhSkPJnlSn5DeXag/znGrtYTYWLPm5tU1n2Zqk8A67hmVeKeZLUSDNHIrX7NSGnFiHTrYs9HjmLpDUQ73Y9TbUOuEus55Cn/RVaA8VY+EqMsO/bYdaUIqFYSwsuSflb1ksOnlWumcV1IJU4go9r5QeC86hTUjIfLUzpl+QKzxh6i2hokKF4uKuQ97rIvWIHo6K8Wts6t2rYRquvPeRMX6y2k9KRp2rVa711OtJMR5v+EMNqaDC3cF/B3ftot2lWUy36NydOtekGoc3doBmrPrnX3KirfOjxdCkkhInjP/f33/AG/IJeTTrq/uLK8QGKV2eD+w+iW9TusDbP29v32BShfFO/IdoFmShzbMa46w5mKiVu8+Cy39jTpjqy6QMukl5WKj/VVN/L3HU7BORt8a6YiCsqCgwn/YjtfsdAAD//9uTc4g="
}
