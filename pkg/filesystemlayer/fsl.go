// Copyright 2021, Chef.  All rights reserved.
// https://github.com/cfeeling/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package filesystemlayer

import "os"

var _ IFileSystemLayer = &FSLDisk{}
var _ IFileSystemLayer = &FSLMemory{}

var _ IFile = &os.File{}
var _ IFile = &file{}
