Nuts consent bridge go client
=============================

Library for contacting consent-bridge. Exposed as Nuts-go engine.

.. image:: https://circleci.com/gh/nuts-foundation/consent-bridge-go-client.svg?style=svg
    :target: https://circleci.com/gh/nuts-foundation/consent-bridge-go-client
    :alt: Build Status

.. image:: https://codecov.io/gh/nuts-foundation/consent-bridge-go-client/branch/master/graph/badge.svg
    :target: https://codecov.io/gh/nuts-foundation/consent-bridge-go-client

.. image:: https://api.codacy.com/project/badge/Grade/ad0e073161b4411690e3219a6c5b2331
    :target: https://www.codacy.com/app/nuts-foundation/consent-bridge-go-client

The consent bridge go client is written in Go and should be part of nuts-go as an engine.

Dependencies
************

This projects is using go modules, so version > 1.12 is recommended. 1.10 would be a minimum. Currently Sqlite is used as database backend.

Running tests
*************

Tests can be run by executing

.. code-block:: shell

    go test ./...

Building
********

This project is part of https://github.com/nuts-foundation/nuts-go. If you do however would like a binary, just use ``go build``.

The client and server API is generated from the nuts-consent-store open-api spec:

.. code-block:: shell

    oapi-codegen -generate types,client -package api ../nuts-consent-bridge/docs/_static/nuts-consent-bridge.yaml > api/generated.go


README
******

The readme is auto-generated from a template and uses the documentation to fill in the blanks.

.. code-block:: shell

    ./generate_readme.sh

This script uses ``rst_include`` which is installed as part of the dependencies for generating the documentation.

Documentation
*************

To generate the documentation, you'll need python3, sphinx and a bunch of other stuff. See :ref:`nuts-documentation-development-documentation`
The documentation can be build by running

.. code-block:: shell

    /docs $ make html

The resulting html will be available from ``docs/_build/html/index.html``

Configuration
*************

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

