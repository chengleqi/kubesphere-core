package controllers

type Controller interface {
	initListerAndInformer()
	sync(chan struct{})
}
