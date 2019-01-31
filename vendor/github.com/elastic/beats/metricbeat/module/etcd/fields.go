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

package etcd

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "etcd", asset.ModuleFieldsPri, AssetEtcd); err != nil {
		panic(err)
	}
}

// AssetEtcd returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/etcd.
func AssetEtcd() string {
	return "eJzMl8+O2zYQxu9+ioEvbYHGD+BDL2mOBYogt6JoZslPFrEUqZBDO377grLktWzKf1YOWh4WC5nzzW+G5HD4gV6xXxNE6QWRGLFY0/KTKL1cEGlEFUwrxrs1/bYgom4m/eF1slgQBVhwxJo2vCCKEDFuE9f01zJGu/yVlrVIu/x7QVQZWB3XncYHctzg6DUP2bdZJfjU9l8KvvP4mo2+kvJO2LhIUVhMFKMiSc1COwRQAGuqgm/o05uLU4JTCgvWCMfPJZYrPHl8HGC65Bz0TsBWJ5PHCRvGOdopXuWt9TuEuFI+OYmjSVOwN4Dz+FKDXGpeEMhXVLGx0MROU0xKIcYqWfrMldDnPz9SwLeEKKNQpsCvwa968QuDIRDjBJvRatwQzOB3q13+U9a2LHBq/6xUU69H4gmsamqBQMaR1CBlU5SLiO9PbK+94i0CbzCZi6jYQv9TWc/ygKxKIcBdWsyUbfi7aVLzdFnjrsreu72OglHYaQ76d2wN5/V9nHiyzrzZvmK/80Ffbs/BNsJWz6tQWe1p9cnou0K6iZgta1By5lsCGQ0npjIIVPnQ/dIgF6vrR/iQY+Mqv3og3fewkdG5TnYn9nAkhkJ/xjWBE4WDiGnOD+g0UUEktQ8pPC2mIlf++zwUqU3s3f4Ux9pF7wFqu+K2hdP95XS4EYpE5WN/g+jtcjy4OV6CB1bnNajmSG3w+UpDefU6zhd2eme01IGlnLMr9e5uzJe9IFKbew8o7/QJp4kZBGZr3IZ+HqoceWf3v0xjt6+bHwp8TOj7mYvwEU7/2L1RCqJrP8dbI57fnCPC/8OuyCBdfvuD32V3RV/ynC3b1E1KTqMyDpq8o2jcxg7VeOhezhrDUZj/+S56fozlYB+u8HcURVBWpF0NdxLUjuPBHQo9wxnTRObfxQMjNcJwSfhw7JcuGxbxAXM6lvw86UROX3jGKZvyorrKh6brx4hffJIuU75F6L4Vz2LNTlvoec3OBpNvmOlqMjIuvFduvFVGaZ3jP873r2EheDdBbz6PIbWaZzD05vMYVMAcht58JoNvWg5gp+OO2/ezjGWexTRzp1wIPc5V1MX31gQ80A4UVXYsqka4Hta/AQAA//9pXiwr"
}