/*
 * Nuts consent bridge api
 * Copyright (C) 2019. Nuts community
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package pkg

import "sync"

const ConfigAddress = "address"
const ConfigAddressDefault = "localhost:8080"

// BridgeClientConfig holds the configuration for connecting to the consent-bridge
type BridgeClientConfig struct {
	Address string
}

var oneConfig sync.Once
var instance *BridgeClientConfig

func ConfigInstance() *BridgeClientConfig {
	oneConfig.Do(func() {
		instance = &BridgeClientConfig{}
	})

	return instance
}