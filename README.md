## CBT Datapath Service

This repository contains prototype that can serve the CSI changed block tracking API.


### Quick Start

TODO: Add deploy steps


#### Get Token

The CBT token can be generated with `VolumeSnapshotDeltaToken` APIs either using kubectl client or client-go sdks

```bash
cat <<EOF | kubectl create -oyaml -f -                                                                                                                    
apiVersion: cbt.storage.k8s.io/v1alpha1
kind: VolumeSnapshotDeltaToken
metadata:
  name: test
spec:
  baseVolumeSnapshotName: vs-00                                                               
  targetVolumeSnapshotName: vs-01                                                                                                                                                            mode: block
EOF                                                                                                                                                                                        apiVersion: cbt.storage.k8s.io/v1alpha1
kind: VolumeSnapshotDeltaToken
metadata:                           
  creationTimestamp: "2023-03-06T17:26:52Z"
  name: test
  namespace: cbt-client
spec:
  baseVolumeSnapshotName: vs-00
  mode: block
  targetVolumeSnapshotName: vs-01
status:
  cabundle: xxxx
  token: xxxxxxxx
  url: cbt-datapath.cbt-svc.svc:80
```
