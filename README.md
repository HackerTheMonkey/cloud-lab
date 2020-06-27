# codematters

Requirements:

```
given domainDoesNotExist then newDomainIsCreated
given domainAlreadyExist then noNewDomainIsCreated

given floatingIPDoesNotExist then newFloatingIPIsCreated
given floatingIPExists then noNewFloatingIPIsCreated

given noARecordExistsInDomainAgainstFloatingIP then createNewARecord
given RecordExistsInDomainAgainstFloatingIP then noNewARecordIsCreated

given rootDropletExists then newRootDropletIsDeleted
given noRootDropletExists then newRootDropletIsCreated floatingIPIsAssigned

TODO: make changes to accommodate for blue-green deployments

```
