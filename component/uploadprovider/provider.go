package uploadprovider

import (
	"200lab-project-1/common"
	"context"
)

type UploadProvider interface {
	SaveFileUploaded(context context.Context, data []byte, dst string) (*common.Image, error)
}
