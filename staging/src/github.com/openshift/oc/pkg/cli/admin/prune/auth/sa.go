package auth

import (
//	"fmt"
	"io"

	corev1 "k8s.io/api/core/v1"
//	kerrors "k8s.io/apimachinery/pkg/api/errors"
//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"

	authv1client "github.com/openshift/client-go/authorization/clientset/versioned/typed/authorization/v1"
//	oauthv1client "github.com/openshift/client-go/oauth/clientset/versioned/typed/oauth/v1"
//	securityv1client "github.com/openshift/client-go/security/clientset/versioned/typed/security/v1"
//	userv1client "github.com/openshift/client-go/user/clientset/versioned/typed/user/v1"
)

func reapForServiceAccount(
	authorizationClient authv1client.AuthorizationV1Interface,
  nsname string,
	name string,
	out io.Writer) error {

	errors := []error{}

  removedSubject := corev1.ObjectReference{Kind: "ServiceAccount", Name: name, Namespace: nsname}
	errors = append(errors, reapClusterBindings(removedSubject, authorizationClient, out)...)
	errors = append(errors, reapNamespacedBindings(removedSubject, authorizationClient, out)...)

	return utilerrors.NewAggregate(errors)
}
