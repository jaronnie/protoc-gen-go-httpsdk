// Code generated by protoc-gen-go-httpsdk. DO NOT EDIT.
package fake

import (
	"github.com/jaronnie/autosdk/typed"
	corev1 "github.com/jaronnie/autosdk/typed/corev1"
	fakecorev1 "github.com/jaronnie/autosdk/typed/corev1/fake"
	"github.com/jaronnie/autosdk/typed/fake"
)

type Clientset struct{}

func (x *Clientset) Direct() typed.DirectInterface {
	return &fake.FakeDirect{}
}

func (x *Clientset) Corev1() corev1.Corev1Interface {
	return &fakecorev1.FakeCorev1{}
}
