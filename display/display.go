/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package display

import (
	"errors"
	"fmt"
	"github.com/apache/skywalking-cli/display/json"
	"github.com/apache/skywalking-cli/display/table"
	"github.com/apache/skywalking-cli/display/yaml"
	"github.com/urfave/cli"
	"strings"
)

const (
	Json  string = "json"
	Yaml  string = "yaml"
	Table string = "table"
)

// Display the object in the style specified in flag --display
func Display(ctx *cli.Context, object interface{}) error {
	displayStyle := ctx.GlobalString("display")

	switch strings.ToLower(displayStyle) {
	case Json:
		return json.Display(object)
	case Yaml:
		return yaml.Display(object)
	case Table:
		return table.Display(object)
	default:
		return errors.New(fmt.Sprintf("unsupported display style: %s", displayStyle))
	}
}
