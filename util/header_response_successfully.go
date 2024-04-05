package util

import "github.com/indrawanagung/food-order-api/model/web"

func HeaderResponseSuccessfully() web.Header {
	return web.Header{
		Message: "Your request has been processed successfully",
		Error:   false,
	}
}
