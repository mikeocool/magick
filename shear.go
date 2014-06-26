package magick

// #include <magick/api.h>
// #include "bridge.h"
import "C"

// AffineTransform returns a new image created by transforming
// the original image as dictated by the affine matrix.
func (im *Image) AffineTransform(m *AffineMatrix) (*Image, error) {
	return im.applyDataFunc("transforming (affine)", C.ImageDataFunc(C.AffineTransformImage), m.matrix())
}

func (im *Image) IntegralRotate(rotations int) (*Image, error) {
    var ex C.ExceptionInfo
    C.GetExceptionInfo(&ex)
    defer C.DestroyExceptionInfo(&ex)
    rotated := C.IntegralRotateImage(im.image, C.size_t(rotations), &ex)
    return checkImage(rotated, nil, &ex, "rotated")
}
