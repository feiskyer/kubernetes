<!-- BEGIN MUNGE: UNVERSIONED_WARNING -->

<!-- BEGIN STRIP_FOR_RELEASE -->

<img src="http://kubernetes.io/img/warning.png" alt="WARNING"
     width="25" height="25">
<img src="http://kubernetes.io/img/warning.png" alt="WARNING"
     width="25" height="25">
<img src="http://kubernetes.io/img/warning.png" alt="WARNING"
     width="25" height="25">
<img src="http://kubernetes.io/img/warning.png" alt="WARNING"
     width="25" height="25">
<img src="http://kubernetes.io/img/warning.png" alt="WARNING"
     width="25" height="25">

<h2>PLEASE NOTE: This document applies to the HEAD of the source tree</h2>

If you are using a released version of Kubernetes, you should
refer to the docs that go with that version.

<strong>
The latest 1.0.x release of this document can be found
[here](http://releases.k8s.io/release-1.0/docs/hyperstack/user-admin.md).

Documentation for other releases can be found at
[releases.k8s.io](http://releases.k8s.io).
</strong>
--

<!-- END STRIP_FOR_RELEASE -->

<!-- END MUNGE: UNVERSIONED_WARNING -->

This document aims to describe how to manage tenants and users in Hypernetes.

### Tenant management

Hypernetes manages tenants and users using OpenStack Keystone. To create a tenant, you must create it in Keystone first, then expose it to Kubernetes:

Create OpenStack `demo` tenant

```sh
openstack project create demo
```

Create a spec file of demo tenant `demo-tenant.yaml`

```yaml
apiVersion: v1
kind: Tenant
metadata:
  name: demo
```

Create demo tenant

```sh
# kubectl create -f ./demo-tenant.yaml
# kubectl get tenant
NAME      LABELS    STATUS    AGE
default   <none>    Active    23h
demo      <none>    Active    57m
```

### User management

You can simply add users to any existing tenant using an OpenStack client. For exmaple, these commands add a `demouser` to tenant `demo`:

```sh
openstack user create --password demo demouser
openstack role add --project demo --user demouser _member_
```



<!-- BEGIN MUNGE: GENERATED_ANALYTICS -->
[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/docs/hyperstack/user-admin.md?pixel)]()
<!-- END MUNGE: GENERATED_ANALYTICS -->
