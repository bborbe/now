// Copyright (c) 2018 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package now

import (
	"fmt"
	"io"
	"time"

	"github.com/golang/glog"
	"github.com/pkg/errors"
)

var DefaultLocations = Locations{
	"Atlantic/Azores",
	"UTC",
	"Atlantic/Canary",
	"Europe/Berlin",
	"Asia/Kolkata",
}

type Location string

func (l Location) String() string {
	return string(l)
}

func (l Location) Load() (*time.Location, error) {
	location, err := time.LoadLocation(l.String())
	return location, errors.Wrapf(err, "load location %s failed", l.String())
}

func (l Location) Write(writer io.Writer) error {
	location, err := l.Load()
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(writer, "%s %s\n",
		time.Now().In(location).Format("2006-01-02T15:04:05"),
		l.String(),
	)
	if err != nil {
		return errors.Wrap(err, "write failed")
	}
	return nil
}

type Locations []Location

func (l Locations) Write(writer io.Writer) error {
	glog.V(4).Infof("write now ...")
	for _, location := range l {
		if err := location.Write(writer); err != nil {
			return err
		}
	}
	return nil
}
