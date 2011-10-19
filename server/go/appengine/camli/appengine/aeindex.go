/*
Copyright 2011 Google Inc.

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

package appengine

import (
	"fmt"
	"os"
	"strings"
	"time"

	"camli/blobref"
	"camli/blobserver"
	"camli/jsonconfig"
	"camli/search"
)

type appengineIndex struct {
	*blobserver.NoImplStorage
	namespace string
}

var _ search.Index = (*appengineIndex)(nil)

func indexFromConfig(ld blobserver.Loader, config jsonconfig.Obj) (storage blobserver.Storage, err os.Error) {
	sto := &appengineIndex{}
	sto.namespace = config.OptionalString("namespace", "")
	if err := config.Validate(); err != nil {
		return nil, err
	}
	// TODO(bradfitz): merge this namespace checking logic with storage.go's copy
	if strings.Contains(sto.namespace, "|") {
		return nil, fmt.Errorf("no pipe allowed in namespace %q", sto.namespace)
	}
	if strings.Contains(sto.namespace, "\x00") {
		return nil, fmt.Errorf("no zero byte allowed in namespace %q", sto.namespace)
	}
	if sto.namespace == "-" {
		return nil, fmt.Errorf("reserved namespace %q", sto.namespace)
	}
	if sto.namespace == "" {
		sto.namespace = "-"
	}
	return sto, nil
}

func (x *appengineIndex) GetRecentPermanodes(dest chan *search.Result,
	owner []*blobref.BlobRef,
	limit int) os.Error {
	// TODO(bradfitz): this will need to be a context wrapper too, like storage
	return os.NewError("TODO: GetRecentPermanodes")
}

func (x *appengineIndex) SearchPermanodesWithAttr(dest chan<- *blobref.BlobRef,
	request *search.PermanodeByAttrRequest) os.Error {
	return os.NewError("TODO: SearchPermanodesWithAttr")
}

func (x *appengineIndex) GetOwnerClaims(permaNode, owner *blobref.BlobRef) (search.ClaimList, os.Error) {
	return nil, os.NewError("TODO: GetOwnerClaims")
}

func (x *appengineIndex) GetBlobMimeType(blob *blobref.BlobRef) (mime string, size int64, err os.Error) {
	err = os.NewError("TODO: GetBlobMimeType")
	return
}


func (x *appengineIndex) ExistingFileSchemas(bytesRef *blobref.BlobRef) ([]*blobref.BlobRef, os.Error) {
	return nil, os.NewError("TODO: xxx")
}

func (x *appengineIndex) GetFileInfo(fileRef *blobref.BlobRef) (*search.FileInfo, os.Error) {
	return nil, os.NewError("TODO: GetFileInfo")
}

func (x *appengineIndex) PermanodeOfSignerAttrValue(signer *blobref.BlobRef, attr, val string) (*blobref.BlobRef, os.Error) {
	return nil, os.NewError("TODO: PermanodeOfSignerAttrValue")
}

func (x *appengineIndex) PathsOfSignerTarget(signer, target *blobref.BlobRef) ([]*search.Path, os.Error) {
	return nil, os.NewError("TODO: PathsOfSignerTarget")
}

func (x *appengineIndex) PathsLookup(signer, base *blobref.BlobRef, suffix string) ([]*search.Path, os.Error) {
	return nil, os.NewError("TODO: PathsLookup")
}

func (x *appengineIndex) PathLookup(signer, base *blobref.BlobRef, suffix string, at *time.Time) (*search.Path, os.Error) {
	return nil, os.NewError("TODO: PathLookup")
}
