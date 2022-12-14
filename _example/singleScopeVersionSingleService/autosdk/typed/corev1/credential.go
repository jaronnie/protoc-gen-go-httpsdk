// Code generated by protoc-gen-go-httpsdk. DO NOT EDIT.
package corev1

import (
	"context"

	"github.com/jaronnie/autosdk/pb/corev1"
	"github.com/jaronnie/autosdk/rest"
)

type CredentialGetter interface {
	Credential() CredentialInterface
}

type CredentialInterface interface {
	InitCredential(ctx context.Context, param *corev1.Credential) (*corev1.Credential, error)
	UpdateCredential(ctx context.Context, param *corev1.Credential) (*corev1.Credential, error)

	CredentialExpansion
}

type credentialClient struct {
	client rest.Interface
}

func newCredentialClient(c *Corev1Client) *credentialClient {
	return &credentialClient{
		client: c.RESTClient(),
	}
}

func (x *credentialClient) InitCredential(ctx context.Context, param *corev1.Credential) (*corev1.Credential, error) {
	var resp corev1.Credential
	err := x.client.Verb("POST").
		SubPath(
			"/gateway/core/api/v1/credential/init",
		).
		Params().
		Body(param).
		Do(ctx).
		Into(&resp, false)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (x *credentialClient) UpdateCredential(ctx context.Context, param *corev1.Credential) (*corev1.Credential, error) {
	var resp corev1.Credential
	err := x.client.Verb("POST").
		SubPath(
			"/gateway/core/api/v1/credential/update",
		).
		Params().
		Body(param).
		Do(ctx).
		Into(&resp, false)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}
