package storage

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/PrasadG193/cbt-datapath/pkg/apis/cbt/v1alpha1"
	cbtv1alpha1 "github.com/PrasadG193/cbt-datapath/pkg/apis/cbt/v1alpha1"
	"github.com/google/uuid"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	genericregistry "k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	builderrest "sigs.k8s.io/apiserver-runtime/pkg/builder/rest"

	"k8s.io/klog"
)

var _ builderrest.ResourceHandlerProvider = CBTHandlerProvider

func CBTHandlerProvider(is *runtime.Scheme, _ genericregistry.RESTOptionsGetter) (rest.Storage, error) {
	obj := &cbtv1alpha1.VolumeSnapshotDeltaToken{}
	return &cbt{
		namespaceScoped: obj.NamespaceScoped(),
		newFunc:         obj.New,
		newListFunc:     obj.NewList,
	}, nil

}

type cbt struct {
	namespaceScoped bool
	newFunc         func() runtime.Object
	newListFunc     func() runtime.Object
}

func (c *cbt) New() runtime.Object {
	return c.newFunc()
}

// NewList returns an empty object that can be used with the List call.
// This object must be a pointer type for use with Codec.DecodeInto([]byte, runtime.Object)
func (c *cbt) NewList() runtime.Object {
	return c.newListFunc()
}

// NamespaceScoped returns true if the storage is namespaced
func (c *cbt) NamespaceScoped() bool {
	return c.namespaceScoped
}

// Create creates a new version of a resource.
func (c *cbt) Create(
	ctx context.Context,
	obj runtime.Object,
	createValidation rest.ValidateObjectFunc,
	options *metav1.CreateOptions) (runtime.Object, error) {

	casted, ok := obj.(*v1alpha1.VolumeSnapshotDeltaToken)
	if !ok {
		return nil, fmt.Errorf("")
	}
	casted.SetCreationTimestamp(metav1.Now())
	out := casted.DeepCopy()

	reqID := uuid.New().String()
	token := NewToken(reqID)
	ca, err := fetchCABundle()
	if err != nil {
		return nil, err
	}
	out.Status = v1alpha1.VolumeSnapshotDeltaTokenStatus{
		Token:    token.Token,
		URL:      token.URL,
		CABundle: ca,
	}

	klog.Infof("created VolumeSnapshotDeltaToken: %s", out.GetName())

	return out, nil
}

func fetchCABundle() ([]byte, error) {
	cacertFile := os.Getenv("CBT_SERVER_CA_BUNDLE")
	if cacertFile == "" {
		return nil, errors.New("Failed to read CA Bundle from " + cacertFile)
	}
	return os.ReadFile(cacertFile)
}
