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

package engine

import (
	"github.com/nuts-foundation/consent-bridge-go-client/pkg"
	engine "github.com/nuts-foundation/nuts-go-core"
	"github.com/spf13/pflag"
)

func NewConsentBridgeClientEngine() *engine.Engine {
	return &engine.Engine{
		Name:      "ConsentBridgeClient",
		Config:    pkg.ConfigInstance(),
		ConfigKey: "cbridge",
		FlagSet:   flagSet(),
	}
}

func flagSet() *pflag.FlagSet {
	flags := pflag.NewFlagSet("cbridge", pflag.ContinueOnError)

	flags.String(pkg.ConfigAddress, pkg.ConfigAddressDefault, "API Address of the consent bridge")

	return flags
}
