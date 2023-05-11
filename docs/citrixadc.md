### Use with Citrix ADC

Example:

```
$ export NS_LOGIN=<LOGIN>
$ export NS_PASSWORD=<PASSWORD>
$ export NS_URL=<URL>
$ export NS_SSLVERIFY=<SSL_VERIFY>
$ terraformer import citrixadc --resources=lbmonitor,cspolicy
```

List of supported Citrix ADC services:


* `servicegroup`
* `cspolicy`
* `lbmonitor`
* `servicegroup_lbmonitor_binding`
* `server`
* `servicegroup_servicegroupmember_binding`
* `lbvserver`
* `rewritepolicy`
* `lbvserver_rewritepolicy_binding`
* `service`
* `lbvserver_service_binding`
* `csvserver`
* `csvserver_cspolicy_binding`
* `responderpolicy`
* `lbvserver_responderpolicy_binding`
* `lbvserver_servicegroup_binding`
* `sslcipher`
* `rewriteaction`
* `sslvserver`
* `responderaction`
