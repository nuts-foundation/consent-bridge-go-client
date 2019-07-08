.. _nuts-consent-bridge-go-client-configuration:

Nuts consent bridge go client configuration
###########################################

.. marker-for-readme

The following configuration parameters are available.

=====================================   ====================    ================================================================
Property                                Default                 Description
=====================================   ====================    ================================================================
nuts.cbridge.address                    localhost:8080          API Address of the consent bridge
=====================================   ====================    ================================================================

As with all other properties for nuts-go, they can be set through yaml:

.. sourcecode:: yaml

    cbridge:
       address: localhost:1323

as commandline property

.. sourcecode:: shell

    ./nuts --cbridge.address localhost:1323

Or by using environment variables

.. sourcecode:: shell

    NUTS_CBRIDGE_ADDRESS=localhost:1323 ./nuts